package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Alfred-Jijo/todoapp-ClI/cmd"
	"github.com/gonuts/commander"
)

const (
	todoFilename = ".todo"
)

func main() {
	filename := ""
	existCurTodo := false
	curDir, err := os.Getwd()
	if err == nil {
		filename = filepath.Join(curDir, todoFilename)
		_, err = os.Stat(filename)
		if err == nil {
			existCurTodo = true
		}
	}
	if !existCurTodo {
		home := os.Getenv("HOME")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		filename = filepath.Join(home, todoFilename)
	}
	command := &commander.Command{
		UsageLine: os.Args[0],
		Short:     "todo for cli",
	}
	command.Subcommands = []*commander.Command{
		cmd.MakeCmdAdd(filename),
		cmd.MakeCmdList(filename),
		cmd.MakeCmdDone(filename),
		cmd.MakeCmdUndone(filename),
		cmd.MakeCmdDelete(filename),
		cmd.MakeCmdUpdate(filename),
	}
	err = command.Dispatch(nil, os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
}
