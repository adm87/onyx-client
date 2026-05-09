package transform

import (
	"github.com/adm87/onyx/internal/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type (
	TransformData struct {
		position geom.Vec2
		rotation float64
		scale    geom.Vec2
		isDirty  bool
	}
)

var (
	Transform = donburi.NewComponentType[TransformData](TransformData{
		scale:   geom.OneVec2(),
		isDirty: true,
	})
	matrix = donburi.NewComponentType[ebiten.GeoM]()
)

func GetTransform(entry *donburi.Entry) *TransformData {
	if !entry.HasComponent(Transform) {
		panic("entry does not have Transform component")
	}
	return donburi.Get[TransformData](entry, Transform)
}

func GetGeoM(entry *donburi.Entry) *ebiten.GeoM {
	if !entry.HasComponent(matrix) {
		entry.AddComponent(matrix)
	}
	return matrix.Get(entry)
}

func GetPosition(entry *donburi.Entry) geom.Vec2 {
	return GetTransform(entry).position
}

func GetRotation(entry *donburi.Entry) float64 {
	return GetTransform(entry).rotation
}

func GetScale(entry *donburi.Entry) geom.Vec2 {
	return GetTransform(entry).scale
}

func SetPosition(entry *donburi.Entry, position geom.Vec2) {
	t := GetTransform(entry)
	t.position = position
	t.isDirty = true
}

func SetPositionXY(entry *donburi.Entry, x, y float64) {
	t := GetTransform(entry)
	t.position.X = x
	t.position.Y = y
	t.isDirty = true
}

func SetRotation(entry *donburi.Entry, rotation float64) {
	t := GetTransform(entry)
	t.rotation = rotation
	t.isDirty = true
}

func SetScale(entry *donburi.Entry, scale geom.Vec2) {
	t := GetTransform(entry)
	t.scale = scale
	t.isDirty = true
}

func SetScaleXY(entry *donburi.Entry, x, y float64) {
	t := GetTransform(entry)
	t.scale.X = x
	t.scale.Y = y
	t.isDirty = true
}
