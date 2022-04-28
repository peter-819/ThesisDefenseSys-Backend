package timex

import "time"

type JsonTime struct {
	time.Time
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = GStringToTime(string(data))
	return err
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	str := "\"" + TimeToString(t.Time) + "\""
	return []byte(str), nil
}
