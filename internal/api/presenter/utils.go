package presenter

import (
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func MicrosToMinutes(micros int64) int32 {
	return int32(micros / 1000000 / 60)
}

func MicrosToSeconds(micros int64) int32 {
	return int32(micros / 1000000)
}

func TimeToString(time pgtype.Time) string {
	seconds := MicrosToSeconds(time.Microseconds)
	minutes := seconds / 60
	hours := minutes / 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes%60, seconds%60)
}

func DateToString(date pgtype.Date) string {
	return date.Time.Format("2006-01-02")
}

func DateTimeToISO(date time.Time) string {
	return date.Format(time.RFC3339)
}
