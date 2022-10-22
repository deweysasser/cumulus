package program

type RingArray struct {
	Values  []int
	Current int
}

func (r RingArray) Average() int {
	total := 0
	for _, v := range r.Values {
		total = total + v
	}

	return total / len(r.Values)
}

func (r *RingArray) Add(val int) {
	if len(r.Values) < cap(r.Values) {
		r.Values = append(r.Values, val)
	} else {
		r.Values[r.Current] = val
		r.Current++
		if r.Current > cap(r.Values) {
			r.Current = 0
		}
	}
}
