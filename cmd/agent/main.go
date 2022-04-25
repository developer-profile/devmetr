package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/developer-profile/devmetr/internal/agent"
	"github.com/developer-profile/devmetr/internal/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	log.Println("Let the agent show begin!")
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-cancelChan
		cancel()
	}()

	address := pflag.StringP("address", "a", config.DefaultServer, "")
	pollInterval := pflag.DurationP("pool-inreval", "p", config.DefaultPollInterval, "")
	reportInterval := pflag.DurationP("report-interval", "r", config.DefaultReportInterval, "")
	key := pflag.StringP("key", "k", "", "")
	pflag.Parse()

	v := viper.New()
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	conf := config.NewAgentConfigWithDefaults(v, *address, *pollInterval, *reportInterval, *key)
	collector := agent.New(*conf)
	collector.Run(ctx)

	log.Println("the end of agent show")
}
