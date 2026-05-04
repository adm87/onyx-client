package transform

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

func LocalToWorld(world donburi.World, entry *donburi.Entry, x, y float64) (worldX, worldY float64) {
	matrix := WorldMatrix(world, entry)
	return matrix.Apply(x, y)
}

func WorldToLocal(world donburi.World, entry *donburi.Entry, x, y float64) (localX, localY float64) {
	matrix := WorldMatrix(world, entry)
	matrix.Invert()
	return matrix.Apply(x, y)
}

func LocalMatrix(entry *donburi.Entry) ebiten.GeoM {
	if !entry.HasComponent(TransformComponent) {
		panic("entry does not have a transform component")
	}
	transform := TransformComponent.Get(entry)

	if !entry.HasComponent(matrixComponent) {
		entry.AddComponent(matrixComponent)
	}
	matrix := matrixComponent.Get(entry)

	if transform.dirty {
		matrix.local.Reset()
		matrix.local.Scale(transform.sx, transform.sy)
		matrix.local.Rotate(transform.rads)
		matrix.local.Translate(transform.x, transform.y)
		transform.dirty = false
	}
	return matrix.local
}

func WorldMatrix(world donburi.World, entry *donburi.Entry) ebiten.GeoM {
	local := LocalMatrix(entry)

	if !entry.HasComponent(hierarchyComponent) {
		return local
	}

	parent := GetParent(world, entry)
	if parent == nil {
		return local
	}
	parentWorld := WorldMatrix(world, parent)

	var matrix ebiten.GeoM
	matrix.Concat(local)
	matrix.Concat(parentWorld)
	return matrix
}
