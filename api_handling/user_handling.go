package api_handling

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"

	"backend_server/db_handling"
	"backend_server/entities"
)

type User = entities.User

func extractUserFromRequest(c *gin.Context) (u *User) {

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&u)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Bad Request",
		})
	}
	return u
}

func HandlePostUser(c *gin.Context) {
	user := extractUserFromRequest(c)

	err := db_handling.WriteUserToDb(user)
	if err == nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal Error Check the Logs",
		})
	}
	if err == nil {
		c.JSON(200, gin.H{"Request": "Success"})
	}
}

// func Handle_user(w http.ResponseWriter, r *http.Request) {

// 	logger.InfoLogger.Println("Handle User func is called with method", r.Method)

// 	// if method post call function to write user in db
// 	//TODO
// 	if r.Method == "POST" {

// 		err := HandlePostUser(w, r)

// 		if err == nil {
// 			logger.InfoLogger.Println("Post method for handle user was successfull")
// 			w.Header().Set("Success", "Post Request Successfull")
// 			w.WriteHeader(200)
// 		}
// 		if err != nil {
// 			logger.WarningLogger.Println("Post method for handle user had an error")
// 		}
// 	}

// 	// if r.Method == "GET" {

// 	// 	HandleGetUser(w, r)
// 	// }

// 	if r.Method == "DELETE" {
// 		err := handleDeleteUser(w, r)
// 		if err == nil {
// 			w.Header().Set("Success", "Post Request Successfull")
// 			w.WriteHeader(200)
// 		}
// 	}

// }

func HandleGetUser(c *gin.Context) {

	// There is a difference between Query Params user?Email and path params user/Email
	userEmail := c.Query("Email")

	userPointer, err := db_handling.GetUserByEmail(userEmail)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal Error Check the Logs",
		})
	}

	userToGet := *userPointer

	fmt.Println("user to get ist hier", userToGet)

	c.JSON(200, gin.H{"user": userToGet})
}

func HandleDeleteUser(c *gin.Context) {
	userEmailToDelete := c.Query("Email")

	err := db_handling.DeleteUserByEmail(userEmailToDelete)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Message": "Internal Error Check the Logs",
		})
	}
	if err == nil {
		c.JSON(200, gin.H{"Request": "Success"})
	}
}
