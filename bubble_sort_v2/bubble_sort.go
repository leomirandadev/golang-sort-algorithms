package bubble_sort_v2

type Number interface {
	int | int32 | int64 | float32 | float64
}

type Sortable interface {
	Number | string
}

func Sort[s Sortable](values []s, order int) {
	isSwapped := true
	for isSwapped {
		isSwapped = bubble(values, order)
	}
}

func bubble[s Sortable](values []s, order int) (isSwapped bool) {

	for i := len(values) - 1; i > 0; i-- {
		if order == 1 && values[i] < values[i-1] || order == -1 && values[i] > values[i-1] {
			values[i], values[i-1] = values[i-1], values[i]
			isSwapped = true
		}
	}

	return
}
