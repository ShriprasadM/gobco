package module2

import "fmt"

func module2(a, b bool) {
	fmt.Println("module2")
	if a || b {
		fmt.Println("inside if as a || b satsfied")
	}
}
