/*


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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	schedulingv1beta1 "gpu-discovery/apis/scheduling/v1beta1"
	versioned "gpu-discovery/generated/gpus/clientset/versioned"
	internalinterfaces "gpu-discovery/generated/gpus/informers/externalversions/internalinterfaces"
	v1beta1 "gpu-discovery/generated/gpus/listers/scheduling/v1beta1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// GPUInformer provides access to a shared informer and lister for
// GPUs.
type GPUInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.GPULister
}

type gPUInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewGPUInformer constructs a new informer for GPU type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewGPUInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredGPUInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredGPUInformer constructs a new informer for GPU type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredGPUInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1beta1().GPUs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SchedulingV1beta1().GPUs().Watch(context.TODO(), options)
			},
		},
		&schedulingv1beta1.GPU{},
		resyncPeriod,
		indexers,
	)
}

func (f *gPUInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredGPUInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *gPUInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&schedulingv1beta1.GPU{}, f.defaultInformer)
}

func (f *gPUInformer) Lister() v1beta1.GPULister {
	return v1beta1.NewGPULister(f.Informer().GetIndexer())
}
