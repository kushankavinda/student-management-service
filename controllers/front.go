package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func RegisterController() {
	uc := implementController()
	http.Handle("/api/user/gen-key", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
