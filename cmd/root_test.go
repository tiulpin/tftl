package cmd_test

import (
	"github.com/tiulpin/tftl/cmd"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseTerraformFile(t *testing.T) {
	content := `resource "aws_instance" "example" {}`
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "test.tf")

	err := os.WriteFile(filePath, []byte(content), 0644)
	require.NoError(t, err)

	targets, parseErr := cmd.ParseTerraformFile(filePath)
	require.NoError(t, parseErr)
	require.Equal(t, 1, len(targets))
	require.Equal(t, "aws_instance.example", targets[0])
}
