package test_demo

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Devision(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}
