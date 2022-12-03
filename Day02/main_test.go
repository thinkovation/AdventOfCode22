package main

import (
	"testing"
)

func TestCalculateScore(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{"1", "A X", 4},
		{"2", "A Y", 8},
		{"3", "A Z", 3},
		{"4", "B X", 1},
		{"5", "B Y", 5},
		{"6", "B Z", 9},
		{"7", "C X", 7},
		{"8", "C Y", 2},
		{"9", "C Z", 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateScore(tt.args); got != tt.want {
				t.Errorf("CalculateScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateScoreRevised(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{"1", "A X", 3},
		{"2", "A Y", 4},
		{"3", "A Z", 8},
		{"4", "B X", 1},
		{"5", "B Y", 5},
		{"6", "B Z", 9},
		{"7", "C X", 2},
		{"8", "C Y", 6},
		{"9", "C Z", 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateScoreRevised(tt.args); got != tt.want {
				t.Errorf("CalculateScoreRevised() = %v, want %v", got, tt.want)
			}
		})
	}
}
