package classes

import (
	"net"
)

type assets struct {
	Face uint32
	Hat1 uint32
	Hat2 uint32
	Hat3 uint32
}

type colors struct {
	Head     uint32
	Torso    uint32
	LeftArm  uint32
	RightArm uint32
	LeftLeg  uint32
	RightLeg uint32
}

// Player type is the structure of the player
type Player struct {
	Socket *net.Conn

	NetID          uint32
	BrickCount     uint32
	UserID         uint32
	Username       string
	Admin          uint8
	MembershipType uint8

	Position       Vector3
	Rotation       Vector3
	Scale          Vector3
	CameraPosition Vector3
	CameraRotation Vector3

	CameraFOV      uint32
	CameraDistance int32
	CameraType     string

	// Default 100
	MaxHealth float32
	Health    float32

	Alive bool

	Speed     uint32 // 4
	JumpPower uint32 // 5
	Score     int32
	Team      uint32

	Assets assets
	Colors colors

	Speech string
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

func (p *Player) SetBodyColor(color uint32) {
	p.Colors.Head = color
	p.Colors.Torso = color

	p.Colors.RightArm = color
	p.Colors.RightLeg = color

	p.Colors.LeftArm = color
	p.Colors.LeftLeg = color
}
