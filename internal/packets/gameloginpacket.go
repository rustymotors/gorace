package packets

import "fmt"

type GameLoginPacket struct {
	GamePacket
	data struct {
		contexttoken        LengthPrefixedString
		unknown1            LengthPrefixedString
		encryptedSessionKey LengthPrefixedString
		gameId              LengthPrefixedString
	}
}

func (p *GameLoginPacket) Unmarshal() {
	offset := 0
	offset += p.data.contexttoken.Unmarshal(p.Data[offset:])
	offset += p.data.unknown1.Unmarshal(p.Data[offset:])
	offset += p.data.encryptedSessionKey.Unmarshal(p.Data[offset:])
	offset += p.data.gameId.Unmarshal(p.Data[offset:])
}

func (p *GameLoginPacket) FromGamePacket(packet GamePacket) {
	p.GamePacket = packet
	p.Unmarshal()
}

func (p *GameLoginPacket) ToString() string {
	return fmt.Sprintf("Login packet: %s, %s, %s, %s",
		p.data.contexttoken.ToString(), p.data.unknown1.ToString(), p.data.encryptedSessionKey.ToString(), p.data.gameId.ToString())
}
