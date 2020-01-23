package internal

import "github.com/ida-wong/leaf/gate"

type Module struct {
	*gate.Client
}

func (module *Module) OnInit() {
	module.Client = Client
}
