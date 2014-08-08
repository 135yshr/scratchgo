package scratchgo

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestMessge(t *testing.T) {
	Describe(t, "Test of Scratch Message class", func() {
		Context("test of parse", func() {
			It("type of broadcast", func() {
				msg := ParseMessage(`broadcast "sample"`)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "broadcast")
				Expect(msg.Variables["command"]).To(Equal, "sample")
			})
			It("type of sensor-update", func() {
				msg := ParseMessage(`sensor-update "name1" 1`)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name1"})
				Expect(msg.Variables["name1"]).To(Equal, "1")
			})
			It("type of sensor-update", func() {
				msg := ParseMessage(`sensor-update "name2" 1.2`)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name2"})
			})
			It("type of sensor-update", func() {
				msg := ParseMessage(`sensor-update "name3" true`)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name3"})
				Expect(msg.Variables["name3"]).To(Equal, "true")
			})
			It("type of sensor-update", func() {
				msg := ParseMessage(`sensor-update "name4" "aaa bbb"`)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name4"})
				Expect(msg.Variables["name4"]).To(Equal, "aaa bbb")
			})
			It("type of sensor-update", func() {
				msg := ParseMessage(`sensor-update "name1" 1 "name2" 1.2 "name3" true "name4" "aaa bbb"`)
				Expect(msg).To(Exist)
				Expect(msg.Type).To(Equal, "sensor-update")
				Expect(msg.GetNames()).To(Equal, []string{"name1", "name2", "name3", "name4"})
				Expect(msg.Variables["name1"]).To(Equal, "1")
				Expect(msg.Variables["name3"]).To(Equal, "true")
				Expect(msg.Variables["name4"]).To(Equal, "aaa bbb")
			})
		})
	})
}
