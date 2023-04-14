package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chrislim1914/xrp-transaction/api/route"
	bootstrap "github.com/chrislim1914/xrp-transaction/bootstrap"
	"github.com/chrislim1914/xrp-transaction/database"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := bootstrap.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	database.ConnectDB(&config)

	server = gin.Default()
}

func main() {
	config, err := bootstrap.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	err = database.Migration(&config)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := route.Routes()
	port := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:         port,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
