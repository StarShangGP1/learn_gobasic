package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	_interface "learn_gobasic/homework/class29/interface"
	"learn_gobasic/pkg/apis"
	"log"
	"net/http"
	"strconv"
	"time"
)

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/learn_go"))
	if err != nil {
		log.Fatal("数据库连接失败：", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var circleServer _interface.ServerInterface = NewDbCircle(conn, NewCircleCache())

	if initCirlce, ok := circleServer.(_interface.CircleInitInterface); ok {
		if err := initCirlce.Init(); err != nil {
			log.Fatal("初始化失败", err)
		}
	}

	r := gin.Default()
	pprof.Register(r)

	r.POST("/post", func(c *gin.Context) {
		var cr *apis.Circle
		if err := c.BindJSON(&cr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "Unable to Post status" + err.Error(),
			})
			return
		}
		cr.ReleaseTime = time.Now().Unix()
		cr.Visible = true

		if err := circleServer.PostStatus(cr); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "Fail to post status",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.GET("/list", func(c *gin.Context) {
		if list, err := circleServer.GetList(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "Unable to get post",
			})
			return
		} else {
			c.JSON(http.StatusOK, list)
		}
	})

	r.DELETE("/delete/:pid", func(c *gin.Context) {

		idInString := c.Param("pid")
		id, err := strconv.Atoi(idInString)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "Invalid personal id",
			})
			return
		}

		if err := circleServer.DeleteStatus(uint32(id)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorMessage": "Unable to delete post under this personal id",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully deleted all posts under this Personal ID",
			})
		}

	})

	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
