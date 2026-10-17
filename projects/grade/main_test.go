
package main

import "testing"

func TestGetGrade(t *testing.T ) {

	tests := [] struct{
		score int
		wantGrade string
		wantValid bool
	}{
		// boundaries: on it, and one below
		{70, "A", true},
		{69, "B", true},
		{60, "B", true},
		{59, "C", true},
		{50, "C", true},
		{49, "D", true},
		{45, "D", true},
		{44, "E", true},
		{40, "E", true},
		{39, "F", true},

		// ends of the valid range
		{0, "F", true},
		{100, "A", true},

		// invalid
		{-1, "", false},
		{101, "", false},
	}

	for _, tc := range tests {
		grade, valid := getGrade(tc.score)
		if grade != tc.wantGrade || valid != tc.wantValid {
			t.Errorf("getGrade(%d) = (%q, %v), expected (%q, %v)",
				tc.score, grade, valid, tc.wantGrade, tc.wantValid)
		}
	}

}


