package main

import (
	"log"
	"net/http"

	"github.com/bk-rim/job-service/adapters/database"
	"github.com/bk-rim/job-service/adapters/repository"
	"github.com/bk-rim/job-service/api"
	wsserver "github.com/bk-rim/job-service/api/ws_server"
	"github.com/bk-rim/job-service/domain"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan wsserver.BroadcastMessage)
	upgrader  = &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func main() {
	loadEnv()
	database.InitDB()
	jobRepo := repository.NewPostgresRepository(database.DB)
	jobExecuterRepo := repository.NewJobExecuterRepository()
	jobService := domain.NewJobService(jobRepo)
	jobExecuterService := domain.NewJobExecuterService(jobExecuterRepo)
	jobController := api.NewJobController(jobService)
	jobExecuterController := api.NewJobExecuterController(jobExecuterService)
	webSocketServer := wsserver.NewWebSocketServer(upgrader, clients, broadcast)

	router := gin.Default()

	configCors := cors.DefaultConfig()
	configCors.AllowOrigins = []string{"*"}
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	configCors.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	router.Use(cors.New(configCors))

	router.GET("/ws", webSocketServer.HandleWebSocket)

	go webSocketServer.BroadcastMessage()

	router.GET("/jobs", jobController.GetAll)
	router.GET("/jobs/:id", jobController.GetByID)
	router.POST("/job-execution", jobExecuterController.LaunchJobExecution)
	router.POST("/jobs", jobController.Create(func(messageType string, data any) {
		// broadcast <- wsserver.BroadcastMessage{MessageType: messageType, Data: data}
	}))
	router.PUT("/jobs/:id", jobController.Update(func(messageType string, data any) {
		broadcast <- wsserver.BroadcastMessage{MessageType: messageType, Data: data}
	}))
	router.DELETE("/jobs/:id", jobController.Delete)

	router.Run(":8080")

}
