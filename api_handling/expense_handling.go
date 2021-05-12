package api_handling

import "net/http"

func Testapi(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(123)
}
