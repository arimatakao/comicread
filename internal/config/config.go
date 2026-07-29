// Package config loads comicread's user configuration.
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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
	Favorites []string  `toml:"favorites"`
	Prerender Prerender `toml:"prerender"`
	Web       Web       `toml:"web"`
}

// Prerender configures how many neighbouring pages are rendered in advance.
type Prerender struct {
	Next     int `toml:"next"`
	Previous int `toml:"previous"`
}

// Web configures the --web local browser reader.
type Web struct {
	// Port the server listens on, on 127.0.0.1. 0 lets the OS assign any
	// free port instead of a fixed one.
	Port int `toml:"port"`
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
		Web: Web{
			Port: 55566,
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
	for i := range settings.Favorites {
		settings.Favorites[i] = strings.TrimSpace(settings.Favorites[i])
	}
	if settings.Graphics == "" {
		settings.Graphics = Default().Graphics
	}
	if settings.Prerender.Next < 0 || settings.Prerender.Previous < 0 {
		return Config{}, fmt.Errorf("parse config %q: prerender counts must not be negative", path)
	}
	if settings.Web.Port < 0 || settings.Web.Port > 65535 {
		return Config{}, fmt.Errorf("parse config %q: web.port must be between 0 and 65535", path)
	}
	return settings, nil
}

// Reset replaces the user config file with the default configuration.
func Reset() error {
	path, err := Path()
	if err != nil {
		return err
	}
	return ResetFile(path)
}

// ResetFile replaces path with the default configuration.
func ResetFile(path string) error {
	return SaveFile(path, Default())
}

// Save writes settings to the default config file.
func Save(settings Config) error {
	path, err := Path()
	if err != nil {
		return err
	}
	return SaveFile(path, settings)
}

// SaveFile writes settings to path.
func SaveFile(path string, settings Config) error {
	contents, err := toml.Marshal(settings)
	if err != nil {
		return fmt.Errorf("encode config: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}
	temporary, err := os.CreateTemp(filepath.Dir(path), ".config-*.tmp")
	if err != nil {
		return fmt.Errorf("create temporary config: %w", err)
	}
	temporaryName := temporary.Name()
	defer os.Remove(temporaryName)
	if _, err := temporary.Write(contents); err != nil {
		temporary.Close()
		return fmt.Errorf("write temporary config: %w", err)
	}
	if err := temporary.Chmod(0o600); err != nil {
		temporary.Close()
		return fmt.Errorf("set config permissions: %w", err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("close temporary config: %w", err)
	}
	if err := os.Rename(temporaryName, path); err != nil {
		return fmt.Errorf("replace config: %w", err)
	}
	return nil
}

// SetOption updates one supported option from a key=value assignment.
func SetOption(settings *Config, assignment string) error {
	key, value, ok := strings.Cut(assignment, "=")
	if !ok {
		return errors.New("config option must use key=value")
	}
	key = strings.ToLower(strings.TrimSpace(key))
	value = strings.TrimSpace(value)
	switch key {
	case "graphics":
		if !isGraphics(value) {
			return fmt.Errorf("unsupported graphics %q", value)
		}
		settings.Graphics = value
	case "view":
		if !isView(value) {
			return fmt.Errorf("unsupported view %q", value)
		}
		settings.View = value
	case "language":
		if value == "" {
			return errors.New("language must not be empty")
		}
		settings.Language = value
	case "directory":
		settings.Directory = value
	case "prerender.next", "prerender.previous":
		count, err := strconv.Atoi(value)
		if err != nil || count < 0 {
			return fmt.Errorf("%s must be a non-negative integer", key)
		}
		if key == "prerender.next" {
			settings.Prerender.Next = count
		} else {
			settings.Prerender.Previous = count
		}
	case "web.port":
		port, err := strconv.Atoi(value)
		if err != nil || port < 0 || port > 65535 {
			return fmt.Errorf("%s must be an integer between 0 and 65535", key)
		}
		settings.Web.Port = port
	default:
		return fmt.Errorf("unsupported config option %q", key)
	}
	return nil
}

func isGraphics(value string) bool {
	for _, allowed := range []string{"auto", "ascii", "ansi", "dots", "kitty", "sixel", "iterm", "iterm2"} {
		if value == allowed {
			return true
		}
	}
	return false
}

func isView(value string) bool {
	for _, allowed := range []string{"single-page", "book-view", "right-view", "circle-view", "right-circle-view"} {
		if value == allowed {
			return true
		}
	}
	return false
}
