// github.com/bartmika/mulberry-server/internal/controllers/controller.go
package controllers

import (
    "net/http"
    "strings"

    "github.com/bartmika/mulberry-server/internal/repositories"
)

type BaseHandler struct {
    UserRepo *repositories.UserRepo
}

func NewBaseHandler(u *repositories.UserRepo) (*BaseHandler) {
    return &BaseHandler{
        UserRepo: u,
    }
}

func (h *BaseHandler) HandleRequests(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Split path into slash-separated parts, for example, path "/foo/bar"
    // gives p==["foo", "bar"] and path "/" gives p==[""]. Our API starts with
    // "/api/v1", as a result we will start the array slice at "3".
    p := strings.Split(r.URL.Path, "/")[3:]
    n := len(p)

    // fmt.Println(p, n) // For debugging purposes only.

    switch {
    case n == 1 && p[0] == "version" && r.Method == http.MethodGet:
        h.getVersion(w, r)
    case n == 1 && p[0] == "time-series-data" && r.Method == http.MethodGet:
        h.getTimeSeriesData(w, r)
    case n == 1 && p[0] == "time-series-data" && r.Method == http.MethodPost:
        h.postTimeSeriesData(w, r)
    case n == 2 && p[0] == "time-series-datum" && r.Method == http.MethodGet:
        h.getTimeSeriesDatum(w, r, p[1])
    case n == 2 && p[0] == "time-series-datum" && r.Method == http.MethodPut:
        h.putTimeSeriesDatum(w, r, p[1])
    case n == 2 && p[0] == "time-series-datum" && r.Method == http.MethodDelete:
        h.deleteTimeSeriesDatum(w, r, p[1])
    default:
        http.NotFound(w, r)
    }
}
