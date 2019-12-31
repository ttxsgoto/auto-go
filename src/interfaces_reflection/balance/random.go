package balance

import (
	"errors"
	"math/rand"
)

type RoundRobinBalance struct {
}

func (p *RoundRobinBalance) DoBalance(insts []*Instance) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("No Instance")
		return
	}
	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}