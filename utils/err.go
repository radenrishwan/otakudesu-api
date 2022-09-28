package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func PanicIfError(err error, w http.ResponseWriter, r *http.Request) {
	if err != nil {
		bytes, err := json.Marshal(DefaultResponse[string]{
			Code: 400,
			Data: err.Error(),
		})

		if err != nil {
			log.Fatalln(err)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(400)
		_, err = fmt.Fprint(w, string(bytes))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
