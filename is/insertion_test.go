package is

import "testing"

//[26 36 89 80 87 66 86 49 56 14]
func TestSortShort(t *testing.T) {
	unsorted := []int{26, 36, 89, 80, 87, 66, 86, 49, 56, 14}
	correct := []int{14, 26, 36, 49, 56, 66, 80, 86, 87, 89}

	sortedData := InsertionSort(unsorted)

	if len(correct) != len(sortedData) {
		t.Error("sortedData is not of correct length")
	}

	for i := 0; i < len(sortedData); i++ {
		if correct[i] != sortedData[i] {
			t.Error("Unmatched data at index", i)
		}
	}

}
