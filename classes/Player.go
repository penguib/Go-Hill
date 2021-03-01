package classes

import (
	"net"
)

// Player type is the structure of the player
type Player struct {
	Socket *net.Conn

	NetID          uint32
	BrickCount     uint32
	UserID         uint32
	Username       string
	Admin          bool
	MembershipType uint8

	Position       Vector3
	Rotation       Vector3
	Scale          Vector3
	CameraPosition Vector3
	CameraRotation Vector3

	// Default 100
	MaxHealth float32
	Health    float32

	Alive bool

	Speed     uint32
	JumpPower uint32
	Score     int32
}

// SetScore sets the score of the player
func (p *Player) SetScore(score int32) {
	p.Score = score
}

// SetSpeed sets the speed of the player
func (p *Player) SetSpeed(speed uint32) {
	p.Speed = speed
}

// SetJumpPower sets the jump power of the player
func (p *Player) SetJumpPower(power uint32) {
	p.JumpPower = power
}
