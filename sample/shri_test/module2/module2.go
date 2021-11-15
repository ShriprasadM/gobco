package module2

import "fmt"

func Module2(a, b bool) {
	fmt.Println("module2")
	if a || b {
		fmt.Println("inside if as a || b satsfied")
	} else {
		fmt.Println("Not inside if as a || b satsfied")
	}
}
