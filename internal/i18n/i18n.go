// Package i18n provides minimal message translation for comicread's
// user-facing strings (status messages, view labels, CLI errors).
package i18n

import (
	"fmt"
	"os"
	"strings"
)

// Lang identifies a supported message language.
type Lang string

const (
	English   Lang = "en"
	Ukrainian Lang = "uk"
)

var current = detect()

// detect picks a language from the COMICREAD_LANG environment variable,
// falling back to English.
func detect() Lang {
	if strings.HasPrefix(strings.ToLower(os.Getenv("COMICREAD_LANG")), "uk") {
		return Ukrainian
	}
	return English
}

// SetLang overrides the active language (used mainly for testing/tools).
func SetLang(lang Lang) {
	current = lang
}

// T returns the message for key in the active language, formatting it with
// args via fmt.Sprintf when args are given. It falls back to English, then
// to the key itself, if a translation is missing.
func T(key string, args ...any) string {
	msg, ok := messages[current][key]
	if !ok {
		msg, ok = messages[English][key]
	}
	if !ok {
		msg = key
	}
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

var messages = map[Lang]map[string]string{
	English:   enMessages,
	Ukrainian: ukMessages,
}
