package day16_test

import (
	"testing"

	"github.com/jdhenke/advent-of-code/2021/day16"
	"github.com/jdhenke/advent-of-code/tester"
)

func TestPart1(t *testing.T) {
	tester.New(t, day16.Part1).Run(
		tester.FromString(`8A004A801A8002F478`).Want(16),
		tester.FromString(`620080001611562C8802118E34`).Want(12),
		tester.FromString(`C0015000016115A2E0802F182340`).Want(23),
		tester.FromString(`A0016C880162017C3686B18A3D4780`).Want(31),
		tester.FromFile("input.txt").Want(965),
	)
}

func TestPart2(t *testing.T) {
	tester.New(t, day16.Part2).Run(
		tester.FromString("C200B40A82").Want(3),
		tester.FromString("04005AC33890").Want(54),
		tester.FromString("880086C3E88112").Want(7),
		tester.FromString("CE00C43D881120").Want(9),
		tester.FromString("D8005AC2A8F0").Want(1),
		tester.FromString("F600BC2D8F").Want(0),
		tester.FromString("9C005AC2F8F0").Want(0),
		tester.FromString("9C0141080250320F1802104A08").Want(1),
	)
}
