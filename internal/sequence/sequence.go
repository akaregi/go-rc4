package sequence

func Sequence(max int) []int {
	a := make([]int, max)

	for i := range a {
		a[i] = i
	}

	return a
}
