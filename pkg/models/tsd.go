// github.com/bartmika/mulberry-server/internal/models/tsd.go
package models

import (
    "time"
)

type TimeSeriesDatum struct {
    Uuid string `json:"uuid"`
    InstrumentUuid string `json:"instrument_uuid"`
    Value float64 `json:"value"`
    Timestamp time.Time `json:"timestamp"`
    UserUuid string `json:"user_uuid"`
}

type TimeSeriesDatumRepository interface {
    Create(uuid string, instrumentUuid string, value float64, timestamp time.Time, userUuid string) error
    ListAll() ([]*TimeSeriesDatum, error)
    FilterByUserUuid(userUuid string) ([]*TimeSeriesDatum, error)
    FindByUuid(uuid string) (*TimeSeriesDatum, error)
    DeleteByUuid(uuid string) error
    Save(datum *TimeSeriesDatum) error
}

type TimeSeriesDatumCreateRequest struct {
    InstrumentUuid string `json:"instrument_uuid"`
    Value float64 `json:"value,string"`
    Timestamp time.Time `json:"timestamp"`
    UserUuid string `json:"user_uuid"`
}

type TimeSeriesDatumCreateResponse struct {
    Uuid string `json:"uuid"`
    InstrumentUuid string `json:"instrument_uuid"`
    Value float64 `json:"value,string"`
    Timestamp time.Time `json:"timestamp"`
    UserUuid string `json:"user_uuid"`
}

type TimeSeriesDatumPutRequest struct {
    InstrumentUuid string `json:"instrument_uuid"`
    Value float64 `json:"value,string"`
    Timestamp time.Time `json:"timestamp"`
    UserUuid string `json:"user_uuid"`
}
