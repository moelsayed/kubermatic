package apiserver

import (
	"github.com/kubermatic/kubermatic/api/pkg/resources"
	"github.com/kubermatic/kubermatic/api/pkg/resources/certificates"
	"github.com/kubermatic/kubermatic/api/pkg/resources/certificates/triple"
)

type etcdClientCertificateCreatorData interface {
	GetRootCA() (*triple.KeyPair, error)
}

// EtcdClientCertificateCreator returns a function to create/update the secret with the client certificate for authenticating against etcd
func EtcdClientCertificateCreator(data etcdClientCertificateCreatorData) resources.NamedSecretCreatorGetter {
	return certificates.GetClientCertificateCreator(
		resources.ApiserverEtcdClientCertificateSecretName,
		"apiserver",
		nil,
		resources.ApiserverEtcdClientCertificateCertSecretKey,
		resources.ApiserverEtcdClientCertificateKeySecretKey,
		data.GetRootCA)
}
