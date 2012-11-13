package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
)

type UUID struct {
	mostSigBits uint64
	lastSigBits uint64
}

func RandomUUID() (*UUID, error) {
	uuid := new(UUID)
	data := make([]byte, 16)

	n, err := rand.Read(data)
	if n != len(data) || err != nil {
		return nil, err
	}

	var msb, lsb uint64
	msb = (uint64(data[0]) & 0xFF) << 56
	msb |= (uint64(data[1]) & 0xFF) << 48
	msb |= (uint64(data[2]) & 0xFF) << 40
	msb |= (uint64(data[3]) & 0xFF) << 32
	msb |= (uint64(data[4]) & 0xFF) << 24
	msb |= (uint64(data[5]) & 0xFF) << 16
	msb |= (uint64(data[6]) & 0x0F) << 8
	msb |= uint64(0x4) << 12 // set the version to 4
	msb |= uint64(data[7]) & 0xFF
	lsb = (uint64(data[8]) & 0x3F) << 56
	lsb |= uint64(0x2) << 62 // set the variant to bits 01
	lsb |= (uint64(data[9]) & 0xFF) << 48
	lsb |= (uint64(data[10]) & 0xFF) << 40
	lsb |= (uint64(data[11]) & 0xFF) << 32
	lsb |= (uint64(data[12]) & 0xFF) << 24
	lsb |= (uint64(data[13]) & 0xFF) << 16
	lsb |= (uint64(data[14]) & 0xFF) << 8
	lsb |= uint64(data[15]) & 0xFF

	uuid.mostSigBits = msb
	uuid.lastSigBits = lsb

	return uuid, nil
}

func (uuid *UUID) String() string {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uuid.mostSigBits)
	binary.Write(buf, binary.BigEndian, uuid.lastSigBits)
	str := hex.EncodeToString(buf.Bytes())
	code := str[0:8] + "-" + str[8:12] + "-" + str[12:16] + "-" + str[16:20] + "-" + str[20:32]

	return code
}

func (uuid1 *UUID) Equal(uuid2 *UUID) bool {
	return (uuid1.mostSigBits == uuid2.mostSigBits) && (uuid1.lastSigBits == uuid2.lastSigBits)
}
