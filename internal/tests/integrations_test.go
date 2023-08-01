package tests

import (
	"context"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const binaryPath = "../../bin/currency-rates"

func TestBasicIntegration(t *testing.T) {

	t.Run("ok with stdin input and stdout result", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		cmd := exec.CommandContext(ctx, binaryPath, "-code", "USD", "-date", "2022-08-10")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		err := cmd.Run()
		assert.NoError(t, err)
		assert.Equal(t, "USD (Доллар США): 60.381400\n", stdout.String())
	})

	t.Run("error with invalid code", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		cmd := exec.CommandContext(ctx, binaryPath, "-code", "", "-date", "2022-08-10")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr
		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
	})

	t.Run("error with invalid date", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		cmd := exec.CommandContext(ctx, binaryPath, "-code", "USD", "-date", "202208-10")
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr
		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
	})

	t.Run("error with no parameters", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		cmd := exec.CommandContext(ctx, binaryPath)
		stdout := &strings.Builder{}
		cmd.Stdout = stdout
		stderr := &strings.Builder{}
		cmd.Stderr = stderr
		err := cmd.Run()

		assert.Error(t, err)
		assert.NotZero(t, stderr.Len())
	})
}
