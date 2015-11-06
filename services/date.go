package services

import "time"

// CustomDate is a custom type used to marshal/unmarshal date fields used in
// applications like SFDC
type CustomDate struct {
	time.Time
}

const customDateLayout = "2006-01-02"

// UnmarshalJSON implements the json.Unmarshaler interface for CustomDate
func (d *CustomDate) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	d.Time, err = time.Parse(customDateLayout, string(b))
	return
}

// MarshalJSON implements the json.Marshaler interface for CustomDate
func (d *CustomDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Time.Format(customDateLayout) + `"`), nil
}

var nilTime = (time.Time{}).UnixNano()

// IsSet can be used to check for CustomDate nil time
func (d *CustomDate) IsSet() bool {
	return d.UnixNano() != nilTime
}
