package main

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"go.uber.org/zap"
)

func runShellCommandsRealtime(ctx context.Context, logger *zap.Logger, cmd string, verbose bool) {
	command := strings.Fields(cmd)
	comm := exec.CommandContext(ctx, command[0], command[1:]...)

	if verbose {
		comm.Stdout = os.Stdout
	}
	comm.Stderr = os.Stderr

	if err := comm.Run(); err != nil {
		// Handle potential errors.
		logger.Error("Failed to run command", zap.String("command", cmd), zap.Error(err))
		os.Stderr.WriteString(err.Error() + "\n")
	}
}

func main() {
	logger := NewLogger()

	runShellCommandsRealtime(context.Background(), logger, "accelerate launch -m axolotl.cli.train examples/openllama-3b/lora.yml", true)
	logger.Info("successfully ran command")
}
