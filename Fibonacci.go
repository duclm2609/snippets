func fibonacci(length int) {
	f, g := 0, 1
	for i := 0; i < length; i++ {
		println(f)
		f = f + g
		f = f - g
	}
}