// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v3

import (
	"context"
	time "time"

	projectcalicov3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	clientset "github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/projectcalico/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v3 "github.com/projectcalico/api/pkg/client/listers_generated/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StagedGlobalNetworkPolicyInformer provides access to a shared informer and lister for
// StagedGlobalNetworkPolicies.
type StagedGlobalNetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v3.StagedGlobalNetworkPolicyLister
}

type stagedGlobalNetworkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewStagedGlobalNetworkPolicyInformer constructs a new informer for StagedGlobalNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStagedGlobalNetworkPolicyInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStagedGlobalNetworkPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredStagedGlobalNetworkPolicyInformer constructs a new informer for StagedGlobalNetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStagedGlobalNetworkPolicyInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().StagedGlobalNetworkPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoV3().StagedGlobalNetworkPolicies().Watch(context.TODO(), options)
			},
		},
		&projectcalicov3.StagedGlobalNetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *stagedGlobalNetworkPolicyInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStagedGlobalNetworkPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *stagedGlobalNetworkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcalicov3.StagedGlobalNetworkPolicy{}, f.defaultInformer)
}

func (f *stagedGlobalNetworkPolicyInformer) Lister() v3.StagedGlobalNetworkPolicyLister {
	return v3.NewStagedGlobalNetworkPolicyLister(f.Informer().GetIndexer())
}
