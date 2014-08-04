package scratch

import (
	"errors"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestMessge(t *testing.T) {
	Describe(t, "Test of Scratch Message class", func() {
		Context("test of parse", func() {
			It("type of broadcast", func() {
				msg, err := ParseMessage(`broadcast "sample"`)
				Expect(err).To(NotExist)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "broadcast")
				Expect(msg.GetNames()).To(Equal, []string{"sample"})
				Expect(msg.Get("sample")).To(NotExist)
			})
			It("type of sensor-update", func() {
				msg, err := ParseMessage(`sensor-update "name1" 1`)
				Expect(err).To(NotExist)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name1"})
				Expect(msg.Get("name1")).To(Equal, 1)
			})
			It("type of sensor-update", func() {
				msg, err := ParseMessage(`sensor-update "name2" 1.2`)
				Expect(err).To(NotExist)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name2"})
			})
			It("type of sensor-update", func() {
				msg, err := ParseMessage(`sensor-update "name3" true`)
				Expect(err).To(NotExist)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name3"})
				Expect(msg.Get("name3")).To(Equal, true)
			})
			It("type of sensor-update", func() {
				msg, err := ParseMessage(`sensor-update "name4" "aaa bbb"`)
				Expect(err).To(NotExist)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name4"})
				Expect(msg.Get("name4")).To(Equal, "aaa bbb")
			})
			It("type of sensor-update", func() {
				msg, err := ParseMessage(`sensor-update "name1" 1 "name2" 1.2 "name3" true "name4" "aaa bbb"`)
				Expect(err).To(NotExist)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name1", "name2", "name3", "name4"})
				Expect(msg.Get("name1")).To(Equal, 1)
				Expect(msg.Get("name3")).To(Equal, true)
				Expect(msg.Get("name4")).To(Equal, "aaa bbb")
			})
			It("unknown method", func() {
				msg, err := ParseMessage(`unknown "sample"`)
				Expect(msg).To(NotExist)
				Expect(err).To(Exist)
				Expect(err).To(Equal, errors.New("un supported type name.[unknown]"))
			})
		})
	})
}
