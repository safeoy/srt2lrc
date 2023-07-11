package main

import (
	"reflect"
	"testing"
)

func TestParseSRT(t *testing.T) {
	data := []byte(`1
00:00:20,000 --> 00:00:24,400
Example subtitle

2
00:00:25,000 --> 00:00:29,400
Another example subtitle
`)
	expected := []Subtitle{
		{
			Index:     1,
			StartTime: "00:00:20,000",
			EndTime:   "00:00:24,400",
			Text:      "Example subtitle",
		},
		{
			Index:     2,
			StartTime: "00:00:25,000",
			EndTime:   "00:00:29,400",
			Text:      "Another example subtitle",
		},
	}

	result := parseSRT(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSrtTimeToLRCTime(t *testing.T) {
	testCases := []struct {
		srtTime string
		lrcTime string
	}{
		{"00:00:20,000", "[00:20.00]"},
		{"00:01:30,500", "[01:30.50]"},
	}

	for _, tc := range testCases {
		result := srtTimeToLRCTime(tc.srtTime)
		if result != tc.lrcTime {
			t.Errorf("Expected %s, got %s", tc.lrcTime, result)
		}
	}
}
