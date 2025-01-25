// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apiconsolev1 "github.com/openshift/api/console/v1"
	versioned "github.com/openshift/client-go/console/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/console/informers/externalversions/internalinterfaces"
	consolev1 "github.com/openshift/client-go/console/listers/console/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConsolePluginInformer provides access to a shared informer and lister for
// ConsolePlugins.
type ConsolePluginInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() consolev1.ConsolePluginLister
}

type consolePluginInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConsolePluginInformer constructs a new informer for ConsolePlugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConsolePluginInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConsolePluginInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConsolePluginInformer constructs a new informer for ConsolePlugin type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConsolePluginInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsolePlugins().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsolePlugins().Watch(context.TODO(), options)
			},
		},
		&apiconsolev1.ConsolePlugin{},
		resyncPeriod,
		indexers,
	)
}

func (f *consolePluginInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConsolePluginInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *consolePluginInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiconsolev1.ConsolePlugin{}, f.defaultInformer)
}

func (f *consolePluginInformer) Lister() consolev1.ConsolePluginLister {
	return consolev1.NewConsolePluginLister(f.Informer().GetIndexer())
}
