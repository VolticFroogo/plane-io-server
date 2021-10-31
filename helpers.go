package main

import (
	"encoding/binary"
	"math"
)

func Float32FromBytes(bytes []byte) float32 {
	return math.Float32frombits(binary.BigEndian.Uint32(bytes))
}

func BytesFromFloat32(num float32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, math.Float32bits(num))
	return bytes
}
