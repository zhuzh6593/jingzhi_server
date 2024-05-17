package main

import (
	"context"
	"os"

	"caict.ac.cn/llm-server/cmd/csghub-server/cmd"
)

func main() {
	command := cmd.RootCmd
	if err := command.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
