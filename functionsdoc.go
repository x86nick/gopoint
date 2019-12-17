package main

import "fmt"

func f() {
	fmt.Println("simple function")
}
func fp(n int) {
	fmt.Println("function with parameter", n)
}
func fmulp(a, b int, s string) {
	fmt.Println("multiple parameter:", a, b, s)
}

/* function with arbitraroy nubmer of parameters, of same type, variadic function */
/* inside a function use a range loop to access individual parameters*/
/* only one of the parameter can be variadic and it must be the last one in paramenter list*/
func funarbi(s ...string) {
	for _, str := range s {
		fmt.Println(str + " ")
	}
	//	fmt.Println()
}

/* function with return value */
func freturn() int {
	return 42
}

/* function with return value, and caller will receive */
func freceive() (int, string, error) {
	return 0, "ok-receive", nil
}

/* you can give each return value a name
These names then defined within function bodies
*/
func funcReturnName() (n int, s string, err error) {
	n = 1
	s = "go-return"
	return n, s, nil
	// return /* you can return a blank, but not advised, you can fool a reader that function does not return anyting */
}

func main() {
	f()
	fp(123)
	/* when calling a function with parameter, pass the parameter in the order specified in parameter def */
	fmulp(1, 2, "go")
	/* */
	funarbi("hello", "this", "is", "arbitrarory", "function", "example")
	/* */
	fmt.Println(freturn())
	/* Receive the return values*/
	x, y, err := freceive()
	fmt.Println(x, y, err)
	/* */
	fmt.Println(funcReturnName())
}
