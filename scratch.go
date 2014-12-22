package scratchgo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

type RspConn struct {
	net.Conn
}

func NewDefaultConnect() (*RspConn, error) {
	return NewConnect("127.0.0.1", 42001)
}

func NewConnect(host string, port int) (*RspConn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return &RspConn{conn}, nil
}

func (self *RspConn) SensorUpdate(value map[string]interface{}) error {
	var data string
	for k, v := range value {
		switch v := v.(type) {
		case int32, int64:
			data += fmt.Sprintf(` "%s" %d`, k, v)
		case float32, float64:
			data += fmt.Sprintf(` "%s" %f`, k, v)
		case bool:
			data += fmt.Sprintf(` "%s" %t`, k, v)
		case string:
			if strings.Index(v, " ") > 0 {
				v = fmt.Sprintf(` "%s"`, v)
			}
			data += fmt.Sprintf(` "%s" %s`, k, v)
		default:
			return fmt.Errorf("not support type [%v]", v)
		}
	}
	return self.send(fmt.Sprintf("sensor-update %s", data))
}

func (self *RspConn) BroadCast(value string) error {
	return self.send(fmt.Sprintf("broadcast %s", value))
}

func (self *RspConn) send(cmd string) error {
	var pkt []byte
	pkt = append(size_to_bytes(len(cmd)), []byte(cmd)...)
	_, err := self.Write(pkt)
	return err
}

func (self *RspConn) Recv() (*Message, error) {
	pkt := make([]byte, 4)
	_, err := self.Read(pkt)
	if err != nil {
		return nil, err
	}

	var pkt_len int32
	buf := bytes.NewBuffer(pkt)
	binary.Read(buf, binary.BigEndian, &pkt_len)

	data := make([]byte, pkt_len)
	_, err = self.Read(data)
	if err != nil {
		return nil, err
	}

	return ParseMessage(string(data)), nil
}

func size_to_bytes(size int) []byte {
	ret := make([]byte, 4)
	ret[0] = byte((size >> 24) & 0xFF)
	ret[1] = byte((size >> 16) & 0xFF)
	ret[2] = byte((size >> 8) & 0xFF)
	ret[3] = byte((size) & 0xFF)
	return ret
}
