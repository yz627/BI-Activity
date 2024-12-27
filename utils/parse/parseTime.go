package parse

import "time"

// TransTimeToDate 将时间字符串转换为 YYYY-MM-DD 格式
func TransTimeToDate(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		println(err)
		return ""
	}
	return t.Format(time.DateOnly)
}

// TransTimeToHour 将时间字符串转换为 HH:MM 格式
func TransTimeToHour(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return t.Format(time.TimeOnly)
}

func TransTimeToTime(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return t.Format(time.DateTime)
}
