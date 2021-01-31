package repositories

import (
    "context"
    "database/sql"
    // "encoding/json"
    "time"

	"github.com/bartmika/mulberry-server/pkg/models"
)

// TimeSeriesDatumRepo implements models.TimeSeriesDatumRepository
type TimeSeriesDatumRepo struct {
	db *sql.DB
}

func NewTimeSeriesDatumRepo(db *sql.DB) *TimeSeriesDatumRepo {
	return &TimeSeriesDatumRepo{
		db: db,
	}
}

func (r *TimeSeriesDatumRepo) Create(ctx context.Context, uuid string, instrumentUuid string, value float64, timestamp time.Time, userUuid string) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "INSERT INTO time_series_data (uuid, instrument_uuid, value, timestamp, user_uuid) VALUES ($1, $2, $3, $4, $5)"

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		uuid,
		instrumentUuid,
		value,
		timestamp,
        userUuid,
	)
	return err
}

func (r *TimeSeriesDatumRepo) ListAll(ctx context.Context) ([]*models.TimeSeriesDatum, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

    query := "SELECT uuid, instrument_uuid, value, timestamp, user_uuid FROM time_series_data"

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
    defer rows.Close()

	var s []*models.TimeSeriesDatum
	for rows.Next() {
		m := new(models.TimeSeriesDatum)
		err = rows.Scan(
			&m.Uuid,
            &m.InstrumentUuid,
            &m.Value,
            &m.Timestamp,
            &m.UserUuid,
		)
		if err != nil {
			return nil, err
		}
		s = append(s, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return s, err
}

func (r *TimeSeriesDatumRepo) FilterByUserUuid(ctx context.Context, userUuid string) ([]*models.TimeSeriesDatum, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

    query := "SELECT uuid, instrument_uuid, value, timestamp, user_uuid FROM time_series_data WHERE user_uuid = $1"

	rows, err := r.db.QueryContext(ctx, query, userUuid)
	if err != nil {
		return nil, err
	}
    defer rows.Close()

	var s []*models.TimeSeriesDatum
	for rows.Next() {
		m := new(models.TimeSeriesDatum)
		err = rows.Scan(
			&m.Uuid,
            &m.InstrumentUuid,
            &m.Value,
            &m.Timestamp,
            &m.UserUuid,
		)
		if err != nil {
			return nil, err
		}
		s = append(s, m)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return s, err
}

func (r *TimeSeriesDatumRepo) FindByUuid(ctx context.Context, uuid string) (*models.TimeSeriesDatum, error) {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	m := new(models.TimeSeriesDatum)

	query := "SELECT uuid, instrument_uuid, value, timestamp, user_uuid FROM time_series_data WHERE uuid = $1"
	err := r.db.QueryRowContext(ctx, query, uuid).Scan(
        &m.Uuid,
        &m.InstrumentUuid,
        &m.Value,
        &m.Timestamp,
        &m.UserUuid,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that uuid.
		if err == sql.ErrNoRows {
			return nil, nil
		} else { // CASE 2 OF 2: All other errors.
			return nil, err
		}
	}
	return m, nil
}

func (r *TimeSeriesDatumRepo) DeleteByUuid(ctx context.Context, uuid string) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

    query := "DELETE FROM time_series_data WHERE uuid = $1;"

    _, err := r.db.Exec(query, uuid)
    if err != nil {
        return err
    }
    return nil
}

func (r *TimeSeriesDatumRepo) Save(ctx context.Context, m *models.TimeSeriesDatum) error {
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	query := "UPDATE time_series_data SET instrument_uuid = $1, value = $2, timestamp = $3, user_uuid = $4 WHERE uuid = $5"
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		m.InstrumentUuid,
		m.Value,
		m.Timestamp,
        m.UserUuid,
		m.Uuid,
	)
	return err
}
