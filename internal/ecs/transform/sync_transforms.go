package transform

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/filter"
)

// OnSynced is an event that is published after the SyncTransformSystem has updated the transformation matrices of all entities.
// The event carries a slice of entries that were synced, allowing other systems to react to the changes.
//
// Note: The slice of entries is only valid during the event callback and should not be stored or used after the callback returns, as it may be reused in subsequent updates.
var OnSynced = events.NewEventType[[]*donburi.Entry]()

type SyncTransformSystem struct {
	query  *donburi.Query
	synced []*donburi.Entry
}

func NewSyncTransformSystem() *SyncTransformSystem {
	return &SyncTransformSystem{
		query: donburi.NewQuery(
			filter.Contains(Archetype()...),
		),
		synced: make([]*donburi.Entry, 0, 128),
	}
}

func (s *SyncTransformSystem) Update(world donburi.World) {
	s.synced = s.synced[:0] // Clear the synced slice
	s.query.Each(world, func(entry *donburi.Entry) {
		m := GetMatrix(entry)

		if !m.isDirty {
			return
		}

		t := GetTransform(entry)

		m.Matrix.Reset()
		m.Matrix.Scale(t.Scale.X, t.Scale.Y)
		m.Matrix.Rotate(t.Rotation)
		m.Matrix.Translate(t.Position.X, t.Position.Y)

		m.isDirty = false
		s.synced = append(s.synced, entry)
	})
	OnSynced.Publish(world, s.synced)
}
