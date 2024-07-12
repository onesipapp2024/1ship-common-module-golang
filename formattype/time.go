package formattype

import (
	"errors"
	"github.com/spf13/viper"
	"time"
)

type FormatTime time.Time
type FormatTimeString string

func (t *FormatTime) MarshalJSON() ([]byte, error) {
	location, err := time.LoadLocation(viper.GetString("CLIENT_TIMEZONE"))
	if err != nil {
		return nil, err
	}
	timeT := time.Time(*t)
	timeT = timeT.In(location)
	return timeT.MarshalJSON()
}

func (t *FormatTime) UnmarshalJSON(data []byte) error {
	timeT := time.Time(*t)
	err := timeT.UnmarshalJSON(data)
	if err != nil {
		return err
	}
	timeT = timeT.UTC()
	*t = FormatTime(timeT)
	return nil
}

func (t *FormatTimeString) MarshalJSON() ([]byte, error) {
	location, err := time.LoadLocation(viper.GetString("CLIENT_TIMEZONE"))
	if err != nil {
		return nil, err
	}
	timeT, err := time.Parse(time.DateTime, string(*t))
	if err != nil {
		return nil, err
	}
	timeT = timeT.In(location)
	tempT := timeT.Format(time.DateTime)
	return []byte("\"" + tempT + "\""), nil
}

func (t *FormatTimeString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == "\"\"" {
		return nil
	}
	if len(data) < 2 || data[0] != '"' || data[len(data)-1] != '"' {
		return errors.New("Time.UnmarshalJSON: input is not a JSON string")
	}
	data = data[len(`"`) : len(data)-len(`"`)]
	timeT, err := time.Parse(time.DateTime, string(data))
	if err != nil {
		return err
	}
	location, err := time.LoadLocation(viper.GetString("CLIENT_TIMEZONE"))
	if err != nil {
		return err
	}
	timeT = time.Date(
		timeT.Year(),
		timeT.Month(),
		timeT.Day(),
		timeT.Hour(),
		timeT.Minute(),
		timeT.Second(),
		timeT.Nanosecond(),
		location,
	)
	timeT = timeT.UTC()
	tempT := timeT.Format(time.DateTime)
	*t = FormatTimeString(tempT)
	return nil
}

func MakeFormatTimeString(t string) (*FormatTimeString, error) {
	if t == "" {
		return nil, nil
	}
	timeT, err := time.Parse(time.DateTime, t)
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocation(viper.GetString("CLIENT_TIMEZONE"))
	if err != nil {
		return nil, err
	}
	timeT = time.Date(
		timeT.Year(),
		timeT.Month(),
		timeT.Day(),
		timeT.Hour(),
		timeT.Minute(),
		timeT.Second(),
		timeT.Nanosecond(),
		location,
	)
	timeT = timeT.UTC()
	tempT := timeT.Format(time.DateTime)
	fts := FormatTimeString(tempT)
	return &fts, nil
}
