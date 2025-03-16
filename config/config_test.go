package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	t.Run("Succes load file config", func(t *testing.T) {
		config := GetConfig()
		assert.NotNil(t, config)
	})
}
