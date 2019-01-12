package main

import (
	"fmt"
	"mb_gate/modbus"
	"net"
)

type Sender struct {
	conn net.Conn
	trId uint16
}

func NewSender() (s *Sender) {
	s = &Sender{}

	conn, err := net.Dial("tcp", "127.0.0.1:1502")
	if err != nil {
		panic("can't write")
	}

	s.conn = conn
	return
}

func (s *Sender) Send(pdu *modbus.ProtocolDataUnit) (ans *modbus.ProtocolDataUnit) {
	data := pdu.ToTCP(s.trId)
	s.trId++

	fmt.Println(pdu)
	s.conn.Write(data)
	res := make([]byte, 255)

	n, err := s.conn.Read(res)
	if err != nil {
		fmt.Printf("error %s\n", err.Error())
		return
	}

	_, ans, _ = modbus.FromTCP(res[:n])
	if ans != nil {
		fmt.Println(ans)
	}
	return
}

func main() {
	s := NewSender()
	defer s.conn.Close()

	pdu := modbus.ReadInputRegisters(1, 1, 1)
	s.Send(pdu)

	pdu = modbus.ReadInputRegisters(1, 2, 1)
	s.Send(pdu)

}
