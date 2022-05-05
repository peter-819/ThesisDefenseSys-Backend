package timex

import (
	"TDS-backend/common/errorx"
	"fmt"
	"time"
)

func TimeToString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.String()
}

func GStringToTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, t)
}

func FStringToTime(t string) (time.Time, error) {
	if t == "" {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, t)
}

func InTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}
func InEqTimeSpan(start, end, check time.Time) bool {
	return (check.After(start) || check.Equal(start)) && (check.Before(end) || check.Equal(end))
}
func IsIntersected(start, end, cstart, cend time.Time) bool {
	return InTimeSpan(start, end, cstart) || InTimeSpan(start, end, cend)
}
func IsIncluded(start, end, cstart, cend time.Time) bool {
	return InEqTimeSpan(start, end, cstart) && InEqTimeSpan(start, end, cend)
}

func IsIntersectedString(start, end, cstart, cend string) (bool, error) {
	_start, err := GStringToTime(start)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_end, err := GStringToTime(end)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_cstart, err := GStringToTime(cstart)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_cend, err := GStringToTime(cend)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}

	return InTimeSpan(_start, _end, _cstart) || InTimeSpan(_start, _end, _cend), nil
}
func IsIncludedString(start, end, cstart, cend string) (bool, error) {
	_start, err := GStringToTime(start)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_end, err := GStringToTime(end)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_cstart, err := GStringToTime(cstart)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_cend, err := GStringToTime(cend)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	fmt.Println("cstart: ", _cstart)
	fmt.Println("cend: ", _cend)
	fmt.Println("start: ", _start)
	fmt.Println("end: ", _end)
	fmt.Println("cstart eq: ", InEqTimeSpan(_start, _end, _cstart))
	fmt.Println("cend eq: ", InEqTimeSpan(_start, _end, _cend))
	return InEqTimeSpan(_start, _end, _cstart) && InEqTimeSpan(_start, _end, _cend), nil
}

func IsSameWeekString(a, b string) (bool, error) {
	_a, err := GStringToTime(a)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	_b, err := GStringToTime(b)
	if err != nil {
		return false, errorx.NewDefaultError("时间格式错误: " + err.Error())
	}
	ay, aw := _a.ISOWeek()
	by, bw := _b.ISOWeek()
	return ay == by && aw == bw, nil
}
