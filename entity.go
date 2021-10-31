package main

const (
	EntityTypePlayer = iota
)

type EntityI interface {
	Step()
	Collide(other *Entity)
	Despawn()

	GetID() uint32
	GetPrev() EntityI
	SetPrev(prev EntityI)
	GetNext() EntityI
	SetNext(next EntityI)
	GetPosition() Vector2
	GetVelocity() Vector2
	GetRotation() float32
	GetRadius() float32

	CreateUpdatePacket() []byte
	CreateInstantiatePacket() []byte
}

type Entity struct {
	ID       uint32
	Prev     EntityI
	Next     EntityI
	Position Vector2
	Velocity Vector2
	Rotation float32
	Radius   float32
}

func (entity *Entity) Step() {

}

func (entity *Entity) Collide(other *Entity) {

}

func (entity *Entity) Despawn() {
	if entity.Prev != nil {
		entity.Prev.SetNext(entity.Next)
		entity.Prev = nil
	} else if entity == server.EntityMap[entity.GetPosition().Chunk()] {
		server.EntityMap[entity.GetPosition().Chunk()] = entity.Next
	}

	if entity.Next != nil {
		entity.Next.SetPrev(entity.Prev)
		entity.Next = nil
	}
}

func (entity *Entity) GetID() uint32 {
	return entity.ID
}

func (entity *Entity) GetPrev() EntityI {
	return entity.Next
}

func (entity *Entity) SetPrev(prev EntityI) {
	entity.Prev = prev
}

func (entity *Entity) GetNext() EntityI {
	return entity.Next
}

func (entity *Entity) SetNext(next EntityI) {
	entity.Next = next
}

func (entity *Entity) GetPosition() Vector2 {
	return entity.Position
}

func (entity *Entity) GetVelocity() Vector2 {
	return entity.Velocity
}

func (entity *Entity) GetRotation() float32 {
	return entity.Rotation
}

func (entity *Entity) GetRadius() float32 {
	return entity.Radius
}

func (entity *Entity) CreateUpdatePacket() []byte {
	return nil
}

func (entity *Entity) CreateInstantiatePacket() []byte {
	return nil
}
