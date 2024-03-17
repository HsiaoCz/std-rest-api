package receivers

import (
	"fmt"

	"github.com/anthdm/hollywood/actor"
)

type Borker struct {
	accountID int64
}

func NewBroker(accountID int64) actor.Producer {
	return func() actor.Receiver {
		return &Borker{
			accountID: accountID,
		}
	}
}

func (b *Borker) Receive(ctx *actor.Context) {
	
}
func (b *Borker) OnStart(ctx *actor.Context) {
	fmt.Println("[BROKER] started")
}
func (b *Borker) OnStop(ctx *actor.Context) {
	fmt.Println("[BROKER] stopped")
}
