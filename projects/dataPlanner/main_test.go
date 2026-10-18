package main

import "testing"

func TestGetStatus(t *testing.T) {
	tests := []struct {
		percent    float64
		wantStatus string
	}{
		//boundaries / just below / over the plan
		{100.1, "Exhausted"},
		{100, "Exhausted"},
		{99.9, "Running low"},
		{80, "Running low"},
		{79.9, "OK"},
		{0, "OK"},
	}

	for _, tc := range tests {
		status := getStatus(tc.percent)
		if status != tc.wantStatus {
			t.Errorf("getStatus %.1f = %q, expected %q", tc.percent, status, tc.wantStatus)
		}
	}
}

func TestCalcUsage(t *testing.T) {
	tests := []struct {
		plan           float64
		used           float64
		wantRemaining  float64
		wantPercentage float64
		wantError      bool
	}{
		// normal usage
		{10, 7.5, 2.5, 75, false},
		{10, 0, 10, 0, false},
		{10, 10, 0, 100, false},

		// over the plan
		{10, 15, 0, 150, false},

		//invalid
		{10, -2, 0, 0, true},
		{0, 10, 0, 0, true},
		{-2, 10, 0, 0, true},
		{0, 0, 0, 0, true},
	}

	for _, tc := range tests {
		rem, perc, err := calcUsage(tc.plan, tc.used)
		gotError := err != nil
		if rem != tc.wantRemaining || perc != tc.wantPercentage || gotError != tc.wantError {
			t.Errorf("calcUsage (%.1f, %.1f) = (%.1f, %.1f, %v), expected(%.1f, %.1f, %v)", tc.plan, tc.used, rem, perc, gotError, tc.wantRemaining, tc.wantPercentage, tc.wantError)
		}
	}

}
