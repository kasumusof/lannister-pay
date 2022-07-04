package pkg

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonify(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func renderResponse(w http.ResponseWriter, data interface{}, status int) {
	jb, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	jsonify(w)
	w.WriteHeader(status)
	w.Write(jb)
}
