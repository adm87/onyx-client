package transform

import "github.com/yohamta/donburi"

func GetParent(world donburi.World, entry *donburi.Entry) *donburi.Entry {
	if !entry.HasComponent(hierarchyComponent) {
		return nil
	}
	hierarchy := hierarchyComponent.Get(entry)
	if hierarchy.parent == donburi.Null {
		return nil
	}
	return world.Entry(hierarchy.parent)
}

func SetParent(world donburi.World, parent, child *donburi.Entry) {
	if child == nil {
		panic("child cannot be nil")
	}
	if child == parent {
		panic("child cannot be its own parent")
	}
	if isDescendant(world, child, parent) {
		panic("cannot set parent to a descendant")
	}

	if !child.HasComponent(hierarchyComponent) {
		if parent == nil {
			return // No parent, no hierarchy needed
		}
		child.AddComponent(hierarchyComponent)
	}

	childEntity := child.Entity()
	childHierarchy := hierarchyComponent.Get(child)
	detachFromParent(world, childHierarchy, childEntity)

	if parent != nil {
		if !parent.HasComponent(hierarchyComponent) {
			parent.AddComponent(hierarchyComponent)
		}
		parentEntity := parent.Entity()
		parentHierarchy := hierarchyComponent.Get(parent)
		attachChild(world, parentHierarchy, childHierarchy, parentEntity, childEntity)
	}
}

func detachFromParent(world donburi.World, childHierarchy *hierarchyModel, childEntity donburi.Entity) {
	if childHierarchy.parent == donburi.Null {
		return
	}

	parent := world.Entry(childHierarchy.parent)
	parentHierarchy := hierarchyComponent.Get(parent)

	if childHierarchy.prev != donburi.Null {
		prev := world.Entry(childHierarchy.prev)
		prevHierarchy := hierarchyComponent.Get(prev)
		prevHierarchy.next = childHierarchy.next
	} else {
		parentHierarchy.first = childHierarchy.next
	}

	if childHierarchy.next != donburi.Null {
		next := world.Entry(childHierarchy.next)
		nextHierarchy := hierarchyComponent.Get(next)
		nextHierarchy.prev = childHierarchy.prev
	}

	childHierarchy.parent = donburi.Null
	childHierarchy.prev = donburi.Null
	childHierarchy.next = donburi.Null
}

func attachChild(world donburi.World, parentHierarchy, childHierarchy *hierarchyModel, parentEntity, childEntity donburi.Entity) {
	childHierarchy.parent = parentEntity
	childHierarchy.prev = donburi.Null
	childHierarchy.next = parentHierarchy.first

	if parentHierarchy.first != donburi.Null {
		first := world.Entry(parentHierarchy.first)
		firstHierarchy := hierarchyComponent.Get(first)
		firstHierarchy.prev = childEntity
	}
	parentHierarchy.first = childEntity
}

func isDescendant(world donburi.World, ancestor, descendant *donburi.Entry) bool {
	if ancestor == nil || descendant == nil {
		return false
	}
	if !descendant.HasComponent(hierarchyComponent) {
		return false
	}
	hierarchy := hierarchyComponent.Get(descendant)
	for hierarchy.parent != donburi.Null {
		if hierarchy.parent == ancestor.Entity() {
			return true
		}
		parent := world.Entry(hierarchy.parent)
		hierarchy = hierarchyComponent.Get(parent)
	}
	return false
}
