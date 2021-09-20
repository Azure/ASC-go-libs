package instrumentation

import (
	"time"

	"github.com/sirupsen/logrus"
)

// interface for Heartbeat sender
type HeartbeatSender interface {
	// Start sending heartbeat
	start()
}

// HeartbeatSenderImpl holds the needed data for sending heartbeat
type HeartbeatSenderImpl struct {
	log          *logrus.Entry
	metricSender PlatformMetricSubmitter
	metric       Metric
	duration     time.Duration
}

// newHeartbeatSender Creates a new object which publish heartbeat every duration time
func newHeartbeatSender(log *logrus.Entry, metricSubmitter PlatformMetricSubmitter, duration time.Duration) HeartbeatSender {
	heartbeatSender := &HeartbeatSenderImpl{
		log:          log,
		metricSender: metricSubmitter,
		metric:       NewDimensionlessMetric("Heartbeat"),
		duration:     duration,
	}

	return heartbeatSender
}

// Start sending heartbeat
func (heartbeatSender *HeartbeatSenderImpl) start() {
	go heartbeatSender.sendHeartbeats()
}

func (heartbeatSender *HeartbeatSenderImpl) sendHeartbeats() {
	for range time.Tick(heartbeatSender.duration) {
		heartbeatSender.log.Trace("Heartbeat")
		heartbeatSender.metricSender.SendMetricToPlatform(1, heartbeatSender.metric)
	}
}
