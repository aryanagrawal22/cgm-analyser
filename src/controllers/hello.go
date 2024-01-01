// controllers/root.go
package controllers

import (
    "net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World Again!"))
}