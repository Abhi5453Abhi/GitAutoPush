package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	interval := time.Second * 30 // Set the interval to 10 minutes

	for {
		fmt.Println("Starting auto-push process...")

		// Step 1: Stage all changess
		err := runCommand("git", "add", ".")
		if err != nil {
			fmt.Println("Error staging files:", err)
			continue
		}

		// Step 2: Commit changes with a timestamp
		commitMessage := fmt.Sprintf("Auto-commit: %s", time.Now().Format(time.RFC1123))
		err = runCommand("git", "commit", "-m", commitMessage)
		if err != nil {
			fmt.Println("Error committing files:", err)
			continue
		}

		// Step 3: Push changes to GitHub
		err = runCommand("git", "push", "origin", "main") // Replace "main" with your branch name if different
		if err != nil {
			fmt.Println("Error pushing files:", err)
			continue
		}

		fmt.Println("Changes pushed successfully!")
		time.Sleep(interval)
	}
}

// Helper function to run shell commands
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = nil
	cmd.Stderr = nil
	return cmd.Run()
}
