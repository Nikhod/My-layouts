package undo_funcs

import (
	"os"
)

// SetEnv sets new environment variable then return undo_func that do rollback
func SetEnv(key, value string) (undo func()) {
	oldValue := os.Getenv(key)
	_ = os.Setenv(key, value)

	return func() {
		_ = os.Setenv(key, oldValue)
	}
}
