package main

import (
	"Golang-Masterclass/simplebank/internal/initialize"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := initialize.Run()
	r.GET("/checkStatus", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Println("Starting server on:9090")
	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
