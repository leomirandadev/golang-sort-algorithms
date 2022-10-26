package simple_sort

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Sortable interface {
	Number | string
}

func Sort[s Sortable](values []s, order int) {
	for i := range values {
		for j := i + 1; j < len(values); j++ {
			if order == 1 && values[i] > values[j] || order == -1 && values[i] < values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
}
