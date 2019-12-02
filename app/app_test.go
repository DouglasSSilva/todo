package app_test

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Starting tests")
	exitVal := m.Run()
	fmt.Println("Finishing tests")
	os.Exit(exitVal)

}
