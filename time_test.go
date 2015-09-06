package timetables

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestNewShortDate(t *testing.T) {
	dateTime := time.Date(2004, 02, 01, 5, 18, 47, 321, time.UTC)
	expectedDate := time.Date(2004, 02, 01, 0, 0, 0, 0, time.UTC)

	shortDate := NewShortDate(dateTime)

	if !reflect.DeepEqual(expectedDate, shortDate.Time) {
		t.Logf("Expected date to equal '%s', got '%s'\n", expectedDate, shortDate.Time)
		t.Fail()
	}

}

func TestDate(t *testing.T) {
	date := time.Date(2004, 02, 01, 0, 0, 0, 0, time.UTC)

	shortDate := Date(2004, 02, 01, time.UTC)

	if !reflect.DeepEqual(date, shortDate.Time) {
		t.Logf("Expected date to equal '%s', got '%s'\n", date, shortDate.Time)
		t.Fail()
	}
}

func TestShortDateUnmarshalJSON(t *testing.T) {
	testJson := `"2014-02-01"`

	var date ShortDate

	err := json.Unmarshal([]byte(testJson), &date)

	if err != nil {
		t.Logf("Expected error to be nil, got %T: %v\n", err, err)
		t.Fail()
	}

	if &date == nil {
		t.Logf("Expected date not to be nil\n")
		t.Fail()
	}

	expectedDate, err := time.Parse("2006-01-02", "2014-02-01")
	expectedShortDate := ShortDate{expectedDate}

	if err != nil {
		t.Logf("Expected error to be nil, got %T: %v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expectedShortDate, date) {
		t.Logf("Expected date to be '%+#v', got '%+#v'\n", expectedShortDate, date)
		t.Fail()
	}
}

func TestShortDateMarshalJSON(t *testing.T) {
	date := ShortDate{time.Date(2014, time.February, 01, 0, 0, 0, 0, time.UTC)}

	bytes, err := json.Marshal(&date)

	if err != nil {
		t.Logf("Expected error to be nil, got %T: %v\n", err, err)
		t.Fail()
	}

	expectedJson := `"2014-02-01"`

	if !reflect.DeepEqual(string(bytes), expectedJson) {
		t.Logf("Expected date to be '%s', got '%s'\n", expectedJson, string(bytes))
		t.Fail()
	}

	// Date is Zero
	date = ShortDate{}

	bytes, err = json.Marshal(&date)

	if err != nil {
		t.Logf("Expected error to be nil, got %T: %v\n", err, err)
		t.Fail()
	}

	expectedJson = `""`

	if !reflect.DeepEqual(string(bytes), expectedJson) {
		t.Logf("Expected date to be '%s', got '%s'\n", expectedJson, string(bytes))
		t.Fail()
	}
}

func TestFrom(t *testing.T) {
	now := time.Now()
	today := Date(now.Year(), now.Month(), now.Day(), now.Location())
	locations := []*time.Location{
		time.UTC,
		mustLoad(time.LoadLocation("America/New_York")),
		mustLoad(time.LoadLocation("Australia/Perth")),
	}

	var tests = []struct {
		startDate       ShortDate
		expectedEndDate ShortDate
	}{
		{
			startDate:       Date(2010, 02, 01, locations[0]),
			expectedEndDate: today,
		},
		{
			startDate:       Date(2010, 02, 01, locations[1]),
			expectedEndDate: Date(now.Year(), now.Month(), now.Day(), locations[1]),
		},
		{
			startDate:       Date(2010, 02, 01, locations[2]),
			expectedEndDate: Date(now.Year(), now.Month(), now.Day(), locations[2]),
		},
	}

	for _, test := range tests {
		actualEndDate := TimeframeFromDate(test.startDate).EndDate
		if !reflect.DeepEqual(test.expectedEndDate, actualEndDate) {
			t.Logf("Expected EndDate to equal '%s', got '%s'\n", test.expectedEndDate, actualEndDate)
			t.Fail()
		}
	}
}

func mustLoad(loc *time.Location, err error) *time.Location {
	if err != nil {
		panic(err)
	}
	return loc
}

func TestNewTimeframe(t *testing.T) {
	timeframe := NewTimeframe(2015, 1, 1, 2015, 2, 1, time.UTC)

	expectedTimeframe := Timeframe{
		StartDate: Date(2015, 1, 1, time.UTC),
		EndDate:   Date(2015, 2, 1, time.UTC),
	}

	if !reflect.DeepEqual(expectedTimeframe, timeframe) {
		t.Logf("Expected new timeframe to equal\n%s\n\tgot:\n%s\n", expectedTimeframe, timeframe)
		t.Fail()
	}
}

func TestTimeframeIsInTimeframe(t *testing.T) {
	tests := []struct {
		date          ShortDate
		timeframe     Timeframe
		isInTimeframe bool
	}{
		{
			Date(2015, 1, 3, time.Local),
			NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local),
			true,
		},
		{
			Date(2015, 3, 1, time.Local),
			NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local),
			false,
		},
		{
			Date(2015, 2, 1, time.Local),
			NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local),
			true,
		},
		{
			ShortDate{time.Date(2015, 2, 1, 23, 59, 59, 999, time.Local)},
			NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local),
			true,
		},
		{
			Date(2015, 1, 1, time.UTC),
			NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local),
			true,
		},
		{
			Date(2015, 2, 2, time.UTC),
			NewTimeframe(2015, 1, 1, 2015, 2, 1, time.Local),
			false,
		},
	}
	for _, test := range tests {
		if ok := test.timeframe.IsInTimeframe(test.date); ok != test.isInTimeframe {
			if ok {
				t.Logf("Expected date %q not to be in timeframe %q, but was\n", test.date, test.timeframe)
				t.Fail()
			} else {
				t.Logf("Expected date %q to be in timeframe %q, but was not\n", test.date, test.timeframe)
				t.Fail()
			}
		}
	}
}

func TestTimeframeMarshalJSON(t *testing.T) {
	startDate := ShortDate{time.Date(2014, time.February, 01, 0, 0, 0, 0, time.UTC)}
	endDate := ShortDate{time.Date(2014, time.April, 01, 0, 0, 0, 0, time.UTC)}

	var tests = []struct {
		timeframe    Timeframe
		expectedJson string
	}{
		{
			timeframe:    Timeframe{StartDate: startDate, EndDate: endDate},
			expectedJson: `"2014-02-01,2014-04-01"`,
		},
		{
			timeframe:    Timeframe{StartDate: startDate},
			expectedJson: `""`,
		},
		{
			timeframe:    Timeframe{EndDate: endDate},
			expectedJson: `""`,
		},
		{
			timeframe:    Timeframe{},
			expectedJson: `""`,
		},
	}

	for _, test := range tests {
		bytes, err := json.Marshal(&test.timeframe)
		if err != nil {
			t.Logf("Expected error to be nil, got %T: %v\n", err, err)
			t.Fail()
		}

		if !reflect.DeepEqual(string(bytes), test.expectedJson) {
			t.Logf("Expected date to be '%s', got '%s'\n", test.expectedJson, string(bytes))
			t.Fail()
		}
	}

}

func TestTimeframeUnmarshalJSON(t *testing.T) {
	startDate := ShortDate{time.Date(2014, time.February, 01, 0, 0, 0, 0, time.UTC)}
	endDate := ShortDate{time.Date(2014, time.April, 01, 0, 0, 0, 0, time.UTC)}

	var tests = []struct {
		testJson          string
		expectedTimeframe Timeframe
	}{
		{
			`"2014-02-01,2014-04-01"`,
			Timeframe{StartDate: startDate, EndDate: endDate},
		},
		{
			`"2014-02-01,"`,
			Timeframe{},
		},
		{
			`""`,
			Timeframe{},
		},
		{
			`","`,
			Timeframe{},
		},
		{
			`"2014-02-01,abcde"`,
			Timeframe{},
		},
		{
			`"abcde,2014-04-01"`,
			Timeframe{},
		},
		{
			`"abcde,abcde"`,
			Timeframe{},
		},
	}

	for _, test := range tests {
		var timeframe Timeframe
		err := json.Unmarshal([]byte(test.testJson), &timeframe)
		if err != nil {
			t.Logf("Expected error to be nil, got %T: %v\n", err, err)
			t.Fail()
		}

		if !reflect.DeepEqual(timeframe, test.expectedTimeframe) {
			t.Logf("Expected date to be '%+#v', got '%+#v'\n", test.expectedTimeframe, timeframe)
			t.Fail()
		}
	}
}

func TestTimeframeDays(t *testing.T) {
	tests := []struct {
		startDate ShortDate
		endDate   ShortDate
		days      int
	}{
		{
			Date(2015, 1, 1, time.Local),
			Date(2015, 1, 10, time.Local),
			10,
		},
		{
			Date(2015, 1, 1, time.Local),
			Date(2015, 1, 20, time.Local),
			20,
		},
	}

	for _, test := range tests {
		timeframe := Timeframe{test.startDate, test.endDate}

		days := timeframe.Days()

		if days != test.days {
			t.Logf("Expected result to equal %d, got %d\n", test.days, days)
			t.Fail()
		}
	}
}
