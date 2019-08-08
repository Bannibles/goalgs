package main

import (
	"fmt"
	"testing"
)

func Test_mooresMajorityAlgorithm(t *testing.T) {
	tcases := []struct {
		name     string
		expMaj   bool
		majValue int
		data     []int
	}{
		{
			"majority_2",
			true,
			2,
			[]int{1, 2, 2},
		},
		{
			"majority_1_11",
			true,
			1,
			[]int{1, 1, 2},
		},
		{
			"majority_1_2",
			true,
			1,
			[]int{1, 2, 1},
		},
		{
			"majority_1_many",
			true,
			1,
			[]int{1, 3, 1, 1, 1, 2, 1, 1, 1, 2, 1, 1, 2, 1, 1, 1},
		},
		{
			"No_majority_1",
			false,
			-1,
			[]int{1, 2},
		},
		{
			"No_majority_2",
			false,
			-1,
			[]int{1, 2, 1, 2},
		},
		{
			"No_majority_3",
			false,
			-1,
			[]int{1, 2, 1, 2, 3},
		},
	}
	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			res := example5_13(tc.data)
			if res == -1 {
				if tc.expMaj {
					fmt.Printf("example5_13: expected majority %d, but got none\n", tc.majValue)
					t.Fail()
				}
			} else if tc.data[res] != tc.majValue {
				fmt.Printf("example5_13: expected majority %d, but got %d\n", tc.majValue, tc.data[res])
				t.Fail()
			}

			res = arrayMajorityWithMap(tc.data)
			if res == -1 {
				if tc.expMaj {
					fmt.Printf("arrayMajorityWithMap: expected majority %d, but got none\n", tc.majValue)
					t.Fail()
				}
			} else if tc.data[res] != tc.majValue {
				fmt.Printf("arrayMajorityWithMap: expected majority %d, but got %d\n", tc.majValue, tc.data[res])
				t.Fail()
			}

			res = mooresMajorityAlgorithm(tc.data)
			if res == -1 {
				if tc.expMaj {
					fmt.Printf("mooresMajorityAlgorithm: expected majority %d, but got none\n", tc.majValue)
					t.Fail()
				}
			} else if tc.data[res] != tc.majValue {
				fmt.Printf("mooresMajorityAlgorithm: expected majority %d, but got %d\n", tc.majValue, tc.data[res])
				t.Fail()
			}
		})
	}
}

func Test_linearSearchInArray(t *testing.T) {
	if linearSearchInArray([]int{}, 1) {
		t.Fail()
	}
	if !linearSearchInArray([]int{1}, 1) {
		t.Fail()
	}
	if !linearSearchInArray([]int{3, 2, 1}, 1) {
		t.Fail()
	}
	if linearSearchInArray([]int{3, 2, 1}, 4) {
		t.Fail()
	}
}
