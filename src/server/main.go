package main

import (
	"github.com/ida-wong/leaf"

	"server/gate"
	"server/logic"
)

func main() {
	leaf.Run(
		logic.Module,
		gate.Module,
	)
}
