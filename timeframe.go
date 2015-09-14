package timetables

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func NewTimeframe(startYear int, startMonth time.Month, startDay int, endYear int, endMonth time.Month, endDay int, location *time.Location) Timeframe {
	return Timeframe{
		StartDate: Date(startYear, startMonth, startDay, location),
		EndDate:   Date(endYear, endMonth, endDay, location),
	}
}

type Timeframe struct {
	StartDate ShortDate
	EndDate   ShortDate
}

// TimeframeFromDate returns a Timeframe with the StartDate set to date and the EndDate set to today.
// The EndDate will use the same timezone location as provided in StartDate
func TimeframeFromDate(date ShortDate) Timeframe {
	endDate := NewShortDate(time.Now().In(date.Location()))
	return Timeframe{date, endDate}
}

func TimeframeFromQuery(params url.Values) (Timeframe, error) {
	from := params.Get("from")
	to := params.Get("to")
	if from == "" || to == "" {
		return Timeframe{}, fmt.Errorf("'from' and/or 'to' must be set")
	}
	startTime, err1 := time.Parse("20060102", from)
	startDate := ShortDate{startTime}
	endTime, err2 := time.Parse("20060102", to)
	endDate := ShortDate{endTime}
	if err1 != nil || err2 != nil {
		return Timeframe{}, fmt.Errorf("Malformed query params")
	}
	return Timeframe{StartDate: startDate, EndDate: endDate}, nil
}

func (tf Timeframe) IsInTimeframe(date ShortDate) bool {
	if tf.StartDate.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
		return true
	}
	if tf.EndDate.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
		return true
	}
	return tf.StartDate.Before(date.Time) && tf.EndDate.After(date.Time)
}

func (tf Timeframe) ToQuery() url.Values {
	params := make(url.Values)
	params.Set("from", tf.StartDate.Format("20060102"))
	params.Set("to", tf.EndDate.Format("20060102"))
	return params
}

func (tf Timeframe) MarshalJSON() ([]byte, error) {
	if tf.StartDate.IsZero() || tf.EndDate.IsZero() {
		return json.Marshal("")
	}
	return json.Marshal(fmt.Sprintf("%s,%s", tf.StartDate.Format("2006-01-02"), tf.EndDate.Format("2006-01-02")))
}

func (tf *Timeframe) UnmarshalJSON(data []byte) error {
	unquotedData, _ := strconv.Unquote(string(data))
	dates := strings.Split(unquotedData, ",")
	if len(dates) != 2 {
		*tf = Timeframe{}
		return nil
	}
	startTime, err1 := time.Parse("2006-01-02", dates[0])
	startDate := ShortDate{startTime}
	endTime, err2 := time.Parse("2006-01-02", dates[1])
	endDate := ShortDate{endTime}
	if err1 != nil || err2 != nil {
		*tf = Timeframe{}
		return nil
	}
	*tf = Timeframe{StartDate: startDate, EndDate: endDate}
	return nil
}

func (tf Timeframe) Days() int {
	return int(tf.EndDate.Add(24*time.Hour).Sub(tf.StartDate.Time) / time.Hour / 24)
}

func (tf Timeframe) IsZero() bool {
	return tf.StartDate.IsZero() && tf.EndDate.IsZero()
}

func (tf Timeframe) String() string {
	return fmt.Sprintf("{%s-%s}", tf.StartDate, tf.EndDate)
}

func (tf Timeframe) MarshalText() ([]byte, error) {
	marshaledStartDate, err := tf.StartDate.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("Error while marshaling StartDate: %v", err)
	}
	marshaledEndDate, err := tf.EndDate.MarshalText()
	if err != nil {
		return nil, fmt.Errorf("Error while marshaling EndDate: %v", err)
	}
	marshaled := fmt.Sprintf("{%s}:{%s}", marshaledStartDate, marshaledEndDate)
	return []byte(marshaled), nil
}

func (tf *Timeframe) UnmarshalText(value []byte) error {
	dates := bytes.SplitN(value, []byte(":"), 2)
	err := tf.StartDate.UnmarshalText(bytes.Trim(dates[0], "{}"))
	if err != nil {
		return fmt.Errorf("Error while unmarshaling StartDate: %v", err)
	}
	err = tf.EndDate.UnmarshalText(bytes.Trim(dates[1], "{}"))
	if err != nil {
		return fmt.Errorf("Error while unmarshaling EndDate: %v", err)
	}
	return nil
}
