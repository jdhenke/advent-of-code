package day16

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestDecodeLiteral(t *testing.T) {
	text, err := toBinary(strings.NewReader("D2FE28"))
	require.NoError(t, err)
	packet := parsePacket(text)
	assert.Equal(t, Packet{
		Version:  6,
		Type:     4,
		Value:    2021,
		Children: nil,
		Size:     len(text) - 3,
	}, *packet)
}

func TestDecodeOperator(t *testing.T) {
	text, err := toBinary(strings.NewReader("38006F45291200"))
	require.NoError(t, err)
	packet := parsePacket(text)
	fmt.Println(packet)
	assert.Len(t, packet.Children, 2)
	assert.Equal(t, 10, packet.Children[0].Value)
	assert.Equal(t, 20, packet.Children[1].Value)
}
