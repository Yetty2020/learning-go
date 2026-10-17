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
			t.Errorf("getStatus %1f = %q, expected %q", tc.percent, status, tc.wantStatus)
		}
	}
}
