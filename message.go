package main

const (
	MessageTypePing = iota
	MessageTypeUpdate
	MessageTypeInstantiate
)

func HandleMessage(player *Player, bytes []byte) {
	offset := 0

	for {
		messageType := uint8(bytes[offset])

		switch messageType {
		case MessageTypeUpdate:
			offset = HandleMessageUpdate(player, bytes, offset)
			break
		}

		if offset >= len(bytes) {
			break
		}
	}
}

func HandleMessageUpdate(player *Player, bytes []byte, offset int) (nextOffset int) {
	// Handle initial type offset.
	offset += 1

	state := State{}

	state.Position.X = Float32FromBytes(bytes[offset : offset+4])
	offset += 4

	state.Position.Y = Float32FromBytes(bytes[offset : offset+4])
	offset += 4

	state.Rotation = Float32FromBytes(bytes[offset : offset+4])
	offset += 4

	state.MovementKeys = uint8(bytes[offset])
	offset += 1

	player.Buffer.Push(state)

	return offset
}
