package anconsol

import (
	"context"

	"github.com/Aninetix/aninetix-app-template/anparam"
	"github.com/Aninetix/core/aninterface"
	"github.com/Aninetix/core/anware"
)

type AnModule struct {
	name string

	ctx context.Context
	in  <-chan anware.AnWareEvent
	mw  *anware.AnWare

	anLocal  aninterface.StaticData
	anLogger aninterface.AnLogger

	Flags  *anparam.Flags
	Config *anparam.Config
}

func init() {
	anware.RegisterModule[*anparam.Flags, *anparam.Config](
		"anConsol",
		NewAnModule[*anparam.Flags, *anparam.Config],
	)
}

func NewAnModule[F *anparam.Flags, C *anparam.Config](
	local aninterface.StaticData,
	config C,
	flags F,
	logger aninterface.AnLogger,
) anware.AnModule {
	return &AnModule{
		anLocal:  local,
		anLogger: logger,
		Flags:    flags,
		Config:   config,
	}
}

func (m *AnModule) Name() string {
	return m.name
}

func (m *AnModule) Param(ctx context.Context, in <-chan anware.AnWareEvent, mw *anware.AnWare) {
	m.ctx = ctx
	m.in = in
	m.mw = mw
	m.name = "anConsol"
}

func (m *AnModule) Stop() error {
	return nil
}

func (m *AnModule) Start() {
	m.HandleAnWare()
	m.ProcessModule()
}
