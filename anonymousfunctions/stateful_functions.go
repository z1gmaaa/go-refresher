package main

func main() {
	ctr := count()
	println(ctr())
	println(ctr())
	println(ctr())
	println(ctr())
	println(ctr())
}

func count() func() int32 {
	x := 0

	return func() int32 {
		x++
		return int32(x)
	}
}
