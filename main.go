package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	interval := time.Second * 10 // Set the interval to 10 minutes

	for {
		fmt.Println("Starting auto-push process...")

		// Step 1: Check if there are changes
		if !hasChanges() {
			fmt.Println("No changes to commit. Skipping this cycle.")
			time.Sleep(interval)
			continue
		}

		// Step 2: Stage all changes
		err := runCommand("git", "add", ".")
		if err != nil {
			fmt.Println("Error staging files:", err)
			continue
		}

		// Step 3: Commit changes with a timestamp
		commitMessage := fmt.Sprintf("Auto-commit: %s", time.Now().Format(time.RFC1123))
		err = runCommand("git", "commit", "-m", commitMessage)
		if err != nil {
			fmt.Println("Error committing files:", err)
			continue
		}

		// Step 4: Push changes to GitHub
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

// Helper function to check for changes
func hasChanges() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error checking git status:", err)
		return false
	}

	return out.Len() > 0 // If there's output, there are changes
}
