package api_handling

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang/gddo/httputil/header"

	"backend_server/db_handling"
	"backend_server/entities"
	"backend_server/logger"
)

type User = entities.User

func extractUserFromRequest(w http.ResponseWriter, r *http.Request) (user *User, err error) {
	// found at https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var u User
	err = dec.Decode(&u)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			http.Error(w, msg, http.StatusBadRequest)

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			http.Error(w, msg, http.StatusBadRequest)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			http.Error(w, msg, http.StatusBadRequest)

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			http.Error(w, msg, http.StatusBadRequest)

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			http.Error(w, msg, http.StatusRequestEntityTooLarge)

		default:
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	user = &u
	fmt.Println(reflect.TypeOf(u))
	fmt.Println(u)
	return user, nil
}

func handlePostUser(w http.ResponseWriter, r *http.Request) error {
	//extract user from request and write it in a variable
	user, err := extractUserFromRequest(w, r)
	if err != nil {
		log.Fatal(err)
	}
	//take that instance and put it into the database
	err = db_handling.WriteUserToDb(user)
	if err != nil {
		http.Error(w, "Bad Request", 400)
	}

	return nil
}

func Handle_user(w http.ResponseWriter, r *http.Request) {

	logger.InfoLogger.Println("Handle User func is called with method", r.Method)

	// if method post call function to write user in db
	//TODO
	if r.Method == "POST" {

		err := handlePostUser(w, r)

		if err == nil {
			logger.InfoLogger.Println("Post method for handle user was successfull")
			w.Header().Set("Success", "Post Request Successfull")
			w.WriteHeader(200)
		}
		if err != nil {
			logger.WarningLogger.Println("Post method for handle user had an error")
		}
	}

	if r.Method == "GET" {

		handleGetUser(w, r)
	}

	if r.Method == "DELETE" {
		err := handleDeleteUser(w, r)
		if err == nil {
			w.Header().Set("Success", "Post Request Successfull")
			w.WriteHeader(200)
		}
	}

}

func handleGetUser(w http.ResponseWriter, r *http.Request) {

	userEmail := r.URL.Query()["Email"]

	userPointer, err := db_handling.GetUserByEmail(userEmail[0])
	if err != nil {
		http.Error(w, "For the Requested Resource nothing could be found", http.StatusNotFound)
		return
	}

	userToGet := *userPointer

	fmt.Println("user to get ist hier", userToGet)

	js, err := json.Marshal(userToGet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) (err error) {
	userEmailToDelete := r.URL.Query()["Email"]

	err = db_handling.DeleteUserByEmail(userEmailToDelete[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	return
}
