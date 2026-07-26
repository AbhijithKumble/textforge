package course

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/adrg/frontmatter"
)

type Lesson struct {
	ID          string     `yaml:"id"`
	Title       string     `yaml:"title"`
	Index       int64      `yaml:"index"`
	Description string     `yaml:"description"`
	Difficulty  string     `yaml:"difficulty"`
	Exercises   []Exercise `yaml:"exercises"`
	Path        string     `yaml:"-"`
	Content     string
}

type Exercise struct {
	ID     string `yaml:"id"`
	Prompt string `yaml:"prompt"`
	Answer string `yaml:"answer"`
}

func LoadLesson(filePath string) (*Lesson, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("read lesson %q: %w", filePath, err)
	}

	defer file.Close()

	var lesson Lesson

	contentByteArray, err := frontmatter.Parse(file, &lesson)

	if err != nil {
		return nil, fmt.Errorf("parse lesson %q: %w", filePath, err)
	}

	lesson.Content = string(contentByteArray)

	return &lesson, nil
}

func DiscoverLessons(root string) ([]*Lesson, error) {
	var paths []string
	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			paths = append(paths, path)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("discover lessons in %q: %w", root, err)
	}

	sort.Strings(paths)
	lessons := make([]*Lesson, 0, len(paths))
	for _, path := range paths {
		lesson, err := LoadLesson(path)
		if err != nil {
			return nil, err
		}
		lesson.Path = path
		lessons = append(lessons, lesson)
	}

	sort.SliceStable(lessons, func(i, j int) bool {
		if lessons[i].Index == lessons[j].Index {
			return lessons[i].Path < lessons[j].Path
		}
		return lessons[i].Index < lessons[j].Index
	})
	return lessons, nil
}
