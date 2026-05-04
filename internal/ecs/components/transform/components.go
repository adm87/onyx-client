package transform

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type (
	transformModel struct {
		x, y   float64
		sx, sy float64
		rads   float64
		dirty  bool
	}
	matrixModel struct {
		local ebiten.GeoM
	}
	hierarchyModel struct {
		parent donburi.Entity
		first  donburi.Entity
		next   donburi.Entity
		prev   donburi.Entity
	}
)

var (
	defaultTransformModel = transformModel{
		x:     0,
		y:     0,
		sx:    1,
		sy:    1,
		rads:  0,
		dirty: true,
	}
	defaultHierarchyModel = hierarchyModel{
		parent: donburi.Null,
		first:  donburi.Null,
		next:   donburi.Null,
		prev:   donburi.Null,
	}
)

var (
	hierarchyComponent = donburi.NewComponentType[hierarchyModel](defaultHierarchyModel)
	matrixComponent    = donburi.NewComponentType[matrixModel](matrixModel{})
)

var TransformComponent = donburi.NewComponentType[transformModel](defaultTransformModel)
