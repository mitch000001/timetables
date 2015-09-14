package timetables

import (
	"encoding/json"
	"strconv"
	"time"
)

type ShortDate struct {
	time.Time
}

func NewShortDate(date time.Time) ShortDate {
	return Date(date.Year(), date.Month(), date.Day(), date.Location())
}

func Date(year int, month time.Month, day int, location *time.Location) ShortDate {
	return ShortDate{time.Date(year, month, day, 0, 0, 0, 0, time.UTC)}
}

func (date *ShortDate) MarshalJSON() ([]byte, error) {
	if date.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(date.Format("2006-01-02"))
}

func (date *ShortDate) UnmarshalJSON(data []byte) error {
	unquotedData, _ := strconv.Unquote(string(data))
	time, err := time.Parse("2006-01-02", unquotedData)
	date.Time = time
	return err
}

func (date *ShortDate) String() string {
	return date.Format("2006-01-02")
}

func (s *ShortDate) MarshalText() ([]byte, error) {
	return []byte(s.Format("2006-01-02")), nil
}

func (s *ShortDate) UnmarshalText(text []byte) error {
	time, err := time.Parse("2006-01-02", string(text))
	if err != nil {
		return err
	}
	*s = ShortDate{time}
	return nil
}
