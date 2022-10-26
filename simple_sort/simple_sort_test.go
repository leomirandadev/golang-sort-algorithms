package simple_sort

import (
	"sort-test/constants"
	"testing"
)

func TestSortInt(t *testing.T) {
	testCases := []struct {
		outputExpected []int
		input          []int
		order          int
	}{
		{
			outputExpected: []int{1, 2, 2, 44, 110},
			input:          []int{2, 1, 44, 2, 110},
			order:          1,
		},
		{
			outputExpected: []int{110, 44, 2, 2, 1},
			input:          []int{2, 1, 44, 2, 110},
			order:          -1,
		},
	}

	for _, testCase := range testCases {

		input := testCase.input

		Sort(input, testCase.order)

		for i, value := range input {
			if value != testCase.outputExpected[i] {
				t.Error("Error on sorting")
				break
			}
		}
	}
}

func TestSortString(t *testing.T) {

	testCases := []struct {
		outputExpected []string
		input          []string
		order          int
	}{
		{
			outputExpected: []string{"a", "b", "c", "g", "z"},
			input:          []string{"z", "c", "b", "g", "a"},
			order:          1,
		},
		{
			outputExpected: []string{"z", "g", "c", "b", "a"},
			input:          []string{"z", "c", "b", "g", "a"},
			order:          -1,
		},
	}

	for _, testCase := range testCases {

		input := testCase.input

		Sort(input, testCase.order)

		for i, value := range input {
			if value != testCase.outputExpected[i] {
				t.Error("Error on sorting")
				break
			}
		}
	}
}

func BenchmarkSimpleSort_biggerSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(constants.BiggerSliceToSort, 1)
	}
}

func BenchmarkSimpleSort_smallerSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort(constants.SmallerSliceToSort, 1)
	}
}
