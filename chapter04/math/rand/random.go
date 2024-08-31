package rand

import (
	"math/rand/v2"

	"github.com/ishtaka/go-game-programming/chapter03/math"
)

func GetFloat() float32 {
	return rand.Float32()
}

func GetFloatRange(min, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

func GetInt() int {
	return rand.Int()
}

func GetIntRange(min, max int) int {
	return min + rand.IntN(max-min)
}

func GetVector2(min, max math.Vector2) math.Vector2 {
	v := math.Vector2{X: GetFloat(), Y: GetFloat()}
	return min.Add(v.Mul(max.Sub(min)))
}

func GetVector3(min, max math.Vector3) math.Vector3 {
	v := math.Vector3{X: GetFloat(), Y: GetFloat(), Z: GetFloat()}
	return min.Add(v.Mul(max.Sub(min)))
}
