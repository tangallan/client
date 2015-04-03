package engine

import (
	"github.com/keybase/client/go/libkb"
)

type ResetEngine struct {
	libkb.Contextified
}

func NewResetEngine() *ResetEngine {
	return &ResetEngine{}
}

func (e *ResetEngine) Name() string {
	return "Reset"
}

func (e *ResetEngine) GetPrereqs() EnginePrereqs {
	return EnginePrereqs{}
}

func (e *ResetEngine) RequiredUIs() []libkb.UIKind {
	return []libkb.UIKind{}
}

func (e *ResetEngine) SubConsumers() []libkb.UIConsumer {
	return []libkb.UIConsumer{}
}

func (e *ResetEngine) Run(ctx *Context) (err error) {
	e.SetGlobalContext(ctx.GlobalContext)

	if err = e.G().LoginState.Logout(); err != nil {
		return
	}
	if err = e.G().LocalDb.Nuke(); err != nil {
		return
	}
	return
}
