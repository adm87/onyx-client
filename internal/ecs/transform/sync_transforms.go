package transform

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"github.com/yohamta/donburi/filter"
)

type TransformSyncSystem struct {
	query  *donburi.Query
	synced []*donburi.Entry

	OnSynced *events.EventType[[]*donburi.Entry]
}

func NewTransformSyncSystem() *TransformSyncSystem {
	return &TransformSyncSystem{
		query: donburi.NewQuery(
			filter.Contains(Transform),
		),
		synced:   make([]*donburi.Entry, 0, 128),
		OnSynced: events.NewEventType[[]*donburi.Entry](),
	}
}

func (s *TransformSyncSystem) Update(world donburi.World) {
	s.synced = s.synced[:0] // Clear the synced slice
	s.query.Each(world, func(entry *donburi.Entry) {
		t := GetTransform(entry)
		if !t.isDirty {
			return
		}
		t.isDirty = false

		m := GetGeoM(entry)
		m.Reset()
		m.Scale(t.scale.X, t.scale.Y)
		m.Rotate(t.rotation)
		m.Translate(t.position.X, t.position.Y)
		matrix.Set(entry, m)

		s.synced = append(s.synced, entry)
	})
	// s.OnSynced.Publish(world, s.synced)
}
