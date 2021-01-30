// github.com/bartmika/mulberry-server/internal/controllers/version.go
package controllers

import (
    "net/http"
)

func (h *BaseHandler) getVersion(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("Mulberry Server v1.0"))
}
