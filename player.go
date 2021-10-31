package main

import (
	"encoding/binary"
	"log"

	"github.com/gorilla/websocket"
)

type Player struct {
	Entity
	Buffer        Buffer
	Conn          *websocket.Conn
	KnownEntities map[uint32]bool
}

func (player *Player) Step() {
	player.Buffer.Step()

	// Send player an update packet.
	packet := []byte{}

	entity := server.EntityMap[0]
	for entity != nil {
		if player == entity {
			entity = entity.GetNext()
			continue
		}

		if _, ok := player.KnownEntities[entity.GetID()]; ok {
			packet = append(packet, entity.CreateUpdatePacket()...)
		} else {
			packet = append(packet, entity.CreateInstantiatePacket()...)
			player.KnownEntities[entity.GetID()] = true
		}

		entity = entity.GetNext()
	}

	if len(packet) == 0 {
		return
	}

	err := player.Conn.WriteMessage(websocket.BinaryMessage, packet)
	if err != nil {
		log.Println(err)
		player.Despawn()
		return
	}
}

func (player *Player) Collide(other *Entity) {

}

func (player *Player) CreateUpdatePacket() (bytes []byte) {
	bytes = append(bytes, uint8(MessageTypeUpdate))

	idBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(idBytes, player.ID)
	bytes = append(bytes, idBytes...)

	bytes = append(bytes, BytesFromFloat32(player.Buffer.Get().Position.X)...)
	bytes = append(bytes, BytesFromFloat32(player.Buffer.Get().Position.Y)...)
	bytes = append(bytes, BytesFromFloat32(player.Buffer.Get().Rotation)...)
	bytes = append(bytes, player.Buffer.Get().MovementKeys)

	return
}

func (player *Player) CreateInstantiatePacket() (bytes []byte) {
	bytes = append(bytes, uint8(MessageTypeInstantiate))
	bytes = append(bytes, uint8(EntityTypePlayer))

	idBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(idBytes, player.ID)
	bytes = append(bytes, idBytes...)

	bytes = append(bytes, BytesFromFloat32(player.Buffer.Get().Position.X)...)
	bytes = append(bytes, BytesFromFloat32(player.Buffer.Get().Position.Y)...)
	bytes = append(bytes, BytesFromFloat32(player.Buffer.Get().Rotation)...)
	bytes = append(bytes, player.Buffer.Get().MovementKeys)

	return
}
