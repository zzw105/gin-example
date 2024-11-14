package main

import (
	"gin-example/config"
	"gin-example/db"
	"gin-example/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	// 连接数据库
	db.Connect(cfg)
	defer db.Close()

	// 创建一个新用户
	// lastID, err := model.CreateUser("Alice")
	// if err != nil {
	// 	log.Fatalf("error creating user: %v", err)
	// }
	// fmt.Printf("Created user with ID: %d\n", lastID)

	// 获取所有用户
	// users, err := model.GetAllUsers()
	// if err != nil {
	// 	log.Fatalf("error getting users: %v", err)
	// }
	// fmt.Println("Users in database:")
	// for _, user := range users {
	// 	fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	// }

	r := gin.Default()

	// 设置信任的代理 IP 地址（例如信任本地或特定 IP 地址）
	err1 := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.0/24"})
	if err1 != nil {
		log.Fatalf("Error setting trusted proxies: %v", err1)
	}

	router.SetupRouter().Run(":8080")
}
