package pfcpUdp

import (
	"net"

	"github.com/free5gc/pfcp"
)

type Message struct {
	RemoteAddr  *net.UDPAddr  `json:"remoteAddr"`
	PfcpMessage *pfcp.Message `json:"pfcpMessage"`
}

func NewMessage(remoteAddr *net.UDPAddr, pfcpMessage *pfcp.Message) (msg Message) {
	msg = Message{}
	msg.RemoteAddr = remoteAddr
	msg.PfcpMessage = pfcpMessage
	return
}
