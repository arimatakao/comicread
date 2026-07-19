// Package config loads comicread's user configuration.
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

const fileName = "config.toml"

// Config contains the configurable reader defaults.
type Config struct {
	Graphics  string    `toml:"graphics"`
	View      string    `toml:"view"`
	Language  string    `toml:"language"`
	Directory string    `toml:"directory"`
	Prerender Prerender `toml:"prerender"`
}

// Prerender configures how many neighbouring pages are rendered in advance.
type Prerender struct {
	Next     int `toml:"next"`
	Previous int `toml:"previous"`
}

// Default returns the configuration used when no config file exists.
func Default() Config {
	return Config{
		Graphics: "auto",
		View:     "single-page",
		Language: "en",
		Prerender: Prerender{
			Next:     1,
			Previous: 1,
		},
	}
}

// Path returns the default config.toml location for the current user.
func Path() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("find user config directory: %w", err)
	}
	return filepath.Join(dir, "comicread", fileName), nil
}

// Load reads the default config file. A missing file uses Default.
func Load() (Config, error) {
	path, err := Path()
	if err != nil {
		return Config{}, err
	}
	return LoadFile(path)
}

// LoadFile reads a TOML configuration file. A missing file uses Default.
func LoadFile(path string) (Config, error) {
	settings := Default()
	contents, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return settings, nil
	}
	if err != nil {
		return Config{}, fmt.Errorf("read config %q: %w", path, err)
	}
	if err := toml.Unmarshal(contents, &settings); err != nil {
		return Config{}, fmt.Errorf("parse config %q: %w", path, err)
	}
	settings.Graphics = strings.TrimSpace(settings.Graphics)
	settings.View = strings.TrimSpace(settings.View)
	settings.Language = strings.TrimSpace(settings.Language)
	settings.Directory = strings.TrimSpace(settings.Directory)
	if settings.Graphics == "" {
		settings.Graphics = Default().Graphics
	}
	if settings.Prerender.Next < 0 || settings.Prerender.Previous < 0 {
		return Config{}, fmt.Errorf("parse config %q: prerender counts must not be negative", path)
	}
	return settings, nil
}
