package main

func main() {

	// simple anonymous function
	func() {
		println("first anonymous function")
	}()

	simpleVarFun := func() {
		println("function as variables: example variable simpleVarFun")
	}
	// invoke function
	simpleVarFun()

	// anonmymous function with parameters
	greet := func(name string) {
		println("Hello", name)
	}
	greet("Jake")

	// anonymous function with return
	add := func(num1, num2 int32) int32 {
		return num1 + num2
	}
	sum := add(1, 2)
	println("sum = ", sum)
}
