package main

import "github.com/gin-gonic/gin"

func main() {
	go StartServer()

	r := gin.Default()

	r.GET("/ws", handleWebsocket)

	r.Run()
}
