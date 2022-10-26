package insertion_sort

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Sortable interface {
	Number | string
}

func Sort[s Sortable](values []s, order int) {
	for i := range values {
		if i == 0 {
			continue
		}

		for j := i; j > 0 && ((order == 1 && values[j-1] > values[j]) || (order == -1 && values[j-1] < values[j])); j-- {
			values[j], values[j-1] = values[j-1], values[j]
		}
	}
}
