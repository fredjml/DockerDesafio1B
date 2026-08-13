package main

import "os"

const message = "Full Cycle Rocks!!\n"

func main() {
	if _, err := os.Stdout.WriteString(message); err != nil {
		os.Exit(1)
	}
}
