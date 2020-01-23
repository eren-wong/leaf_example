package command

import "fmt"

func init() {
	skeleton.RegisterCommand("echo", "echo user inputs", onEcho)
}

func onEcho(args []interface{}) interface{} {
	return fmt.Sprintf("%v", args)
}
