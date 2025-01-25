// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apinetworkv1 "github.com/openshift/api/network/v1"
	versioned "github.com/openshift/client-go/network/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/network/informers/externalversions/internalinterfaces"
	networkv1 "github.com/openshift/client-go/network/listers/network/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NetNamespaceInformer provides access to a shared informer and lister for
// NetNamespaces.
type NetNamespaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() networkv1.NetNamespaceLister
}

type netNamespaceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNetNamespaceInformer constructs a new informer for NetNamespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetNamespaceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetNamespaceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNetNamespaceInformer constructs a new informer for NetNamespace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetNamespaceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().NetNamespaces().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1().NetNamespaces().Watch(context.TODO(), options)
			},
		},
		&apinetworkv1.NetNamespace{},
		resyncPeriod,
		indexers,
	)
}

func (f *netNamespaceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetNamespaceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *netNamespaceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apinetworkv1.NetNamespace{}, f.defaultInformer)
}

func (f *netNamespaceInformer) Lister() networkv1.NetNamespaceLister {
	return networkv1.NewNetNamespaceLister(f.Informer().GetIndexer())
}
