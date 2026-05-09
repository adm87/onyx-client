package transform

import (
	"github.com/adm87/onyx/internal/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type (
	TransformModel struct {
		Position geom.Vec2
		Rotation float64
		Scale    geom.Vec2
	}
	MatrixModel struct {
		Matrix  ebiten.GeoM
		isDirty bool
	}
)

var (
	Transform = donburi.NewComponentType[TransformModel](TransformModel{
		Scale: geom.OneVec2(),
	})
	Matrix = donburi.NewComponentType[MatrixModel](MatrixModel{
		isDirty: true,
	})
)

// Archetype returns the list of components that make up a common Transform archetype.
func Archetype() []donburi.IComponentType {
	return []donburi.IComponentType{
		Transform,
		Matrix,
	}
}

func GetTransform(entry *donburi.Entry) *TransformModel {
	if !entry.HasComponent(Transform) {
		panic("entry does not have Transform component")
	}
	return donburi.Get[TransformModel](entry, Transform)
}

func GetMatrix(entry *donburi.Entry) *MatrixModel {
	if !entry.HasComponent(Matrix) {
		panic("entry does not have Matrix component")
	}
	return donburi.Get[MatrixModel](entry, Matrix)
}

func GetPosition(entry *donburi.Entry) geom.Vec2 {
	return GetTransform(entry).Position
}

func GetRotation(entry *donburi.Entry) float64 {
	return GetTransform(entry).Rotation
}

func GetScale(entry *donburi.Entry) geom.Vec2 {
	return GetTransform(entry).Scale
}

func SetPosition(entry *donburi.Entry, position geom.Vec2) {
	t := GetTransform(entry)
	t.Position = position
	GetMatrix(entry).isDirty = true
}

func SetRotation(entry *donburi.Entry, rotation float64) {
	t := GetTransform(entry)
	t.Rotation = rotation
	GetMatrix(entry).isDirty = true
}

func SetScale(entry *donburi.Entry, scale geom.Vec2) {
	t := GetTransform(entry)
	t.Scale = scale
	GetMatrix(entry).isDirty = true
}
