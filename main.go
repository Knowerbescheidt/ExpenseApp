package main

import (
	"log"
	"net/http"

	"backend_server/api_handling"
	"backend_server/db_handling"
	"backend_server/logger"
)

func testapi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(123)
}

func main() {

	logger.InitLogger()

	db_handling.Set_up_db()

	http.HandleFunc("/user", api_handling.Handle_user)

	http.HandleFunc("/transaction", api_handling.Testapi)

	http.HandleFunc("/test", testapi)

	log.Fatal(http.ListenAndServe(":8081", nil))

}
