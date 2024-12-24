package utils

import "testing"

func TestTransTimeToDate(t *testing.T) {
	tm := TransTimeToDate("2024-12-18T20:37:39+08:00")
	t.Log(tm)
}
