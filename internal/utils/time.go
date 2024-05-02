package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	responsedto "github.com/quocanh1897/sample-gin-server/internal/model/dto/response"
)

func ConvertRFC3339StringToTime(text string) *time.Time {
	result, err := time.Parse(time.RFC3339, text)
	if err != nil {
		slog.Warn("[convertStringToRFC3339Format] Cannot parse time to RFC3339 format")
		return nil
	}
	return &result
}

func ConvertMiliSecondToTime(mSec int64) *time.Time {
	result := time.UnixMilli(mSec).UTC()
	return &result
}

func ConvertToISO8601(value string, format string) string {
	t, err := time.Parse(format, value)
	if err != nil {
		slog.Warn(fmt.Sprintf("[ConvertToISO8601] Cannot parse time from format %s to ISO8601", format))
		return ""
	}
	return t.Format(time.RFC3339)
}

func LoadListTimeZone(bytes []byte) ([]responsedto.SystemTimeZone, error) {
	var timezones []responsedto.TimeZone
	err := json.Unmarshal(bytes, &timezones)
	if err != nil {
		return []responsedto.SystemTimeZone{}, err
	}
	var result []responsedto.SystemTimeZone
	for _, timeZone := range timezones {
		result = append(result, responsedto.SystemTimeZone{
			Id:          timeZone.TimezoneId,
			Description: timeZone.WindowsDisplayName,
			WindowsId:   timeZone.WindowsId,
		})
	}
	return result, nil
}
