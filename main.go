package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
		logger.Error("Failed to run command", zap.String("command", cmd), zap.Error(err))
		os.Stderr.WriteString(err.Error() + "\n")
	}
}

func sendStatusUpdate(status string) {
	logger := GetLogger()
	url := "https://train-status.computehubs.com/train/status"
	if os.Getenv("CHUB_ENV") == "local" {
		url = "https://da9b-83-136-182-39.ngrok-free.app/train/status"
	}

	payload := map[string]string{
		"status": "PREPROCESSING",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CHUB-API-KEY", os.Getenv("CHUB_HTTP_TK"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error sending status update", zap.Error(err))
		return
	}
	defer resp.Body.Close()

	logger.Info("Status update sent", zap.String("status", status))
}

func downloadFromS3(ctx context.Context, logger *zap.Logger) {
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/user-yaml-config/%s.yaml",
		os.Getenv("CHUB_CONTAINER_REGISTRY"),
		os.Getenv("CHUB_KEY"))
	logger.Debug(url)

	cmd := []string{
		fmt.Sprintf("curl %s -o custom.yaml", url),
		"mkdir -p examples/custom_model",
		"mv custom.yaml examples/custom_model/custom.yaml",
	}

	for _, c := range cmd {
		runShellCommandsRealtime(ctx, logger, c, true)
	}
}

func main() {
	logger := NewLogger()
	sendStatusUpdate("Preprocessing started")

	downloadFromS3(context.Background(), logger)

	cmd := []string{
		"pip install --no-deps -e .",
		"pwd",
		"accelerate launch -m axolotl.cli.train examples/custom_model/custom.yaml",
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, c := range cmd {
		runShellCommandsRealtime(ctx, logger, c, true)
	}
	logger.Info("successfully ran command")
}
