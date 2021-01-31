// FILE LOCATION: github.com/bartmika/mulberry-server/internal/controllers/tsd.go
package controllers

import (
    "fmt"
    "net/http"
    "encoding/json"

    "github.com/google/uuid"

    "github.com/bartmika/mulberry-server/pkg/models"
)

// To run this API, try running in your console:
// $ http get 127.0.0.1:5000/api/v1/time-series-data
func (h *BaseHandler) getTimeSeriesData(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()

    //TODO: Add filtering based on the authenticated user account. For now just list all the records.
    //      In a future article we will update this code.
    results, err := h.TsdRepo.ListAll(ctx)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Encode our results
    if err := json.NewEncoder(w).Encode(&results); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// To run this API, try running in your console:
// $ http post 127.0.0.1:5000/api/v1/time-series-data instrument_uuid="lalala" value="123" timestamp="2021-01-30T10:20:10.000Z" user_uuid="lalala"
func (h *BaseHandler) postTimeSeriesData(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    var requestData models.TimeSeriesDatumCreateRequest
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // For debugging purposes only.
    fmt.Println(requestData.InstrumentUuid)
    fmt.Println(requestData.Value)
    fmt.Println(requestData.Timestamp)
    fmt.Println(requestData.UserUuid)

    // Generate a `UUID` for our record.
    uid := uuid.New().String()

    // Save to our database.
    err := h.TsdRepo.Create(ctx, uid, requestData.InstrumentUuid, requestData.Value, requestData.Timestamp, requestData.UserUuid)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)

    // Return our record.
    responseData := models.TimeSeriesDatumCreateResponse{
        Uuid: uid,
        InstrumentUuid: requestData.InstrumentUuid,
        Value: requestData.Value,
        Timestamp: requestData.Timestamp,
        UserUuid: requestData.UserUuid,
    }
    if err := json.NewEncoder(w).Encode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// To run this API, try running in your console:
// $ http get 127.0.0.1:5000/api/v1/time-series-datum/f3e7b442-f3d4-4c2f-8f8d-d347982c1569
func (h *BaseHandler) getTimeSeriesDatum(w http.ResponseWriter, r *http.Request, uuid string) {
    ctx := r.Context()

    // Lookup our record.
    tsd, err := h.TsdRepo.FindByUuid(ctx, uuid)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return our record to the user as a response.
    if err := json.NewEncoder(w).Encode(&tsd); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// To run this API, try running in your console:
// $ http put 127.0.0.1:5000/api/v1/time-series-datum/f3e7b442-f3d4-4c2f-8f8d-d347982c1569 instrument_uuid="lalala" value="321" timestamp="2021-01-30T10:20:10.000Z" user_uuid="lalala"
func (h *BaseHandler) putTimeSeriesDatum(w http.ResponseWriter, r *http.Request, uid string) {
    ctx := r.Context()

    var requestData models.TimeSeriesDatumPutRequest
    if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // For debugging purposes only.
    fmt.Println(uid)
    fmt.Println(requestData.InstrumentUuid)
    fmt.Println(requestData.Value)
    fmt.Println(requestData.Timestamp)
    fmt.Println(requestData.UserUuid)

    // Update our record.
    tsd := models.TimeSeriesDatum{
        Uuid: uid,
        InstrumentUuid: requestData.InstrumentUuid,
        Value: requestData.Value,
        Timestamp: requestData.Timestamp,
        UserUuid: requestData.UserUuid,
    }
    err := h.TsdRepo.Save(ctx, &tsd)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return our record to the user as a response.
    if err := json.NewEncoder(w).Encode(&tsd); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// To run this API, try running in your console:
// $ http delete 127.0.0.1:5000/api/v1/time-series-datum/f3e7b442-f3d4-4c2f-8f8d-d347982c1569
func (h *BaseHandler) deleteTimeSeriesDatum(w http.ResponseWriter, r *http.Request, uid string) {
    ctx := r.Context()

    if err := h.TsdRepo.DeleteByUuid(ctx, uid); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.WriteHeader(http.StatusOK) // Note: https://tools.ietf.org/html/rfc7231#section-6.3.1
}
