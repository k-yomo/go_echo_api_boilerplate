package clock

import (
	"encoding/json"
	"gopkg.in/guregu/null.v3"
	"strings"
	"time"
)

// Now returns current time.
var Now = func() time.Time {
	return time.Now()
}

// RFC3339DATE represents date format
var RFC3339DATE = "2006-01-02"

// Date represents nullable date
type NullableDate null.Time

// ToNullTime returns null.Time
func (d NullableDate) ToNullTime() null.Time {
	return null.Time(d)
}

// MarshalJSON martial date to json string
func (d NullableDate) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(d.Time.Format(RFC3339DATE))
}

// MarshalJSON unmarshal json string to date
func (d *NullableDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(RFC3339DATE, s)
	if err != nil {
		*d = NullableDate{Time: t, Valid: false}
	} else {
		*d = NullableDate{Time: t, Valid: true}
	}
	return nil
}
