package pfcpType

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalNodeID(t *testing.T) {
	testData := NodeID{
		NodeIdType:  NodeIdTypeFqdn,
		NodeIdValue: []byte("free5gc.local"),
	}
	buf, err := testData.MarshalBinary()

	assert.Nil(t, err)
	assert.Equal(t, []byte{2, 102, 114, 101, 101, 53, 103, 99, 46, 108, 111, 99, 97, 108}, buf)
}

func TestJSONMarshalNodeID(t *testing.T) {
	testData := []struct {
		NodeID NodeID
		Result string
	}{
		{
			NodeID: NodeID{
				NodeIdType:  NodeIdTypeFqdn,
				NodeIdValue: []byte("free5gc.local"),
			},
			Result: `{"type":"Fqdn","value":"free5gc.local"}`,
		},
		{
			NodeID: NodeID{
				NodeIdType:  NodeIdTypeIpv4Address,
				NodeIdValue: net.IPv4(1, 2, 3, 4),
			},
			Result: `{"type":"IPv4","value":"1.2.3.4"}`,
		},
		{
			NodeID: NodeID{
				NodeIdType:  NodeIdTypeIpv6Address,
				NodeIdValue: net.IPv6zero,
			},
			Result: `{"type":"IPv6","value":"::"}`,
		},
	}

	for _, tt := range testData {
		buf, err := tt.NodeID.MarshalJSON()

		assert.Nil(t, err)
		assert.Equal(t, []byte(tt.Result), buf)
	}
}

func TestUnmarshalNodeID(t *testing.T) {
	buf := []byte{2, 102, 114, 101, 101, 53, 103, 99, 46, 108, 111, 99, 97, 108}
	var testData NodeID
	err := testData.UnmarshalBinary(buf)

	assert.Nil(t, err)
	expectData := NodeID{
		NodeIdType:  NodeIdTypeFqdn,
		NodeIdValue: []byte("free5gc.local"),
	}
	assert.Equal(t, expectData, testData)
}
