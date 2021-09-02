package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type loginInfo struct {
	UserName string `json:"user_name"`
	PassPort string `json:"pass_port"`
}

func setupRouter()  {
	router := gin.Default()
	router.GET("/demo/login", loginCheck)
	router.GET("/demo/ready", func(ctx *gin.Context) {
		ctx.String(200, "success")
	})
	_ = http.ListenAndServe(":8989", router)
}

func loginCheck(ctx *gin.Context) {
	inParams := &loginInfo{}
	if err := ctx.ShouldBindJSON(&inParams); err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(200, fmt.Sprintf("%s -> %s", inParams.UserName, inParams.PassPort))
}

func main()  {
	setupRouter()
}

