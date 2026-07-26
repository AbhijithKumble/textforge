package progress

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Store struct {
	Completed map[string]bool `json:"completed"`
}

func Load(path string) (*Store, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Store{Completed: make(map[string]bool)}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read progress: %w", err)
	}

	var store Store
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("parse progress: %w", err)
	}
	if store.Completed == nil {
		store.Completed = make(map[string]bool)
	}
	return &store, nil
}

func (s *Store) Complete(id string) {
	s.Completed[id] = true
}

func (s *Store) IsComplete(id string) bool {
	return s.Completed[id]
}

func (s *Store) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create progress directory: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("encode progress: %w", err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
		return fmt.Errorf("write progress: %w", err)
	}
	return nil
}
