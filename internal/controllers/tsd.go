// FILE LOCATION: github.com/bartmika/mulberry-server/internal/controllers/tsd.go
package controllers

import (
    "net/http"
)

func (h *BaseHandler) getTimeSeriesData(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("TODO: List Time Series Data")) //TODO: IMPLEMENT.
}

func (h *BaseHandler) postTimeSeriesData(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("TODO: Create Series Data")) //TODO: IMPLEMENT.
}

func (h *BaseHandler) getTimeSeriesDatum(w http.ResponseWriter, req *http.Request, uuid string) {
    w.Write([]byte("TODO: Get Series Datum with UUID: " + uuid)) //TODO: IMPLEMENT.
}

func (h *BaseHandler) putTimeSeriesDatum(w http.ResponseWriter, req *http.Request, uuid string) {
    w.Write([]byte("TODO: Update Series Datum with UUID: " + uuid)) //TODO: IMPLEMENT.
}

func (h *BaseHandler) deleteTimeSeriesDatum(w http.ResponseWriter, req *http.Request, uuid string) {
    w.Write([]byte("TODO: Delete Series Datum with UUID: " + uuid)) //TODO: IMPLEMENT.
}
