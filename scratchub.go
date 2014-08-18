package scratchgo

type Port interface {
	Action(*Message) error
}

type Scratchub struct {
	Ports []Port
}

func NewScratchub(ports []Port) *Scratchub {
	return &Scratchub{Ports: ports}
}

func (self *Scratchub) Action(msg *Message) error {
	for _, port := range self.Ports {
		if err := port.Action(msg); err != nil {
			return err
		}
	}
	return nil
}
