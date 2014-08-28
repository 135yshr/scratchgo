package scratchgo

type Port interface {
	Action(*Message) error
}

type Hub struct {
	Ports []Port
}

func NewHub(ports []Port) *Hub {
	return &Hub{Ports: ports}
}

func (self *Hub) Action(msg *Message) error {
	for _, port := range self.Ports {
		if err := port.Action(msg); err != nil {
			return err
		}
	}
	return nil
}
