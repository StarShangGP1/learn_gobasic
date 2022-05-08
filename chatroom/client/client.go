package client

import (
	"bufio"
	"chat-project/api"
	"chat-project/global"
	"chat-project/proto"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

func init() {
	wait = &sync.WaitGroup{}
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

var grpcLog glog.LoggerV2
var client proto.ChatServiceClient
var wait *sync.WaitGroup

func Register(remoteServerHost, nickname, password string) {
	client = GrpcClientStart(remoteServerHost)
	req, err := client.RegisterUser(context.Background(), &proto.RegisterUserInfo{
		Nickname: nickname,
		Password: password,
	})
	if err != nil {
		fmt.Println("注册密码出错")
	}

	fmt.Printf("{Acccount: %s}\n", req.Account)
}

func Login(remoteServerHost, account, password string) {
	client = GrpcClientStart(remoteServerHost)
	req, err := client.Login(context.Background(), &proto.LoginInfo{
		Account:  account,
		Password: password,
	})
	if err != nil {
		fmt.Println("登陆错误")
	}
	if req.IsLogin != true {
		fmt.Println("登陆失败")
	} else {
		fmt.Println("登陆成功")
	}
}

func ChattingRecords(remoteServerHost, account string, whoDoYouWantToChatWithstring []string) {
	user := &proto.User{
		Account: account,
	}
	//client = GrpcClientStart(remoteServerHost)
	client = GrepClientStartTokenAuth(account, remoteServerHost)
	req, err := client.ChattingRecords(context.Background(), &proto.Connect{
		User:         user,
		Active:       false,
		ChattingWith: whoDoYouWantToChatWithstring,
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range req.Data {
		fmt.Printf("%s ————> %s： %s \n", v.Sender, v.Receiver, v.Message)
	}

	for _, v := range req.Data {
		fmt.Printf("<<<<<<欢迎%s|quit退出聊天>>>>>>\n", v.Sender)
		break
	}
	fmt.Printf("$ ")
}

// GrepClientStartTokenAuth TOKEN认证
func GrepClientStartTokenAuth(account, remoteServerHost string) proto.ChatServiceClient {
	user := api.User{}
	requestToken := new(api.AuthToekn)

	fp, err := os.OpenFile(global.FilePath, os.O_RDONLY, 0644)
	br := bufio.NewReader(fp)

	for {
		data, _, c := br.ReadLine()
		if c == io.EOF {
			fmt.Println("用户未登陆，请登录")
			os.Exit(1)
		}

		err = json.Unmarshal(data[:len(data)], &user)
		if err != nil {
			return nil
		}
		if account == user.Account {
			requestToken.Token = user.Token
			conn, err := grpc.Dial(fmt.Sprintf("%s", remoteServerHost), grpc.WithInsecure(), grpc.WithPerRPCCredentials(requestToken))
			if err != nil {
				panic(err)
			}
			return proto.NewChatServiceClient(conn)
		}
	}
	return nil
}

// GrpcClientStart 无token认证
func GrpcClientStart(remoteServerHost string) proto.ChatServiceClient {
	conn, err := grpc.Dial(fmt.Sprintf("%s", remoteServerHost), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return proto.NewChatServiceClient(conn)
}

func UnRegisterClient(nickname string) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client = proto.NewChatServiceClient(conn)
	_, err = client.UnRegisterClient(context.Background(), &proto.User{Nickname: nickname})
	if err != nil {
		return
	}
}

func ToUserExist(whoDoYouWantToChatWith []string, remoteServerHost string) bool {
	client = GrpcClientStart(remoteServerHost)
	req, err := client.GetAllUsers(context.Background(), &proto.PageSize{
		Pn:    1,
		PSize: 1000000000,
	})
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range req.Data {
		for _, u := range whoDoYouWantToChatWith {
			if v.Account == u {
				return true
			}
		}
	}
	return false
}

func StartWithClient(name, whoDoYouWantToChatWith, remoteServerHost string) {
	// 判断目标用户是否存在，不存在则退出
	to := ToUserExist(strings.Split(whoDoYouWantToChatWith, ","), remoteServerHost)
	if to != true {
		fmt.Println("目标用户未注册")
		os.Exit(3)
	}

	timestamp := time.Now()
	done := make(chan int)

	client = GrepClientStartTokenAuth(name, remoteServerHost)

	user := &proto.User{
		Account: name,
	}

	err := connect(name, user, strings.Split(whoDoYouWantToChatWith, ","), remoteServerHost)
	if err != nil {
		panic(err)
	}

	wait.Add(1)
	go func() {
		defer wait.Done()
		//mu := sync.Mutex{}
		running := true
		for running {
			// take input from command line
			var cmd string
			fmt.Printf("$ ")
			fmt.Scanf("%s", &cmd)

			if cmd == "send" {
				test := make(chan string, 2)
				rd := bufio.NewReader(os.Stdin)
				lineBuf, _, _ := rd.ReadLine()
				line := string(lineBuf)
				// initialize the message
				msg := &proto.Message{
					User:      user,
					Message:   line,
					Timestamp: timestamp.String(),
				}

				userinfo := strings.Split(whoDoYouWantToChatWith, ",")
				for _, temp := range userinfo {
					req, _ := client.GetOnLineUser(context.Background(), &emptypb.Empty{})
					for _, v := range req.Data {
						if temp != v.Account {
							// TODO 无需输出
						} else {
							test <- "ok"
							break
						}
					}
				}

				select {
				case <-test:
					info, _ := client.BroadcastMessage(context.Background(), msg)
					go quitClient(msg, info.Nickname)
				default:
					SaveUnreadRecords(strings.Split(whoDoYouWantToChatWith, ","), remoteServerHost, msg.Message, name)
					fmt.Println("对方用户未登陆，数据保存数据库")

				}
			}
			if cmd == "quit" {
				fmt.Println("注销账号，", name)
				UnRegisterClient(name)
				running = false
			}

			if cmd == "list" {
				GetClientList(remoteServerHost)
			}

			if cmd == "subscribe" {
				LoadUnreadRecords(name, remoteServerHost)
			}

			// TODO 平滑切换用户聊天 未实现
			if cmd == "switch" {
				rd := bufio.NewReader(os.Stdin)
				lineBuf, _, _ := rd.ReadLine()
				line := string(lineBuf)

				stream, _ := client.CreateChatStream(context.Background())
				_ = stream.Send(&proto.Connect{
					User:         user,
					Active:       true,
					ChattingWith: strings.Split(line, ","),
				})
				fmt.Println(strings.Split(line, ","))
				ChattingRecords(remoteServerHost, name, strings.Split(line, ","))
			}
		}
	}()

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
}

func GetUserList(remoteServerHost string, pn, pSize int32) {
	client = GrpcClientStart(remoteServerHost)
	req, err := client.GetAllUsers(context.Background(), &proto.PageSize{
		Pn:    pn,
		PSize: pSize,
	})
	if err != nil {
		return
	}
	fmt.Printf("目前聊天室共注册用户：%d\n", req.Total)

	for _, v := range req.Data {
		fmt.Printf("名称：%s  账号：%s\n", v.Nickname, v.Account)
	}
}

func GetClientList(remoteServerHost string) {
	client = GrpcClientStart(remoteServerHost)
	user, err := client.GetOnLineUser(context.Background(), &emptypb.Empty{})
	if err != nil {
		return
	}
	fmt.Println("获取在线用户列表")

	for _, v := range user.Data {
		fmt.Printf("名称：%s  账号：%s\n", v.Nickname, v.Account)
	}
}

func SaveUnreadRecords(whoDoYouWantToChatWith []string, remoteServerHost, msg, account string) {
	user := &proto.User{
		Account: account,
	}
	client = GrepClientStartTokenAuth(account, remoteServerHost)
	_, err := client.SaveUnreadRecords(context.Background(), &proto.Connect{
		User:         user,
		Active:       false,
		ChattingWith: whoDoYouWantToChatWith,
		Message:      msg,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func LoadUnreadRecords(account, remoteServerHost string) {
	client = GrepClientStartTokenAuth(account, remoteServerHost)
	req, err := client.LoadUnreadRecords(context.Background(), &proto.User{
		Account: account,
	})

	if err != nil {
		fmt.Println(err)
	}

	if len(req.Data) == 0 {
		fmt.Println("没有最新的消息")
	}
	// TODO 未实现位置调换
	for _, v := range req.Data {
		fmt.Printf("%s ————> %s： %s \n", v.Sender, v.Receiver, v.Message)
	}
}

func quitClient(msg *proto.Message, nickname string) {
	for {
		if msg.Message == "quit" {
			UnRegisterClient(nickname)
			os.Exit(1)
		}
	}
}

func connect(name string, user *proto.User, chattingWith []string, remoteServerHost string) error {
	var streamerror error

	stream, err := client.CreateChatStream(context.Background())
	go func() {
		defer wait.Done()
		_ = stream.Send(&proto.Connect{
			User:         user,
			Active:       true,
			ChattingWith: chattingWith,
		})
		ChattingRecords(remoteServerHost, name, chattingWith)
	}()

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	wait.Add(1)
	go func(str proto.ChatService_CreateChatStreamClient) {
		defer wait.Done()
		for {
			msg, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("error reading message: %v", err)
				break
			}

			grpcLog.Infof("%s : %s", msg.User.Nickname, msg.Message)
		}
	}(stream)
	return streamerror
}
