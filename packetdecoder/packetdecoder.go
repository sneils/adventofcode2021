package packetdecoder

import (
	"fmt"

	"github.com/sneils/adventofcode2021/convert"
)

type Header struct {
	Version, TypeID int
}

func (header Header) ToBinary() string {
	return fmt.Sprintf("%03s%03s",
		convert.ToBinary(header.Version),
		convert.ToBinary(header.TypeID),
	)
}

type Packet struct {
	Decoded      string
	Header       Header
	Body         string
	Junk         string
	value        int
	LengthTypeID int
	LengthValue  int
	Packets      []*Packet
}

func (packet *Packet) Content() string {
	return packet.Header.ToBinary() + packet.Body
}

func ParseHeader(input string) Header {
	return Header{
		Version: convert.FromBinary(input[:3]),
		TypeID:  convert.FromBinary(input[3:6]),
	}
}

func Decode(input string) *Packet {
	decoded := ""

	for _, hex := range input {
		i := convert.FromHex(string(hex))
		bin := convert.ToBinary(i)
		padded := fmt.Sprintf("%04s", bin)
		decoded += padded
	}

	return parsePacket(decoded)
}

func parsePacket(decoded string) *Packet {
	header := ParseHeader(decoded)

	if header.TypeID == 4 {
		value, junk := parseLiteral(decoded[6:])

		return &Packet{
			Decoded: decoded,
			Header:  header,
			Body:    decoded[6 : len(decoded)-len(junk)],
			Junk:    junk,
			value:   value,
		}
	}

	lengthTypeID, lengthValue, packets, junk := parsePackets(decoded[6:])

	return &Packet{
		Decoded:      decoded,
		Header:       header,
		Body:         decoded[6 : len(decoded)-len(junk)],
		Junk:         junk,
		LengthTypeID: lengthTypeID,
		LengthValue:  lengthValue,
		Packets:      packets,
	}
}

func parseLiteral(content string) (int, string) {
	literal, i := "", 0
	for ; i < len(content); i += 5 {
		literal += content[i+1 : i+5]
		if content[i] == '0' {
			break
		}
	}
	return convert.FromBinary(literal), content[i+5:]
}

func parsePackets(content string) (int, int, []*Packet, string) {
	if content[0] == '0' {
		n := convert.FromBinary(content[1 : 1+15])
		packets, Junk := parseAllPackets(content[1+15:], n)
		return 0, n, packets, Junk
	}

	// content[0] == '1'
	n := convert.FromBinary(content[1 : 1+11])
	packets, Junk := parseNumPackets(content[1+11:], n)
	return 1, n, packets, Junk
}

func parseAllPackets(content string, n int) ([]*Packet, string) {
	packets := []*Packet{}
	for i := 0; i < n; {
		packet := parsePacket(content)
		packets = append(packets, packet)
		content = packet.Junk
		i += len(packet.Content())
	}
	return packets, content
}

func parseNumPackets(content string, n int) ([]*Packet, string) {
	packets := make([]*Packet, n)
	for i := 0; i < n; i++ {
		packets[i] = parsePacket(content)
		content = packets[i].Junk
	}
	return packets, content
}

func (packet *Packet) GetVersionSum() int {
	sum := packet.Header.Version
	for _, sub := range packet.Packets {
		sum += sub.GetVersionSum()
	}
	return sum
}

func (packet *Packet) GetValue() int {
	switch packet.Header.TypeID {
	case 0:
		sum := 0
		for _, sub := range packet.Packets {
			sum += sub.GetValue()
		}
		return sum
	case 1:
		prod := 1
		for _, sub := range packet.Packets {
			prod *= sub.GetValue()
		}
		return prod
	case 2:
		min := 0
		for _, sub := range packet.Packets {
			val := sub.GetValue()
			if min < 1 || min > val {
				min = val
			}
		}
		return min
	case 3:
		max := 0
		for _, sub := range packet.Packets {
			val := sub.GetValue()
			if max < val {
				max = val
			}
		}
		return max
	case 4:
		return packet.value
	case 5:
		if packet.Packets[0].GetValue() > packet.Packets[1].GetValue() {
			return 1
		}
		return 0
	case 6:
		if packet.Packets[0].GetValue() < packet.Packets[1].GetValue() {
			return 1
		}
		return 0
	case 7:
		if packet.Packets[0].GetValue() == packet.Packets[1].GetValue() {
			return 1
		}
		return 0
	}
	panic("Unknown typeID detected :(")
}
