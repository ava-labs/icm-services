package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

func RunRelayer(
	ctx context.Context,
	relayerConfigPath string,
) (context.CancelFunc, chan struct{}) {
	relayerCtx, relayerCancel := context.WithCancel(ctx)
	fmt.Print(relayerConfigPath)
	relayerCmd := exec.CommandContext(relayerCtx, "./build/icm-relayer", "--config-file", relayerConfigPath)
	healthCheckURL := fmt.Sprintf("http://localhost:%d/health", 9090)

	readyChan := runExecutable(
		relayerCmd,
		relayerCtx,
		"icm-relayer",
		healthCheckURL,
	)
	return func() {
		relayerCancel()
		<-relayerCtx.Done()
	}, readyChan
}

func runExecutable(
	cmd *exec.Cmd,
	ctx context.Context,
	appName string,
	healthCheckUrl string,
) chan struct{} {
	cmdOutput := make(chan string)
	// file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	// Set up a pipe to capture the command's output
	cmdStdOutReader, _ := cmd.StdoutPipe()
	cmdStdErrReader, _ := cmd.StderrPipe()

	// Start the command
	log.Info("Starting executable", "appName", appName)
	_ = cmd.Start()

	readyChan := make(chan struct{})

	// Start goroutines to read and output the command's stdout and stderr
	go func() {
		scanner := bufio.NewScanner(cmdStdOutReader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		cmdOutput <- "Command execution finished"
	}()
	go func() {
		scanner := bufio.NewScanner(cmdStdErrReader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		cmdOutput <- "Command execution finished"
	}()
	go func() {
		err := cmd.Wait()
		// Context cancellation is the only expected way for the process to exit, otherwise log an error
		// Don't panic to allow for easier cleanup
		if !errors.Is(ctx.Err(), context.Canceled) {
			log.Error("Executable exited abnormally", "appName", appName, "err", err)
		}
	}()
	go func() { // wait for health check to report healthy
		for {
			resp, err := http.Get(healthCheckUrl)
			if err == nil && resp.StatusCode == 200 {
				log.Info("Health check passed", "appName", appName)
				close(readyChan)
				break
			}
			log.Info("Health check failed", "appName", appName, "err", err)
			time.Sleep(time.Second * 1)
		}
	}()
	return readyChan
}

func BuildAllExecutables(ctx context.Context) {
	cmd := exec.Command("./scripts/build.sh")
	out, _ := cmd.CombinedOutput()
	log.Info(string(out))
}
