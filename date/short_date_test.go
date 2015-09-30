package date

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

func TestShortDateMarshalText(t *testing.T) {
	expected := "2015-01-01"

	date := Date(2015, 1, 1, time.Local)

	var marshaled []byte
	var err error

	marshaled, err = date.MarshalText()

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	if expected != string(marshaled) {
		t.Logf("Expected marshaled value to equal\n%q\n\tgot:\n%q\n", expected, marshaled)
		t.Fail()
	}
}

func TestShortDateUnmarshalText(t *testing.T) {
	expected := Date(2015, 1, 1, time.Local)

	marshaled := "2015-01-01"

	date := ShortDate{}

	var err error

	err = date.UnmarshalText([]byte(marshaled))

	if err != nil {
		t.Logf("Expected no error, got %T:%v\n", err, err)
		t.Fail()
	}

	if !reflect.DeepEqual(expected, date) {
		t.Logf("Expected unmarshaled value to equal\n%q\n\tgot:\n%q\n", expected, date)
		t.Fail()
	}
}
