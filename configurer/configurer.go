package configurer

import (
	"errors"

	"github.com/cloudfoundry-incubator/cf-tcp-router/configurer/haproxy"
	"github.com/cloudfoundry-incubator/cf-tcp-router/models"
	"github.com/pivotal-golang/lager"
)

const (
	HaProxyConfigurer = "HAProxy"
)

//go:generate counterfeiter -o fakes/fake_configurer.go . RouterConfigurer
type RouterConfigurer interface {
	Configure(routingTable models.RoutingTable) error
}

func NewConfigurer(logger lager.Logger, tcpLoadBalancer string, tcpLoadBalancerBaseCfg string, tcpLoadBalancerCfg string, scriptRunner haproxy.ScriptRunner) RouterConfigurer {
	switch tcpLoadBalancer {
	case HaProxyConfigurer:
		routerHostInfo, err := haproxy.NewHaProxyConfigurer(logger, tcpLoadBalancerBaseCfg, tcpLoadBalancerCfg, scriptRunner)
		if err != nil {
			logger.Fatal("could not create tcp load balancer",
				err,
				lager.Data{"tcp_load_balancer": tcpLoadBalancer})
			return nil
		}
		return routerHostInfo
	default:
		logger.Fatal("not-supported", errors.New("unsupported tcp load balancer"), lager.Data{"tcp_load_balancer": tcpLoadBalancer})
		return nil
	}
}
