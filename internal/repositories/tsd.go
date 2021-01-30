package repositories

import (
    "encoding/json"
    "time"

	"github.com/sdomino/scribble"
	"github.com/bartmika/mulberry-server/pkg/models"
)

// TimeSeriesDatumRepo implements models.TimeSeriesDatumRepository
type TimeSeriesDatumRepo struct {
	db *scribble.Driver
}

func NewTimeSeriesDatumRepo(db *scribble.Driver) *TimeSeriesDatumRepo {
	return &TimeSeriesDatumRepo{
		db: db,
	}
}

func (r *TimeSeriesDatumRepo) Create(uuid string, instrumentUuid string, value float64, timestamp time.Time, userUuid string) error {
	tsd := models.TimeSeriesDatum{
		Uuid: uuid,
		InstrumentUuid: instrumentUuid,
		Value: value,
		Timestamp: timestamp,
        UserUuid: userUuid,
	}
	if err := r.db.Write("time_series_data", uuid, &tsd); err != nil {
		return err
	}
	return nil
}

func (r *TimeSeriesDatumRepo) ListAll() ([]*models.TimeSeriesDatum, error) {
    var results []*models.TimeSeriesDatum
    records, err := r.db.ReadAll("time_series_data")
	if err != nil {
        return nil, err
	}

    for _, f := range records {
		tsdFound := models.TimeSeriesDatum{}
        if err := json.Unmarshal([]byte(f), &tsdFound); err != nil {
            return nil, err
		}
		results = append(results, &tsdFound)
	}
	return results, nil
}

func (r *TimeSeriesDatumRepo) FilterByUserUuid(userUuid string) ([]*models.TimeSeriesDatum, error) {
    var results []*models.TimeSeriesDatum
    records, err := r.db.ReadAll("time_series_data")
	if err != nil {
        return nil, err
	}

    for _, f := range records {
		tsdFound := models.TimeSeriesDatum{}
        if err := json.Unmarshal([]byte(f), &tsdFound); err != nil {
            return nil, err
		}
		if tsdFound.UserUuid == userUuid {
			results = append(results, &tsdFound)
		}
	}
	return results, nil
}

func (r *TimeSeriesDatumRepo) FindByUuid(uuid string) (*models.TimeSeriesDatum, error) {
    tsd := models.TimeSeriesDatum{}
	if err := r.db.Read("time_series_data", uuid, &tsd); err != nil {
		return nil, err
	}
	return &tsd, nil
}


func (r *TimeSeriesDatumRepo) DeleteByUuid(uuid string) error {
    if err := r.db.Delete("time_series_data", uuid); err != nil {
		return err
	}
	return nil
}

func (r *TimeSeriesDatumRepo) Save(tsd *models.TimeSeriesDatum) error {
	if err := r.db.Write("time_series_data", tsd.Uuid, tsd); err != nil {
		return err
	}
	return nil
}
