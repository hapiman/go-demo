package log_demo

import "testing"

func TestZapLogger(t *testing.T) {
	t.Log("------log---------")
	MainLogger.Debug("hello main Debug")
	MainLogger.Info("hello main Info")
	GatewayLogger.Debug("Hi Gateway Im Debug")
	GatewayLogger.Info("Hi Gateway  Im Info")
}
