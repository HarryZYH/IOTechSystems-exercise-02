package main

import (
	"testing"
)

var devices = []Device{
	{
		Name:      "A-device",
		Type:      "ImaginaryType",
		Info:      "A Fake device uuid:c84738ea-c12d-19bc-8fea-0242ac120002, used for imaginary function",
		Value:     "MjA=",
		Timestamp: "1674573123",
	},
	{
		Name:      "B-device",
		Type:      "ImaginaryType",
		Info:      "A Fake device uuid:c9278-e57c-11ec-8fea-0242ac234242, used for imaginary function",
		Value:     "MTk5",
		Timestamp: "1674573555",
	},
	{
		Name:      "C-device",
		Type:      "ImaginaryType",
		Info:      "A Fake device uuid:c, used for imaginary function",
		Value:     "ODc=",
		Timestamp: "1714399922",
	},
}

func TestIsBefore(t *testing.T) {
	testCases := []struct {
		device Device
		want bool
	}{
		{devices[0], true},
		{devices[1], true},
		{devices[2], false},
	}

	for _, tc := range testCases {
		got := IsBefore(&tc.device)
		if got != tc.want {
			t.Errorf("got %t, want %t", got, tc.want)
		}
	}
}

func TestGetUuid(t *testing.T) {
	testCases := []struct {
		device Device
		want string
	}{
		{devices[0], "c84738ea-c12d-19bc-8fea-0242ac120002"},
		{devices[1], "c9278-e57c-11ec-8fea-0242ac234242"},
		{devices[2], "c"},
	}

	for _, tc := range testCases {
		got := GetUuid(&tc.device)
		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}

func TestDecodeValue(t *testing.T) {
	testCases := []struct {
		device Device
		want int
	}{
		{devices[0], 20},
		{devices[1], 199},
		{devices[2], 87},
	}

	for _, tc := range testCases {
		got := DecodeValue(&tc.device)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}
}
