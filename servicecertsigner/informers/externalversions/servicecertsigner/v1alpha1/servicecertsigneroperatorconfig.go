// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	apiservicecertsignerv1alpha1 "github.com/openshift/api/servicecertsigner/v1alpha1"
	versioned "github.com/openshift/client-go/servicecertsigner/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/servicecertsigner/informers/externalversions/internalinterfaces"
	servicecertsignerv1alpha1 "github.com/openshift/client-go/servicecertsigner/listers/servicecertsigner/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceCertSignerOperatorConfigInformer provides access to a shared informer and lister for
// ServiceCertSignerOperatorConfigs.
type ServiceCertSignerOperatorConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() servicecertsignerv1alpha1.ServiceCertSignerOperatorConfigLister
}

type serviceCertSignerOperatorConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewServiceCertSignerOperatorConfigInformer constructs a new informer for ServiceCertSignerOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceCertSignerOperatorConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceCertSignerOperatorConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredServiceCertSignerOperatorConfigInformer constructs a new informer for ServiceCertSignerOperatorConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceCertSignerOperatorConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecertsignerV1alpha1().ServiceCertSignerOperatorConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecertsignerV1alpha1().ServiceCertSignerOperatorConfigs().Watch(context.TODO(), options)
			},
		},
		&apiservicecertsignerv1alpha1.ServiceCertSignerOperatorConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceCertSignerOperatorConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceCertSignerOperatorConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceCertSignerOperatorConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiservicecertsignerv1alpha1.ServiceCertSignerOperatorConfig{}, f.defaultInformer)
}

func (f *serviceCertSignerOperatorConfigInformer) Lister() servicecertsignerv1alpha1.ServiceCertSignerOperatorConfigLister {
	return servicecertsignerv1alpha1.NewServiceCertSignerOperatorConfigLister(f.Informer().GetIndexer())
}
