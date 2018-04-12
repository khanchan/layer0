package config

import (
	"time"

	"github.com/quintilesims/layer0/common/models"
)

const (
	DefaultAWSRegion               = "us-west-2"
	DefaultTimeBetweenRequests     = time.Millisecond * 10
	DefaultLockExpiry              = time.Hour * 1
	DefaultPort                    = 9090
	DefaultEnvironmentInstanceType = "t2.small"
	DefaultEnvironmentOS           = "linux"
	DefaultServiceScale            = 1
	DefaultLoadBalancerType        = models.ApplicationLoadBalancerType
	DefaultTargetGroupPort         = 80
	DefaultTargetGroupProtocol     = "HTTP"

	// https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	DefaultFargatePlatformVersion = "1.1.0"
)

func DefaultLoadBalancerHealthCheck() models.HealthCheck {
	return models.HealthCheck{
		Target:             "TCP:80",
		Path:               "/",
		Interval:           30,
		Timeout:            5,
		HealthyThreshold:   2,
		UnhealthyThreshold: 2,
	}
}
func DefaultLoadBalancerPort() models.Port {
	return models.Port{
		ContainerPort: 80,
		HostPort:      80,
		Protocol:      "HTTP",
	}
}