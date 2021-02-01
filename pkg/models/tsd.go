// github.com/bartmika/mulberry-server/internal/models/tsd.go
package models

import (
    "context"
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
    Create(ctx context.Context, uuid string, instrumentUuid string, value float64, timestamp time.Time, userUuid string) error
    ListAll(ctx context.Context) ([]*TimeSeriesDatum, error)
    FilterByUserUuid(ctx context.Context, userUuid string) ([]*TimeSeriesDatum, error)
    FindByUuid(ctx context.Context, uuid string) (*TimeSeriesDatum, error)
    DeleteByUuid(ctx context.Context, uuid string) error
    Save(ctx context.Context, datum *TimeSeriesDatum) error
}

type TimeSeriesDatumCreateRequest struct {
    InstrumentUuid string `json:"instrument_uuid"`
    Value float64 `json:"value,string"`
    Timestamp time.Time `json:"timestamp"`
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
}
