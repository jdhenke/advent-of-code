package day16

import (
	"bytes"
	"io"
	"io/ioutil"
	"strconv"
)

type Packet struct {
	Version  int
	Type     int
	Value    int
	Children []*Packet
	Size     int // bytes
}

func Part1(r io.Reader) (ans int, err error) {
	text, err := toBinary(r)
	if err != nil {
		return 0, nil
	}
	packet := parsePacket(text)
	var traverse func(packet *Packet) int
	traverse = func(packet *Packet) int {
		out := packet.Version
		for _, c := range packet.Children {
			out += traverse(c)
		}
		return out
	}
	sumOfVersions := traverse(packet)
	return sumOfVersions, nil
}

func Part2(r io.Reader) (ans int, err error) {
	text, err := toBinary(r)
	if err != nil {
		return 0, nil
	}
	packet := parsePacket(text)
	return calculate(packet), nil
}

func calculate(packet *Packet) int {
	switch packet.Type {
	case 0:
		// sum
		sum := 0
		for _, c := range packet.Children {
			sum += calculate(c)
		}
		return sum
	case 1:
		// product
		product := 1
		for _, c := range packet.Children {
			product *= calculate(c)
		}
		return product
	case 2:
		// min
		min := calculate(packet.Children[0])
		for _, c := range packet.Children[1:] {
			if val := calculate(c); val < min {
				min = val
			}
		}
		return min
	case 3:
		// max
		max := calculate(packet.Children[0])
		for _, c := range packet.Children[1:] {
			if val := calculate(c); val > max {
				max = val
			}
		}
		return max
	case 4:
		return packet.Value
	case 5:
		// >
		if calculate(packet.Children[0]) > calculate(packet.Children[1]) {
			return 1
		}
		return 0
	case 6:
		// <
		if calculate(packet.Children[0]) < calculate(packet.Children[1]) {
			return 1
		}
		return 0
	case 7:
		// =
		if calculate(packet.Children[0]) == calculate(packet.Children[1]) {
			return 1
		}
		return 0
	}
	panic("bad input")
}

func toBinary(r io.Reader) (string, error) {
	hexText, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	for i := range hexText {
		_, _ = b.WriteString(map[string]string{
			"0": "0000",
			"1": "0001",
			"2": "0010",
			"3": "0011",
			"4": "0100",
			"5": "0101",
			"6": "0110",
			"7": "0111",
			"8": "1000",
			"9": "1001",
			"A": "1010",
			"B": "1011",
			"C": "1100",
			"D": "1101",
			"E": "1110",
			"F": "1111",
		}[string(hexText[i:i+1])])
	}
	text := string(b.Bytes())
	return text, nil
}

func parsePacket(text string) *Packet {
	packetVersion := num(text[0:3])
	packetType := num(text[3:6])
	var packetValue int
	var packetChildren []*Packet
	var packetSize int
	switch packetType {
	case 4:
		last := false
		literalBytes := ""
		var i int
		for i = 6; !last; i += 5 {
			if text[i:i+1] == "0" {
				last = true
			}
			literalBytes += text[i+1 : i+5]
		}
		packetValue = num(literalBytes)
		packetSize = i
	default:
		lengthTypeID := text[6:7]
		switch lengthTypeID {
		case "0":
			numBits := num(text[7 : 7+15])
			var childrenSize int
			for childrenSize = 0; childrenSize < numBits; {
				child := parsePacket(text[22+childrenSize:])
				packetChildren = append(packetChildren, child)
				childrenSize += child.Size
			}
			packetSize = 22 + childrenSize
		case "1":
			numPackets := num(text[7 : 7+11])
			var childrenSize int
			for i := 0; i < numPackets; i++ {
				child := parsePacket(text[18+childrenSize:])
				packetChildren = append(packetChildren, child)
				childrenSize += child.Size
			}
			packetSize = 18 + childrenSize
		}
	}
	return &Packet{
		Version:  packetVersion,
		Type:     packetType,
		Value:    packetValue,
		Children: packetChildren,
		Size:     packetSize,
	}
}

func num(s string) int {
	d, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(d)
}
