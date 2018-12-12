package sorting

func SelectionSort(data []int, ordered func(int, int) bool) []int  {

	for i := 0; i< len(data); i++ {

		for j := 1 + i; j < len(data); j ++ {

			if !ordered(data[i], data[j]) {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	return data
}
