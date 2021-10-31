package main

const (
	BufferBuildThreshold   = 3
	BufferOptimalSize      = 5
	BufferDepleteThreshold = 7
	BufferBuildFrequency   = 5
)

type State struct {
	Position     Vector2
	Velocity     Vector2
	Rotation     float32
	MovementKeys uint8
}

type Buffer struct {
	states   []State
	building bool
}

func NewBuffer(state State) Buffer {
	return Buffer{
		states:   []State{state},
		building: false,
	}
}

func (buffer *Buffer) Push(state State) {
	buffer.states = append(buffer.states, state)

	if len(buffer.states) > BufferDepleteThreshold {
		buffer.states = buffer.states[1:]
	}
}

func (buffer *Buffer) Get() State {
	return buffer.states[0]
}

func (buffer *Buffer) Step() {
	if len(buffer.states) == 1 {
		return
	}

	if len(buffer.states) < BufferBuildThreshold {
		buffer.building = true
	}

	if buffer.building {
		if len(buffer.states) >= BufferOptimalSize {
			buffer.building = false
		} else if (server.Tick % BufferBuildFrequency) == 0 {
			return
		}
	}

	buffer.states = buffer.states[1:]
}
