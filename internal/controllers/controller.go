// github.com/bartmika/mulberry-server/internal/controllers/controller.go
package controllers

import (
    "net/http"

    "github.com/bartmika/mulberry-server/internal/repositories"
)

type BaseHandler struct {
    UserRepo *repositories.UserRepo
    TsdRepo *repositories.TimeSeriesDatumRepo
}

func NewBaseHandler(u *repositories.UserRepo, tsd *repositories.TimeSeriesDatumRepo) (*BaseHandler) {
    return &BaseHandler{
        UserRepo: u,
        TsdRepo: tsd,
    }
}

func (h *BaseHandler) HandleRequests(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Get our URL paths which are slash-seperated.
    ctx := r.Context()
    p := ctx.Value("url_split").([]string)
    n := len(p)

    // Get our authorization information.
    isAuthorized := ctx.Value("is_authorized").(bool)

    switch {
    case n == 1 && p[0] == "version" && r.Method == http.MethodGet:
        h.getVersion(w, r)
    case n == 1 && p[0] == "login" && r.Method == http.MethodPost:
        h.postLogin(w, r)
    case n == 1 && p[0] == "register" && r.Method == http.MethodPost:
        h.postRegister(w, r)
    case n == 1 && p[0] == "time-series-data" && r.Method == http.MethodGet:
        if isAuthorized {
            h.getTimeSeriesData(w, r)
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    case n == 1 && p[0] == "time-series-data" && r.Method == http.MethodPost:
        if isAuthorized {
            h.postTimeSeriesData(w, r)
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    case n == 2 && p[0] == "time-series-datum" && r.Method == http.MethodGet:
        if isAuthorized {
            h.getTimeSeriesDatum(w, r, p[1])
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    case n == 2 && p[0] == "time-series-datum" && r.Method == http.MethodPut:
        if isAuthorized {
            h.putTimeSeriesDatum(w, r, p[1])
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    case n == 2 && p[0] == "time-series-datum" && r.Method == http.MethodDelete:
        if isAuthorized {
            h.deleteTimeSeriesDatum(w, r, p[1])
        } else {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
        }
    default:
        http.NotFound(w, r)
    }
}
