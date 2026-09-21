// Package fm reads and writes frontmatter and YAML files.
package fm

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var ErrNoFrontmatter = errors.New("no frontmatter")

// Split returns the YAML frontmatter of a Markdown file and its body.
func Split(b []byte) (front, body []byte, err error) {
	b = bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
	if !bytes.HasPrefix(b, []byte("---\n")) {
		return nil, b, ErrNoFrontmatter
	}
	rest := b[4:]
	end := bytes.Index(rest, []byte("\n---\n"))
	if end < 0 {
		if bytes.HasSuffix(rest, []byte("\n---")) {
			return rest[:len(rest)-4], nil, nil
		}
		return nil, b, ErrNoFrontmatter
	}
	return rest[:end+1], rest[end+5:], nil
}

// ReadFront decodes the frontmatter of a Markdown file into v, and also returns it as
// a generic map so that field presence can be checked.
func ReadFront(path string, v any) (map[string]any, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	front, _, err := Split(b)
	if err != nil {
		return nil, err
	}
	raw := map[string]any{}
	if err := yaml.Unmarshal(front, &raw); err != nil {
		return nil, err
	}
	if v != nil {
		if err := yaml.Unmarshal(front, v); err != nil {
			return raw, err
		}
	}
	return raw, nil
}

// ReadYAML decodes a YAML file into v. A missing file is reported as os.ErrNotExist.
func ReadYAML(path string, v any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(b, v)
}

// WriteYAML writes v as YAML, creating directories as needed.
func WriteYAML(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
