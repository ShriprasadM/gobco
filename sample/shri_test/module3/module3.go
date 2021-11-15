package module3

import "fmt"

func Module3(a, b bool) {
	fmt.Println("module1")
	if a || b {
		fmt.Println("inside if as a || b satsfied")
	} else {
		fmt.Println("Not inside if as a || b satsfied")
	}
}
