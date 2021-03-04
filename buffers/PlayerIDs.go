package buffers

import (
	"Go-Hill/classes"
)

// CreatePlayerIDBuffer handles the buffer for the player
func CreatePlayerIDBuffer(player *classes.Player, idString string) PacketBuilder {
	var buffer []byte
	figurePacket := New(&buffer, Enums["Figure"])

	figurePacket.Write("uint32", player.NetID)
	figurePacket.Write("string", idString)

	for _, v := range idString {
		switch string(v) {

		// Position
		case "A":
			figurePacket.Write("float", player.Position.X)
		case "B":
			figurePacket.Write("float", player.Position.Y)
		case "C":
			figurePacket.Write("float", player.Position.Z)

		// Rotation
		case "D":
			figurePacket.Write("float", player.Rotation.X)
		case "E":
			figurePacket.Write("float", player.Rotation.Y)
		case "F":
			figurePacket.Write("float", player.Rotation.X)

		// Scale
		case "G":
			figurePacket.Write("float", player.Scale.X)
		case "H":
			figurePacket.Write("float", player.Scale.Y)
		case "I":
			figurePacket.Write("float", player.Scale.Z)

		//case "J"

		// Colors
		case "K":
			figurePacket.Write("uint32", player.Colors.Head)
		case "L":
			figurePacket.Write("uint32", player.Colors.Torso)
		case "M":
			figurePacket.Write("uint32", player.Colors.LeftArm)
		case "N":
			figurePacket.Write("uint32", player.Colors.RightArm)
		case "O":
			figurePacket.Write("uint32", player.Colors.LeftLeg)
		case "P":
			figurePacket.Write("uint32", player.Colors.RightLeg)

		// Assets
		case "Q":
			figurePacket.Write("uint32", player.Assets.Face)
		case "U":
			figurePacket.Write("uint32", player.Assets.Hat1)
		case "V":
			figurePacket.Write("uint32", player.Assets.Hat2)
		case "W":
			figurePacket.Write("uint32", player.Assets.Hat3)

		// Misc
		case "X":
			figurePacket.Write("int32", player.Score)
		case "Y":
			figurePacket.Write("uint32", player.Team)
		case "1":
			figurePacket.Write("uint32", player.Speed)
		case "2":
			figurePacket.Write("uint32", player.JumpPower)

		// Camera
		case "3":
			figurePacket.Write("uint32", player.CameraFOV)
		case "4":
			figurePacket.Write("int32", player.CameraDistance)

		case "5":
			figurePacket.Write("float", player.CameraPosition.X)
		case "6":
			figurePacket.Write("float", player.CameraPosition.Y)
		case "7":
			figurePacket.Write("float", player.CameraPosition.Z)

		case "8":
			figurePacket.Write("float", player.CameraRotation.X)
		case "9":
			figurePacket.Write("float", player.CameraRotation.Y)
		case "a":
			figurePacket.Write("float", player.CameraRotation.Z)

		case "b":
			figurePacket.Write("string", player.CameraType)
		// case "c":
		// 	figurePacket.Write("string", player.CameraType)

		case "e":
			figurePacket.Write("float", player.Health)
		case "g":
			figurePacket.Write("string", player.Speech)
		case "f":
			figurePacket.Write("uint32", uint32(0))
			figurePacket.Write("uint32", uint32(0))
		case "h":
			break
		}
	}

	return figurePacket
}
