package sort

// Bubble Sort
func bsort(x []int) []int {
	var temp int

	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				temp = x[i]
				x[i] = x[j]
				x[j] = temp
			}
		}
	}
	return x
}
