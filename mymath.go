package mymath

func Add(a, b, c int) int {
	return a + b + c
}

func Sub(a, b int) int {
	return a - b
}

func Div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
