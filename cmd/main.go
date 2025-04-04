package main

import (
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/interfaces"
)

func main() {
	// بررسی اینکه آیا کاربر دستوری وارد کرده است یا نه
	if len(os.Args) < 2 {
		fmt.Println("لطفاً دستور را وارد کنید.")
		os.Exit(1)
	}

	// اجرای دستورات مختلف CLI
	interfaces.RunCLI()
}
