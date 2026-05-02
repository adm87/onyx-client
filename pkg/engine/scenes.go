package engine

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	SceneExitCodeNone SceneExitCode = 0
)

const (
	SceneIDNone SceneId = ""
)

type SceneId string

func (id SceneId) IsNone() bool {
	return id == SceneIDNone
}

type SceneExitCode uint8

func (c SceneExitCode) IsNone() bool {
	return c == SceneExitCodeNone
}

type SceneRouter map[SceneId]map[SceneExitCode]SceneId

func (r SceneRouter) AddRoute(from SceneId, exitCode SceneExitCode, to SceneId) {
	if _, ok := r[from]; !ok {
		r[from] = make(map[SceneExitCode]SceneId)
	}
	r[from][exitCode] = to
}

func (r SceneRouter) GetNext(from SceneId, exitCode SceneExitCode) (SceneId, bool) {
	if routes, ok := r[from]; ok {
		to, exists := routes[exitCode]
		return to, exists
	}
	return SceneIDNone, false
}

type Scene interface {
	OnEnter() error
	OnExit() error
	Update() (SceneExitCode, error)
	Draw(screen *ebiten.Image) error
}

type SceneEntry struct {
	Id          SceneId
	Ctor        func() Scene
	Transitions map[SceneExitCode]SceneId
}

type Scenes struct {
	logger *Logger

	current   Scene
	currentID SceneId
	pendingID SceneId

	scenes map[SceneId]*SceneEntry
	router SceneRouter
}

func NewScenes(logger *Logger) *Scenes {
	return &Scenes{
		logger: logger,
		scenes: make(map[SceneId]*SceneEntry),
		router: make(SceneRouter),
	}
}

func (s *Scenes) Add(scenes ...*SceneEntry) {
	for _, entry := range scenes {
		if entry.Id.IsNone() {
			s.logger.Warn("scene with ID 0 is reserved for 'none', skipping")
			continue
		}

		if _, ok := s.scenes[entry.Id]; ok {
			s.logger.Warn("scene already exists, skipping", "id", entry.Id)
			continue
		}

		s.scenes[entry.Id] = entry
		if len(entry.Transitions) > 0 {
			s.buildRouting(entry.Id, entry.Transitions)
		}
	}
}

func (s *Scenes) Start(id SceneId) error {
	if err := s.Exit(); err != nil {
		return err
	}

	entry, ok := s.scenes[id]
	if !ok {
		return fmt.Errorf("scene not found %s", id)
	}

	s.current = entry.Ctor()
	s.currentID = id

	s.logger.Debug("starting scene", "id", id)
	return s.current.OnEnter()
}

func (s *Scenes) Exit() error {
	if s.current == nil {
		return nil
	}

	if err := s.current.OnExit(); err != nil {
		return fmt.Errorf("failed to exit current scene id=%s error=%w", s.currentID, err)
	}
	s.logger.Debug("exited scene", "id", s.currentID)

	s.current = nil
	s.currentID = SceneIDNone
	return nil
}

func (s *Scenes) Update() error {
	if !s.pendingID.IsNone() {
		if err := s.Start(s.pendingID); err != nil {
			return fmt.Errorf("failed to start pending scene id=%s error=%w", s.pendingID, err)
		}
		s.pendingID = SceneIDNone
	}

	if s.current == nil {
		return nil
	}

	exitCode, err := s.current.Update()
	if err != nil {
		return err
	}
	if exitCode.IsNone() {
		return nil
	}

	if nextID, exists := s.router.GetNext(s.currentID, exitCode); exists {
		s.pendingID = nextID
	}
	return nil
}

func (s *Scenes) Draw(screen *ebiten.Image) error {
	if s.current != nil {
		return s.current.Draw(screen)
	}
	return nil
}

func (s *Scenes) buildRouting(id SceneId, transitions map[SceneExitCode]SceneId) {
	for exitCode, to := range transitions {
		s.logger.Debug("adding scene route", "from_scene_id", id, "exit_code", exitCode, "to_scene_id", to)
		s.router.AddRoute(id, exitCode, to)
	}
}
