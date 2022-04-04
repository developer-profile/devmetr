package agent

import (
	"context"
	"log"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"

	"github.com/developer-profile/devmetr/internal/models"
)

type MetricStorer interface {
	AddMetricValue(mtype, name, value string)
	GetMetricValue(mtype, name string) (string, error)
	GetAll() ([]models.Metric, error)
}

type Transporter interface {
	SendMetric(m models.Metric) error
}

type TrackingMetricsStorer interface {
	Add(...interface{})
	GetCustomMetrics() []models.CustomMetric
	GetRuntimeMetric() []models.RuntimeMetric
}

type Agent struct {
	pollInterval         int
	reportInterval       int
	trackingMetricsStore TrackingMetricsStorer
	metricStore          MetricStorer
	transport            Transporter
}

func New(
	secPollInterval int,
	secReportInterval int,
	client Transporter,
	metricStore MetricStorer,
	trackingMetricsStore TrackingMetricsStorer) *Agent {
	return &Agent{
		pollInterval:         secPollInterval,
		reportInterval:       secReportInterval,
		trackingMetricsStore: trackingMetricsStore,
		metricStore:          metricStore,
		transport:            client,
	}
}

func (a *Agent) AddMetrics(metrics ...interface{}) {
	a.trackingMetricsStore.Add(metrics)
}

func (a *Agent) updateMetric() {

	rms := &runtime.MemStats{}
	runtime.ReadMemStats(rms)
	r := reflect.Indirect(reflect.ValueOf(rms))
	log.Println("updateMetric() start")
	for _, m := range a.trackingMetricsStore.GetRuntimeMetric() {
		v, _ := m.UpdateFunc(r, m)
		a.metricStore.AddMetricValue(m.TypeM, m.Name, v)
	}

	for _, m := range a.trackingMetricsStore.GetCustomMetrics() {
		v, _ := a.metricStore.GetMetricValue(m.TypeM, m.Name)
		value, _ := m.UpdateFunc(v)
		a.metricStore.AddMetricValue(m.TypeM, m.Name, value)
	}
	log.Println("updateMetric() finish")
}

func (a *Agent) sendMetric(ch chan<- models.Metric) {
	log.Println("sendMetric() start")
	metrics, _ := a.metricStore.GetAll()
	for _, m := range metrics {
		ch <- m
	}
	log.Println("sendMetric() finish")
}

func (a *Agent) startSend(ctx context.Context, wg *sync.WaitGroup, ch <-chan models.Metric) {
	defer wg.Done()
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case v := <-ch:
			log.Println("StartSend() start")
			a.transport.SendMetric(v)
		}
	}
	log.Println("StartSend() finish")
}

func (a *Agent) Start(ctx context.Context, cF context.CancelFunc, wg *sync.WaitGroup) {

	mCh := make(chan models.Metric, 5)
	wg.Add(1)
	go a.startSend(ctx, wg, mCh)

	tickerPollInter := time.NewTicker(time.Duration(a.pollInterval) * time.Second)
	tickerRepInter := time.NewTicker(time.Duration(a.reportInterval) * time.Second)
	cycleCounter := 1
	wg.Add(1)
	go func() {
		defer wg.Done()
	LOOP:
		for {

			select {
			case <-ctx.Done():
				close(mCh)
				break LOOP
			case <-tickerPollInter.C:
				a.updateMetric()
			case <-tickerRepInter.C:
				a.sendMetric(mCh)
				cycleCounter += 1
				if cycleCounter > 2 {
					log.Println(cycleCounter)
					err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
					if err != nil {
						return
					}
					//ctx.Done()
					//close(mCh)
					break LOOP

				}
			}
		}
	}()
}
