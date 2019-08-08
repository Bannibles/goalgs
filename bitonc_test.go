package main

import (
	"testing"
)

func Test_BitonicSearch(t *testing.T) {
	// TODO: cases
	tcases := []struct {
		name     string // name of the test
		data     []int  // data to test
		value    int    // item to find
		expected int    // expected return value
	}{
		{
			"invalid length : less than 3",
			[]int{1, 2},
			2,
			-1,
		},
		{
			"mono increase 123",
			[]int{1, 2, 3},
			3,
			-1,
		},
		{
			"bitonic 1232",
			[]int{1, 2, 3, 2},
			3,
			2,
		},
		{
			"mono decrease 321",
			[]int{3, 2, 1},
			3,
			-1,
		},
		{
			"bitonic 1321",
			[]int{1, 3, 2, 1},
			3,
			1,
		},
		{
			"bitonic 131",
			[]int{1, 3, 1},
			1,
			0,
		},
		{
			"bitonic 231",
			[]int{2, 3, 1},
			1,
			2,
		},
	}
	for _, tc := range tcases {
		t.Logf("Running %s...", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			found := SearchBitonicArray(tc.data, tc.value)
			if tc.expected != found {
				t.Logf("Error in test: %s, extected %d, but returned %d", tc.name, tc.expected, found)
				t.Fail()
			}
		})
	}
}

func Test_isBitonicArray(t *testing.T) {
	// TODO: cases
	tcases := []struct {
		name     string // name of the test
		data     []int  // data to test
		expected bool   // expected return value
	}{
		{
			"bitonic 121",
			[]int{1, 2, 1},
			true,
		},
		{
			"bitonic 1321",
			[]int{1, 3, 2, 1},
			true,
		},
		{
			"bitonic 12",
			[]int{1, 2},
			false,
		},
		{
			"bitonic 123",
			[]int{1, 2, 3},
			false,
		},
		{
			"bitonic 321",
			[]int{3, 2, 1},
			false,
		},
	}
	for _, tc := range tcases {
		t.Logf("Running %s...", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			res := isBitonicArray(tc.data)
			if tc.expected != res {
				t.Logf("Error in test: %s, extected %t, but returned %t", tc.name, tc.expected, res)
				t.Fail()
			}
		})
	}
}
