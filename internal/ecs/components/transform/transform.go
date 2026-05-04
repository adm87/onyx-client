package transform

import (
	"github.com/yohamta/donburi"
)

func Reset(entry *donburi.Entry) {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	donburi.SetValue(entry, TransformComponent, defaultTransformModel)
}

func Position(entry *donburi.Entry) (x, y float64) {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	model := TransformComponent.Get(entry)
	return model.x, model.y
}

func Scale(entry *donburi.Entry) (sx, sy float64) {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	model := TransformComponent.Get(entry)
	return model.sx, model.sy
}

func Rotation(entry *donburi.Entry) float64 {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	model := TransformComponent.Get(entry)
	return model.rads
}

func SetPosition(entry *donburi.Entry, x, y float64) {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	model := TransformComponent.Get(entry)
	model.x = x
	model.y = y
	model.dirty = true
}

func SetScale(entry *donburi.Entry, sx, sy float64) {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	model := TransformComponent.Get(entry)
	model.sx = sx
	model.sy = sy
	model.dirty = true
}

func SetRotation(entry *donburi.Entry, rads float64) {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	model := TransformComponent.Get(entry)
	model.rads = rads
	model.dirty = true
}
