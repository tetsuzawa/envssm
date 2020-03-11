package internal

import (
	"bufio"
	"fmt"
	"os"
)

func createFile(path string) (*os.File, error) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		sc := bufio.NewScanner(os.Stdin)
		fmt.Printf("overwrite %s? (y/n [y]): ", path)
		sc.Scan()
		if sc.Text() != "Y" && sc.Text() != "" {
			fmt.Println("Aborting...")
			os.Exit(1)
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return nil, fmt.Errorf("OpenFile: %w", err)
	}
	return f, nil
}
