package anconsol

import (
	"fmt"

	"github.com/Aninetix/core/anware"
)

func (m *AnModule) HandleAnWare() {
	go func() {
		for {
			select {
			case <-m.ctx.Done():
				return
			case msg := <-m.in:
				m.anLogger.Info(fmt.Sprintf("[%v] => %#v\n", m.name, msg))
				m.msgChanAnWare(msg)
			}
		}
	}()
}

func (m *AnModule) msgChanAnWare(msg anware.AnWareEvent) {
	switch msg.Type {
	default:
		fmt.Printf("[%v] Message reçu: %+v\n", m.name, msg)
	}
}
