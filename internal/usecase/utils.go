package usecase

import (
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pkg/errors"
)

func DurationFromMinutes(min int32) pgtype.Interval {
	return pgtype.Interval{
		Microseconds: int64(min) * 60 * 1000000,
		Valid:        true,
	}
}

func DateFromYear(result *pgtype.Date, year int) error {
	return result.Scan(time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC))
}

func DateFromString(result *pgtype.Date, strDate string) error {
	return DateFromISOString(result, strDate+"T00:00:00Z")
}

func DateFromISOString(result *pgtype.Date, isoDate string) error {
	var dateTime time.Time
	err := dateTime.UnmarshalText([]byte(isoDate))
	if err != nil {
		return err
	}
	return result.Scan(dateTime)
}

func ErrorIsNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}
