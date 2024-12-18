package utils

import "time"

// TransTimeToDate 将时间字符串转换为 YYYY-MM-DD 格式
func TransTimeToDate(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		println(err)
		return ""
	}
	return t.Format("2006-01-02")
}

// TransTimeToHour 将时间字符串转换为 HH:MM 格式
func TransTimeToHour(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return t.Format("15:04")
}

func TransTimeToTime(timeStr string) string {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}
