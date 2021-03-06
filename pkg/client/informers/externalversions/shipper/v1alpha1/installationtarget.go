/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	shipper_v1alpha1 "github.com/bookingcom/shipper/pkg/apis/shipper/v1alpha1"
	versioned "github.com/bookingcom/shipper/pkg/client/clientset/versioned"
	internalinterfaces "github.com/bookingcom/shipper/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/bookingcom/shipper/pkg/client/listers/shipper/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InstallationTargetInformer provides access to a shared informer and lister for
// InstallationTargets.
type InstallationTargetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InstallationTargetLister
}

type installationTargetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInstallationTargetInformer constructs a new informer for InstallationTarget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInstallationTargetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInstallationTargetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInstallationTargetInformer constructs a new informer for InstallationTarget type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInstallationTargetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ShipperV1alpha1().InstallationTargets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ShipperV1alpha1().InstallationTargets(namespace).Watch(options)
			},
		},
		&shipper_v1alpha1.InstallationTarget{},
		resyncPeriod,
		indexers,
	)
}

func (f *installationTargetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInstallationTargetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *installationTargetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&shipper_v1alpha1.InstallationTarget{}, f.defaultInformer)
}

func (f *installationTargetInformer) Lister() v1alpha1.InstallationTargetLister {
	return v1alpha1.NewInstallationTargetLister(f.Informer().GetIndexer())
}
