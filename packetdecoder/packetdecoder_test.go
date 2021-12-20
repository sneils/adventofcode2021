package packetdecoder

import "testing"

func TestDecode(t *testing.T) {
	tests := map[string]string{
		"D2FE28":         "110100101111111000101000",
		"38006F45291200": "00111000000000000110111101000101001010010001001000000000",
		"EE00D40C823060": "11101110000000001101010000001100100000100011000001100000",
	}
	for input, expected := range tests {
		actual := Decode(input).Decoded
		if expected != actual {
			t.Errorf("Expected deceded message to look like %s, found %s.", expected, actual)
		}
	}
}

func TestDecodeVersion4(t *testing.T) {
	packet := Decode("D2FE28")
	exVersion, exTypeID, exValue, exJunk := 6, 4, 2021, "000"
	version, typeID, value, junk := packet.Header.Version, packet.Header.TypeID, packet.GetValue(), packet.Junk
	if exVersion != version {
		t.Errorf("Expected version to be %d, found %d.", exVersion, version)
	}
	if exTypeID != typeID {
		t.Errorf("Expected typeID to be %d, found %d.", exTypeID, typeID)
	}
	if exValue != value {
		t.Errorf("Expected value to be %d, found %d.", exValue, value)
	}
	if exJunk != junk {
		t.Errorf("Expected Junk to be %s, found %s.", exJunk, junk)
	}
}

func TestDecodeLengthType0(t *testing.T) {
	packet := Decode("38006F45291200")
	exVersion, exTypeID, exLengthTypeID, exLengthValue, exPackets := 1, 6, 0, 27, 2
	version, typeID, lengthTypeID, lengthValue, packets := packet.Header.Version, packet.Header.TypeID, packet.LengthTypeID, packet.LengthValue, len(packet.Packets)
	if exVersion != version {
		t.Errorf("Expected version to be %d, found %d.", exVersion, version)
	}
	if exTypeID != typeID {
		t.Errorf("Expected typeID to be %d, found %d.", exTypeID, typeID)
	}
	if exLengthTypeID != lengthTypeID {
		t.Errorf("Expected lengthTypeID to be %d, found %d.", exLengthTypeID, lengthTypeID)
	}
	if exLengthValue != lengthValue {
		t.Errorf("Expected lengthValue to be %d, found %d.", exLengthValue, lengthValue)
	}
	if exPackets != packets {
		t.Errorf("Expected to find %d packets, found %d.", exPackets, packets)
	}
}

func TestDecodeLengthType1(t *testing.T) {
	packet := Decode("EE00D40C823060")
	exVersion, exTypeID, exLengthTypeID, exLengthValue, exPackets := 7, 3, 1, 3, 3
	version, typeID, lengthTypeID, lengthValue, packets := packet.Header.Version, packet.Header.TypeID, packet.LengthTypeID, packet.LengthValue, len(packet.Packets)
	if exVersion != version {
		t.Errorf("Expected version to be %d, found %d.", exVersion, version)
	}
	if exTypeID != typeID {
		t.Errorf("Expected typeID to be %d, found %d.", exTypeID, typeID)
	}
	if exLengthTypeID != lengthTypeID {
		t.Errorf("Expected lengthTypeID to be %d, found %d.", exLengthTypeID, lengthTypeID)
	}
	if exLengthValue != lengthValue {
		t.Errorf("Expected lengthValue to be %d, found %d.", exLengthValue, lengthValue)
	}
	if exPackets != packets {
		t.Errorf("Expected to find %d packets, found %d.", exPackets, packets)
	}
}

func TestGetVersionSum(t *testing.T) {
	tests := map[string]int{
		"8A004A801A8002F478":             16,
		"620080001611562C8802118E34":     12,
		"C0015000016115A2E0802F182340":   23,
		"A0016C880162017C3686B18A3D4780": 31,
	}
	for input, expected := range tests {
		actual := Decode(input).GetVersionSum()
		if expected != actual {
			t.Errorf("Expected version sum to be %d, found %d.", expected, actual)
		}
	}
}

func TestGetValue(t *testing.T) {
	tests := map[string]int{
		"C200B40A82":                 3,
		"04005AC33890":               54,
		"880086C3E88112":             7,
		"CE00C43D881120":             9,
		"D8005AC2A8F0":               1,
		"F600BC2D8F":                 0,
		"9C005AC2F8F0":               0,
		"9C0141080250320F1802104A08": 1,
	}
	for input, expected := range tests {
		actual := Decode(input).GetValue()
		if expected != actual {
			t.Errorf("Expected value to be %d, found %d.", expected, actual)
		}
	}
}
