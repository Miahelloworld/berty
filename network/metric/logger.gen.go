// Code generated by berty.tech/network/.scripts/generate-logger.sh

package metric

import "go.uber.org/zap"

func logger() *zap.Logger {
	return zap.L().Named("network.metric")
}