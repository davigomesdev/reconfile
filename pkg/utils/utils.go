package utils

import (
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func ToInt(s string) int {
	if s == "" {
		return 0
	}
	n, _ := strconv.Atoi(s)
	return n
}

func ToFloat(s string) float64 {
	if s == "" {
		return 0
	}
	v, _ := strconv.ParseFloat(s, 64)
	return float64(v)
}

var excelBaseDate = time.Date(1899, 12, 30, 0, 0, 0, 0, time.UTC)

func ExcelDateToTime(excelSerial float64) time.Time {
	days := int(excelSerial)
	frac := excelSerial - float64(days)

	duration := time.Duration(frac * 24 * float64(time.Hour))
	return excelBaseDate.AddDate(0, 0, days).Add(duration)
}

func ToDate(s string) time.Time {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}
	}

	if f, err := strconv.ParseFloat(s, 64); err == nil {
		if f > 59 {
			return ExcelDateToTime(f)
		} else {
			return ExcelDateToTime(f)
		}
	}

	var t time.Time
	var err error

	switch {
	case len(s) == 8 && s[2] == '-' && s[5] == '-':
		t, err = time.Parse("02-01-06", s)
	case len(s) == 10 && s[2] == '/' && s[5] == '/':
		t, err = time.Parse("02/01/2006", s)
		if err != nil {
			t, err = time.Parse("01/02/2006", s)
		}
	case len(s) == 10 && s[2] == '-' && s[5] == '-':
		t, err = time.Parse("02-01-2006", s)
	case len(s) == 10 && s[4] == '-' && s[7] == '-':
		t, err = time.Parse("2006-01-02", s)
	case len(s) >= 20 && s[4] == '-' && s[10] == 'T':
		t, err = time.Parse(time.RFC3339, s)
	default:
		return time.Time{}
	}

	if err != nil {
		return time.Time{}
	}

	return t
}
func ToStringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ToMapPtr(s string) *map[string]string {
	if s == "" {
		return nil
	}
	var m map[string]string
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		return nil
	}
	return &m
}

func Safe(row []string, idx int) string {
	if idx < len(row) {
		return row[idx]
	}
	return ""
}
