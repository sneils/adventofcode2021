package days

import "github.com/sneils/adventofcode2021/packetdecoder"

func (day *Day) Day20() (int, int) {
	return 0, 0
}

func (day *Day) Day19() (int, int) {
	return 0, 0
}

func (day *Day) Day18() (int, int) {
	return 0, 0
}

func (day *Day) Day17() (int, int) {
	return 0, 0
}

func (day *Day) Day16() (int, int) {
	packet := packetdecoder.Decode(day.Inputs[0])
	return packet.GetVersionSum(), packet.GetValue()
}
