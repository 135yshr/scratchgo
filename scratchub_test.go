package scratchgo

import (
	"fmt"
	. "github.com/r7kamura/gospel"
	"testing"
)

type ErrorNotExistPort struct{}

func (self ErrorNotExistPort) Action(msg *Message) error {
	return nil
}

type ErrorExistPort struct{}

func (self ErrorExistPort) Action(msg *Message) error {
	return fmt.Errorf("ErrorExistPort")
}

func TestScaratchub(t *testing.T) {
	Describe(t, "Scratchub TestCase", func() {
		Context("create hub", func() {
			It("create new hub.", func() {
				actual := Newhub([]Port{ErrorNotExistPort{}})
				Expect(actual).To(Exist)
			})
		})
		Context("run func Action.", func() {
			It("error not exist", func() {
				actual := Newhub([]Port{ErrorNotExistPort{}})
				err := actual.Action(&Message{})
				Expect(err).To(NotExist)
			})
			It("error exist", func() {
				actual := Newhub([]Port{ErrorExistPort{}, ErrorNotExistPort{}})
				err := actual.Action(&Message{})
				Expect(err).To(Exist)
			})
		})
	})
}
