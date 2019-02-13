package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"mras/conf"
	"mras/router"
	"net/http"
	"time"
)

func main() {
	conf.Init()
	// 初始化gin
	gin.SetMode(viper.GetString("RunningMode"))
	g := gin.New()

	// 中间件列表
	middlewares := []gin.HandlerFunc{}

	// 加载路由
	router.Load(
		g,
		middlewares...,
	)

	// 启动ping自检服务
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("ServeAddr"))
	log.Printf(http.ListenAndServe(viper.GetString("ServeAddr"), g).Error())
}

// ping自检
func pingServer() error {
	for i := 0; i < viper.GetInt("MaxSelfPingTimes"); i++ {
		resp, err := http.Get("http://" + viper.GetString("ServeAddr") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
