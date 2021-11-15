package sample

import (
	"fmt"
	"os"
	"testing"

	"github.com/junhwi/gobco"
	shritest "github.com/junhwi/gobco/sample/shri_test"
)

func TestMain(m *testing.M) {
	retCode := m.Run()
	gobco.ReportCoverage()
	gobco.ReportProfile("gobco.cover.out")
	os.Exit(retCode)
}

func TestFoo(t *testing.T) {
	if !Foo(9) {
		// t.Error("wrong")
	}

	// Foo(0)
}

func TestRecusive(t *testing.T) {
	fmt.Println("Inside TestRecursive")
	shritest.Recursive()
}
