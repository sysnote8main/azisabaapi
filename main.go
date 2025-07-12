package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/sysnote8main/azisabaapi/pkg/aziapi"
)

var (
	API_ENDPOINT string = "https://api-ktor.azisaba.net"
)

func main() {
	apiToken, ok := os.LookupEnv("AZI_API_TOKEN")
	if !ok {
		slog.Error("Failed to get api token from environment. Please set api token as AZI_API_TOKEN")
		os.Exit(1)
	}

	apiClient := aziapi.NewAziApiClient(
		apiToken,
		API_ENDPOINT,
	)

	res, err := apiClient.GetCounts()
	if err != nil {
		slog.Error("Failed to get counts", slog.Any("error", err))
		os.Exit(1)
	}
	fmt.Printf("Data: %v\n", *res)
	fmt.Printf("Total players: %d", res.TotalPlayers)
}

func smain() {
	apiToken, ok := os.LookupEnv("AZI_API_TOKEN")
	if !ok {
		slog.Error("Failed to get api token from environment. Please set AZI_API_TOKEN")
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, "https://api-ktor.azisaba.net/counts", nil)
	if err != nil {
		slog.Error("Failed to create new request", slog.Any("error", err))
		os.Exit(1)
	}

	req.Header.Add("Authorization", "Bearer "+apiToken)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send get request", slog.Any("error", err))
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		slog.Error("Failed to request", slog.Int("status_code", res.StatusCode))
		os.Exit(1)
	}

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Failed to read body", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("Result", slog.String("data", string(b)))
}
