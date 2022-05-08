package main

import "chat-project/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		return
	}
}
