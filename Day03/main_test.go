package main

import "testing"

func Test_checkRucksack(t *testing.T) {

	tests := []struct {
		name string
		line string
		want int
	}{
		// TODO: Add test cases.
		{name: "1", line: "vJrwpWtwJgWrhcsFMMfFFhFp", want: 16},
		{name: "2", line: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", want: 38},
		{name: "3", line: "PmmdzqPrVvPwwTWBwg", want: 42},
		{name: "4", line: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", want: 22},
		{name: "5", line: "ttgJtRGJQctTZtZT", want: 20},
		{name: "6", line: "CrZsJsPPZsGzwwsLwLmpwMDw", want: 19},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkRucksack(tt.line); got != tt.want {
				t.Errorf("checkRucksack() = %v, want %v", got, tt.want)
			}
		})
	}
}
