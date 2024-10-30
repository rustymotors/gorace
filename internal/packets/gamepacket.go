package packets

import "fmt"

type NPSHeader struct {
	// message id
	Id uint16 // 2 bytes
	// Packet size
	Size uint16 // 2 bytes
}

type NPSPacket struct {
	// Packet header
	Header NPSHeader
	// Packet data
	Data []byte
}

type GamePacket struct {
	// Packet header
	Header struct {
		NPSHeader
		// message id
		Id uint16 // 2 bytes
		// Packet size
		Size uint16 // 2 bytes
		// version
		Version uint16 // 2 bytes
		// Reserved
		Reserved uint16 // 2 bytes
		// Checksum
		Checksum uint32 // 4 bytes
	}
	// Packet data
	Data []byte
}

func (p *GamePacket) MessageId() uint16 {
	return p.Header.Id
}

func (p *GamePacket) GetData() []byte {
	return p.Data
}

func (p *GamePacket) FromBytes(data []byte) {
	p.Header.Id = uint16(data[0]) | uint16(data[1])<<8
	p.Header.Size = uint16(data[2]) | uint16(data[3])<<8
	p.Header.Version = uint16(data[4]) | uint16(data[5])<<8
	p.Header.Reserved = uint16(data[6]) | uint16(data[7])<<8
	p.Header.Checksum = uint32(data[8]) | uint32(data[9])<<8 | uint32(data[10])<<16 | uint32(data[11])<<24
	p.Data = data[12:]
}

func (p *GamePacket) ToString() string {
	return fmt.Sprintf("Message ID: %d, Size: %d, Version: %d, Reserved: %d, Checksum: %d, Data: %s",
		p.Header.Id, p.Header.Size, p.Header.Version, p.Header.Reserved, p.Header.Checksum, p.Data)
}