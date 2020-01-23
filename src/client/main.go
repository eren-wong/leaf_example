package main

import (
	"github.com/ida-wong/leaf"

	"client/gate"
	"client/logic"
)

func main() {
	leaf.Run(
		logic.Module,
		gate.Module,
	)
}
