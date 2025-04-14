package is

func InsertionSort(data []int) []int {
	// todo: implementation
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			for j := i; j > 0; j-- {
				if data[j-1] > data[j] {
					data[j-1], data[j] = data[j], data[j-1]
				}
			}
		}
	}
	return data
}
