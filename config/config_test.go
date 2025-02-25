package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfg(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		config := GetConfig()
		require.NotNil(t, config, "Config should not be nil!")

	})
}
