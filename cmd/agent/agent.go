package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/developer-profile/devmetr/internal/agent"
	delaultMetrics "github.com/developer-profile/devmetr/internal/agent/delault_metrics"
	repositiry "github.com/developer-profile/devmetr/internal/agent/repository"
	trackingMetricStore "github.com/developer-profile/devmetr/internal/agent/tracking_metric_store"
	"github.com/developer-profile/devmetr/internal/agent/transport"
)

func main() {

	listMetricForTrack := trackingMetricStore.New()
	listMetricForTrack.Add(delaultMetrics.DefaultRuntimeMetric)
	listMetricForTrack.Add(delaultMetrics.DefaultCustomMetric)
	log.Println("Server started properly.")
	a := agent.New(
		2,
		10,
		transport.NewHTTPClient("127.0.0.1:8080", &http.Client{}),
		repositiry.NewRepoMem(),
		listMetricForTrack)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := &sync.WaitGroup{}

	a.Start(ctx, cancel, wg)

	SignalChanel := make(chan os.Signal, 1)
	signal.Notify(SignalChanel,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
