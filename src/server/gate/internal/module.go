package internal

import "github.com/ida-wong/leaf/gate"

type Module struct {
	*gate.Server
}

func (module *Module) OnInit() {
	module.Server = Server
}
