package shritest

import (
	"github.com/junhwi/gobco/sample/shri_test/module1"
	"github.com/junhwi/gobco/sample/shri_test/module2"
)

func Recursive() {
	module1.Module1(true, false)
	// module1.Module1(false, true)
	module2.Module2(false, true)
	// module2.Module2(false, true)
}
