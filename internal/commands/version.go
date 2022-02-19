package commands

import (
	"fmt"
	"os"
)

func init() {
	JmCtl.Version = func() {
		fmt.Println("1.0.0")
		os.Exit(0)
	}
}
