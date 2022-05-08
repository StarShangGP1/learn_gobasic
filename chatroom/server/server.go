package server

import (
	"chat-project/api"
	"chat-project/global"
	"chat-project/models"
	"chat-project/proto"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

type Connection struct {
	stream       proto.ChatService_CreateChatStreamServer
	Nickname     string
	Account      string
	chattingWith []string
	active       bool
	error        chan error
}

type ChatServer struct {
	ConnDb     *gorm.DB
	Connection map[string]*Connection
	mu         sync.RWMutex
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		ConnDb:     connectDb(),
		Connection: make(map[string]*Connection),
		mu:         sync.RWMutex{},
	}
}

// Crontab 备份聊天记录到文件中
func Crontab() {
	db := connectDb()
	c := cron.New()
	EntryID, err := c.AddFunc("00 */1 * * *", func() {
		BackupRecords(db)
	})
	fmt.Println(time.Now(), EntryID, err)

	c.Start()
}

// BackupRecords 备份聊天记录到文件中，不考虑服务端Down机情况下
// 可以数据库通过增加字段处理备份过的打Tag OR 新建表记录最后备份时的状态
func BackupRecords(db *gorm.DB) {
	var chats []*models.Chat

	db.Table("chat_records").Find(&chats)

	for _, chat := range chats {
		records := api.Chat{
			Sender:          chat.Sender,
			SenderAccount:   chat.SenderAccount,
			Receiver:        chat.Receiver,
			ReceiverAccount: chat.ReceiverAccount,
			Message:         chat.Message,
			Status:          chat.Status,
		}

		data, err := json.Marshal(records)
		if err != nil {
			fmt.Println("marshal 出错：", err)
		}

		file, err := os.OpenFile(global.FilePathRecord, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666) // linux file permission settings
		if err != nil {
			fmt.Println("无法打开文件", global.FilePathRecord, "错误信息是：", err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.Write(append(data, '\n'))
	}

}

// GetAllUsers 查看所有注册用户
func (c *ChatServer) GetAllUsers(ctx context.Context, page *proto.PageSize) (*proto.UserOnLineList, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var total []*proto.User
	var users []*proto.User

	c.ConnDb.Table("user").Find(&total)

	start := page.PSize * (page.Pn - 1)
	c.ConnDb.Table("user").Limit(int(page.PSize)).Offset(int(start)).Find(&users)
	return &proto.UserOnLineList{
		Data:  users,
		Total: int32(len(total)),
	}, nil
}

// LoadUnreadRecords 加载未读聊天记录
func (c *ChatServer) LoadUnreadRecords(ctx context.Context, user *proto.User) (*proto.RecordsResponses, error) {
	api.CheckAuth(ctx)
	c.mu.RLock()
	c.mu.RUnlock()

	var chat []*proto.ChatRecords
	var chatUpdate models.Chat
	c.ConnDb.Table("chat_records").Where("receiver_account = ? AND status = ?", user.Account, "未读").Find(&chat)

	// 改变状态为已读
	c.ConnDb.Table("chat_records").Model(&chatUpdate).Where("receiver_account = ? AND status = ?", user.Account, "未读").Update("status", "已读")
	return &proto.RecordsResponses{Data: chat}, nil
}

// SaveUnreadRecords 保存未读聊天记录
func (c *ChatServer) SaveUnreadRecords(ctx context.Context, connect *proto.Connect) (*emptypb.Empty, error) {
	api.CheckAuth(ctx)
	c.mu.Lock()
	defer c.mu.Unlock()

	senderUser := models.User{}
	receiverUser := models.User{}
	c.ConnDb.Table("user").Where("account = ?", connect.GetUser().Account).Find(&senderUser)
	c.ConnDb.Table("user").Where("account = ?", connect.ChattingWith[0]).Find(&receiverUser)
	chat := models.Chat{
		Sender:          senderUser.NickName,
		SenderAccount:   senderUser.Account,
		Receiver:        receiverUser.NickName,
		ReceiverAccount: receiverUser.Account,
		Message:         connect.Message,
		Status:          "未读",
	}

	c.ConnDb.Table("chat_records").Create(&chat)
	return &emptypb.Empty{}, nil
}

// ChattingRecords 获取聊天记录
func (c *ChatServer) ChattingRecords(ctx context.Context, connect *proto.Connect) (*proto.RecordsResponses, error) {
	api.CheckAuth(ctx)
	c.mu.RLock()
	c.mu.RUnlock()
	var chat []*proto.ChatRecords

	s := &api.Stack{
		Data: []*proto.ChatRecords{},
	}

	c.ConnDb.Order("created_at desc").Limit(10).Where(
		"sender_account = ? AND receiver_account = ? OR sender_account = ? AND receiver_account = ?",
		connect.GetUser().Account, connect.GetChattingWith()[0], connect.GetChattingWith()[0], connect.GetUser().Account).Find(&chat)

	// 倒叙聊天记录保存
	for _, v := range chat {
		s.Push(&proto.ChatRecords{
			Sender:   v.Sender,
			Receiver: v.Receiver,
			Message:  v.Message,
		})
	}
	return &proto.RecordsResponses{Data: s.Data}, nil
}

// UnRegisterClient 注销客户端
func (c *ChatServer) UnRegisterClient(ctx context.Context, user *proto.User) (*emptypb.Empty, error) {
	c.removeClient(user.Nickname)
	return &emptypb.Empty{}, nil
}

// 删除客户端用户
func (c *ChatServer) removeClient(nickname string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var user *models.User
	c.ConnDb.Table("user").Where("account = ?", nickname).Find(&user)

	fmt.Println("退出账号：", user.NickName, user.Account)

	delete(c.Connection, nickname)
	delete(c.Connection, user.NickName)
}

// GetOnLineUser 查看在线用户列表
// TODO 分页功能未实现，可以通过客户端传参控制取索引的长度实现
func (c *ChatServer) GetOnLineUser(ctx context.Context, empty *emptypb.Empty) (*proto.UserOnLineList, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var users []*proto.User
	for _, user := range c.Connection {
		users = append(users, &proto.User{
			Nickname: user.Nickname,
			Account:  user.Account,
		})
	}

	return &proto.UserOnLineList{
		Data: users,
	}, nil
}

// CreateChatStream 创建通道
func (c *ChatServer) CreateChatStream(server proto.ChatService_CreateChatStreamServer) error {
	for {
		req, err := server.Recv()
		if err == io.EOF {
			log.Println("no more data")
			break
		}
		if err != nil {
			return err
		}

		var user *models.User
		c.ConnDb.Table("user").Where("account = ?", req.GetUser().Account).Find(&user)
		fmt.Println(user.NickName, user.Account)

		var withUsers *models.User
		var chattingWiths []string

		for _, v := range req.GetChattingWith() {
			c.ConnDb.Table("user").Where("account = ?", v).Find(&withUsers)
			chattingWiths = append(chattingWiths, withUsers.NickName)
		}

		conn := &Connection{
			stream:       server,
			Nickname:     user.NickName,
			Account:      req.GetUser().Account,
			chattingWith: chattingWiths,
			active:       true,
			error:        make(chan error),
		}

		if _, ok := c.Connection[conn.Nickname]; ok {
			return errors.New("")
		} else {
			c.Connection[conn.Nickname] = conn
		}
		return <-conn.error
	}
	return nil
}

// BroadcastMessage 广播消息
func (c *ChatServer) BroadcastMessage(ctx context.Context, message *proto.Message) (*proto.User, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	var user *models.User
	c.ConnDb.Table("user").Where("account = ?", message.User.GetAccount()).Find(&user)

	for sendingTo, conn := range c.Connection {
		wait.Add(1)
		go func(sendingTo string, msg *proto.Message, conn *Connection) {
			defer wait.Done()

			senderConn := c.Connection[user.NickName]
			/*
				senderConn: 当前客户端用户信息
				conn：所有客户端的用户信息
			*/

			if conn.active && c.canChatWith(senderConn.chattingWith, conn.chattingWith, senderConn.Nickname, senderConn.Account, conn.Nickname, conn.Account, msg.Message) {
				grpcLog.Infof("sending message to %s: %v", sendingTo, conn.stream)
				if err := conn.stream.Send(&proto.Message{
					User: &proto.User{
						Nickname: user.NickName,
						Account:  msg.User.GetAccount(),
					},
					Message: msg.Message,
				}); err != nil {
					conn.active = false
					conn.error <- err
				}
			}
		}(sendingTo, message, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done

	return &proto.User{
		Nickname: user.NickName,
	}, nil
}

// 消息判断， 存入数据库
func (c *ChatServer) canChatWith(senderChattingWith, currentUserChattingWith []string, senderUserDisplayName, senderAccount, currentUserDisplayName, currentAccount, msg string) bool {
	// 退出聊天
	if msg == "quit" {
		return false
	}

	// TODO 处理对方客户端未登陆，聊天信息保存数据库
	fmt.Println("currentUserChattingWith", currentUserChattingWith[0])
	fmt.Println("senderChattingWith", senderChattingWith[0])
	fmt.Printf("当前客户端名子：%s, 目标客户端名字：%s\n", senderUserDisplayName, currentUserDisplayName)

	fmt.Println("打印登陆用户开始")

	// 自己发送的消息不要发送给自己
	if senderUserDisplayName == currentUserDisplayName {
		return false
	}

	if len(senderChattingWith) == 0 {
		return false
	}

	// TODO 消息同时传输给多个人
	if len(senderChattingWith) == 1 && strings.EqualFold(senderChattingWith[0], "all") {
		grpcLog.Infof("sender %s, current user %s, current user chatting with %v", senderUserDisplayName, currentUserDisplayName, currentUserChattingWith)
		if len(currentUserChattingWith) == 0 {
			return false
		} else if len(currentUserChattingWith) == 1 && strings.EqualFold(currentUserChattingWith[0], "all") {
			return true
		}
		for _, userName := range currentUserChattingWith {
			if senderUserDisplayName == userName {
				return true
			}
		}
		return false
	}

	fmt.Println("写入数据库，对方登陆")
	for _, userName := range senderChattingWith {
		if currentUserDisplayName == userName {
			// 聊天记录保存到数据库
			chat := models.Chat{
				Sender:          senderUserDisplayName,
				SenderAccount:   senderAccount,
				Receiver:        currentUserDisplayName,
				Message:         msg,
				ReceiverAccount: currentAccount,
				Status:          "已读",
			}
			c.ConnDb.Create(&chat)
			return true
		}
	}

	return false
}

// RegisterUser 注册用户
func (c *ChatServer) RegisterUser(ctx context.Context, info *proto.RegisterUserInfo) (*proto.User, error) {
	userAccount := models.User{}

	rand.Seed(time.Now().UnixNano())
	rankNum := rand.Intn(100000)

	// 判断账号是否已经存在
	c.ConnDb.Table("user").Where("account = ?", rankNum).Find(&userAccount)

	for i := 0; i <= 10; i++ {
		if userAccount.Account == "" {
			break
		} else if userAccount.Account == strconv.Itoa(rankNum) {
			rankNum = rand.Intn(10)
			continue
		}
		if i == 10 {
			fmt.Println("没有账号可以注册了，请升级服务")
			os.Exit(2)
		}
	}

	user := models.User{
		NickName: info.Nickname,
		PassWord: info.Password,
		Account:  strconv.Itoa(rankNum),
	}

	c.ConnDb.Create(&user)

	return &proto.User{
		Nickname: info.Nickname,
		Account:  strconv.Itoa(rankNum),
	}, nil
}

// Login 登陆用户
func (c *ChatServer) Login(ctx context.Context, info *proto.LoginInfo) (*proto.LoginResponse, error) {
	var user models.User
	obj := c.ConnDb.Where("account = ? AND password = ?", info.Account, info.Password).First(&user)
	if err := obj.Error; err != nil {
		return &proto.LoginResponse{IsLogin: false}, err
	}

	tokenString := api.CreateToken(info.Account)
	users := api.User{
		Nickname: user.NickName,
		Account:  user.Account,
		Password: user.PassWord,
		Token:    tokenString,
	}

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Println("marshal 出错：", err)
	}

	file, err := os.OpenFile(global.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666) // linux file permission settings
	if err != nil {
		fmt.Println("无法打开文件", global.FilePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(append(data, '\n'))

	return &proto.LoginResponse{IsLogin: true, Token: tokenString}, nil
}

// StartServer 启动server服务
func StartServer(port string, ctx context.Context) {
	go Crontab()

	conn, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := NewChatServer()

	if err := server.ConnDb.AutoMigrate(&models.User{}, &models.Chat{}); err != nil {
		log.Println("数据库迁移数据失败")
	}

	proto.RegisterChatServiceServer(s, server)
	go func() {
		select {
		case <-ctx.Done():
			s.Stop()
		}
	}()
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("chat:HbiyJloEvxevVCDp@tcp(39.97.251.109:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	//fmt.Println("连接数据库成功")

	return conn
}
