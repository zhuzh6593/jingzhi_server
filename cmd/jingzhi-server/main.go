package main

import (
	"context"
	"os"

	"jingzhi-server/cmd/jingzhi-server/cmd"
)

func main() {
	command := cmd.RootCmd
	if err := command.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
