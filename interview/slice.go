package interview

func CopySlice(s []int) (old, new []int) {
	if s == nil {
		return nil, nil
	}
	old = s
	new = old[:]
	old = append(old, 1)
	new = append(new, 2)
	return old, new
}
