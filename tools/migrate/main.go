package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/YuukiHayashi0510/todo-app/pkg/empty"
)

func main() {
	// フラグの設定
	name := flag.String("name", "", "Migration name (e.g., create_users)")
	flag.Parse()

	if empty.Is(*name) {
		fmt.Println("Error: name is required")
		fmt.Println("Usage: migrate-create -name <name>")
		os.Exit(1)
	}

	// 命名規則のチェック
	pattern := regexp.MustCompile(`^[a-z][a-z_]+[a-z]$`)
	if !pattern.MatchString(*name) {
		fmt.Println("Error: name must be snake_case and contain only lowercase letters")
		os.Exit(1)
	}

	// golang-migrateコマンドの実行
	cmd := exec.Command("migrate",
		"create",
		"-ext", "sql",
		"-dir", "db/migrations",
		"-seq",
		*name,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error executing migrate: %v\n%s\n", err, output)
		os.Exit(1)
	}

	fmt.Printf("Successfully created migration files: %s\n", output)
}
