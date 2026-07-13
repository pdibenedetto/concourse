package kubernetes

import (
	"code.cloudfoundry.org/lager/v3"
	"k8s.io/client-go/kubernetes"

	"github.com/concourse/concourse/atc/creds"
)

type kubernetesFactory struct {
	logger lager.Logger

	client          kubernetes.Interface
	namespacePrefix string
	sharedPath      string
}

func NewKubernetesFactory(logger lager.Logger, client kubernetes.Interface, namespacePrefix, sharedPath string) *kubernetesFactory {
	factory := &kubernetesFactory{
		logger:          logger,
		client:          client,
		namespacePrefix: namespacePrefix,
		sharedPath:      sharedPath,
	}

	return factory
}

func (factory *kubernetesFactory) NewSecrets() creds.Secrets {
	return &Secrets{
		logger:          factory.logger,
		client:          factory.client,
		namespacePrefix: factory.namespacePrefix,
		sharedPath:      factory.sharedPath,
	}
}
