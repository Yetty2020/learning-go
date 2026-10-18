package main

import (
	"testing"

	
)

func TestCalcTip(t *testing.T){
	tests := [] struct{
		bill float64
		tipPercent float64
		tip float64
		wantTip bool

	} {

		// normal cases

{10, 10, 1, false},
{10, 0, 0,  false},



//invalid cases
{0, 10, 0, true},
{-1, 10,  0,true},
{10, -2, 0, true},
{-1, -1, 0,true},

// realistic case
{12000, 10, 1200, false},

	}

	for _, tc := range tests {

		tipAmount, err := calcTip(tc.bill, tc.tipPercent)
		gotTip := err != nil

		if tc.tip != tipAmount || gotTip != tc.wantTip{
			t.Errorf("calcTip(%.2f, %.1f) = (%.2f, %v), expected: (%.2f, %v)", tc.bill, tc.tipPercent, tipAmount, gotTip, tc.tipPercent, tc.wantTip)
		}


	}


}

func TestSplitBill(t *testing.T){
	tests := [] struct {
		total float64
		people int
		splittedBill float64
		wantError bool
	}{
		//normal case
{10, 1, 10, false},
{10, 2, 5, false},
{13200, 4, 3300, false},

// invalid people (good total)
{10, 0,  0, true},
{10, -1, 0, true},

// invalid total (good people)
{0,    10, 0, true},
{-100, 2,  0, true},

// both invalid
{0,  0,  0, true},
{-1, 0,  0, true},
{-1, -1, 0, true},

	}

	for _, tc := range tests{
		splitBill, err := calcSplitBill(tc.total, tc.people)
		gotError := err != nil
		if tc.splittedBill != splitBill || tc.wantError != gotError{
			t.Errorf("calcSplitBill(%.2f, %d) = (%.2f, %v), expected(%.2f, %v)", tc.total, tc.people, splitBill, gotError, tc.splittedBill, tc.wantError)
		}
	}
}






