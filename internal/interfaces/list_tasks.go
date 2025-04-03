package interfaces

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/rezaabaskhanian/to_do-list-step-by-step/internal/infrastructure"
)

func ListTasks() {
	listTasksCmd := flag.NewFlagSet("list-tasks", flag.ExitOnError)
	listTasksCmd.Parse(os.Args[2:])

	tasks, err := infrastructure.LoadTasks()
	if err != nil {
		fmt.Println("خطا در بارگذاری تسک‌ها:", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("هیچ تسکی وجود ندارد.")
		return
	}

	fmt.Println("لیست تسک‌ها:")
	for _, task := range tasks {
		taskJSON, _ := json.MarshalIndent(task, "", "  ")
		fmt.Println(string(taskJSON))
	}
}
