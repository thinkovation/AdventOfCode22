package main

import (
	"testing"
)

func Test_scanString(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		teststring string
		want       int
	}{
		// TODO: Add test cases.
		{name: "1", teststring: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 7},
		{name: "2", teststring: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 5},
		{name: "3", teststring: "nppdvjthqldpwncqszvftbrmjlhg", want: 6},
		{name: "4", teststring: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 10},
		{name: "4", teststring: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanString(tt.teststring); got != tt.want {
				t.Errorf("scanString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanStringForMessageMarker(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		teststring string
		want       int
	}{
		// TODO: Add test cases.
		{name: "1", teststring: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", want: 19},
		{name: "2", teststring: "bvwbjplbgvbhsrlpgdmjqwftvncz", want: 23},
		{name: "3", teststring: "nppdvjthqldpwncqszvftbrmjlhg", want: 23},
		{name: "4", teststring: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", want: 29},
		{name: "4", teststring: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", want: 26},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanStringForMessageMarker(tt.teststring); got != tt.want {
				t.Errorf("scanString() = %v, want %v", got, tt.want)
			}
		})
	}
}
