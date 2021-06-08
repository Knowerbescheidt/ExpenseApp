package api_handling

import (
	"backend_server/logger"
	"net/http"
)

func Testapi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(123)
}

func Handle_transaction(w http.ResponseWriter, r *http.Request) {

	logger.InfoLogger.Println("Handle Transaction func is called with method", r.Method)

	// if method post call function to write user in db
	//TODO
	if r.Method == "POST" {

		err := handlePostTransaction(w, r)

		if err == nil {
			logger.InfoLogger.Println("Post method for handle user was successfull")
			w.Header().Set("Success", "Post Request Successfull")
			w.WriteHeader(200)
		}
		if err != nil {
			logger.WarningLogger.Println("Post method for handle user had an error")
		}
	}

	// if r.Method == "GET" {

	// 	handleGetTransaction(w, r)
	// }

	// if r.Method == "DELETE" {
	// 	err := handleDeleteTransaction(w, r)
	// 	if err == nil {
	// 		w.Header().Set("Success", "Post Request Successfull")
	// 		w.WriteHeader(200)
	// 	}
	// }

}

func handlePostTransaction(w http.ResponseWriter, r *http.Request) (err error) {
	return nil
}
