package internal

import (
	"github.com/ida-wong/leaf/module"
)

type Module struct {
	*module.Skeleton
}

func (module *Module) OnInit() {
	module.Skeleton = Skeleton
}

func (module *Module) OnDestroy() {

}
