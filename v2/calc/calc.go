package calc

func Sub(a, b int) int {
	return a - b
}

func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}
