package gorace

type GamePacket struct {
	// Packet header
	Header struct {
		// Packet type
		Type uint8 // 1 byte
		// Packet size
		Size uint16 // 2 bytes
	}
	// Packet data
	Data []byte
}

func (p *GamePacket) GetHeader() uint8 {
	return p.Header.Type
}

func (p *GamePacket) GetData() []byte {
	return p.Data
}

type GameLoginPacket struct {
	GamePacket
	// Packet data
	Data struct {
		// User name
		Username string
		// Password
		Password string
	}
}

func (p *GameLoginPacket) GetUsername() string {
	return p.Data.Username
}

func (p *GameLoginPacket) GetPassword() string {
	return p.Data.Password
}