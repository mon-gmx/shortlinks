package handlers

import (
    "fmt"
    "net/http"
)


func IsValidMethod(w http.ResponseWriter, r *http.Request, allowedMethods ...string) bool {
    for _, method := range allowedMethods {
        if r.Method == method {
            return true
        }
    }
    w.WriteHeader(http.StatusMethodNotAllowed)
    fmt.Fprintln(w, "Method Not Allowed")
    return false
}

