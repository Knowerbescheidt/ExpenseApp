package main

import (
	"backend_server/api_handling"
	"backend_server/db_handling"
	"backend_server/logger"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	logger.InitLogger()

	db_handling.Set_up_db()

	router.GET("user", api_handling.HandleGetUser)

	router.POST("user", api_handling.HandlePostUser)

	router.DELETE("user", api_handling.HandleDeleteUser)

	router.Run(":8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
