/*
Copyright (c) 2024-2024 VMware, Inc. All Rights Reserved.

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

package cache

import (
	"context"

	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/watch"
)

// ListOptions are used for listing and watching objects. Please note the
// MaxObjects property is ignored for watches.
type ListOptions struct {
	types.PropertyFilterSpec
	types.RetrieveOptions
}

// Lister is any object that knows how to perform an initial list.
type Lister interface {
	// List should return a list of managed entities.
	List(ctx context.Context, options ListOptions) ([]mo.Reference, error)
}

// Watcher is any object that knows how to start a watch on a managed entity.
type Watcher interface {
	// Watch should begin a watch at the specified version.
	Watch(ctx context.Context, options ListOptions) (watch.Watcher, error)
}

// ListerWatcher is any object that knows how to perform an initial list and
// start a watch on a managed entity.
type ListerWatcher interface {
	Lister
	Watcher
}

// ListFunc knows how to list managed entities.
type ListFunc func(
	ctx context.Context,
	options ListOptions) ([]mo.Reference, error)

// WatchFunc knows how to watch managed entities.
type WatchFunc func(
	ctx context.Context,
	options ListOptions) (watch.Watcher, error)

// ListWatch knows how to list and watch a set of managed entities.
// ListFunc and WatchFunc must not be nil.
type ListWatch struct {
	ListFunc  ListFunc
	WatchFunc WatchFunc
}

// Getter interface knows how to access Get method from RESTClient.
type Getter interface {
	// Get returns the requested objects with the requested properties.
	Get(
		ctx context.Context,
		propertyFilter types.PropertyFilterSpec,
		retrieveOptions types.RetrieveOptions) ([]mo.Reference, error)

	// Watch returns an interface that can be used to watch the requested
	// objects.
	Watch(
		ctx context.Context,
		propertyFilter types.PropertyFilterSpec) (watch.Watcher, error)
}

// NewListWatchFromClient creates a new ListWatch from the specified
// client and options.
func NewListWatchFromClient(c Getter, options ListOptions) *ListWatch {
	listFunc := func(ctx context.Context, options ListOptions) ([]mo.Reference, error) {
		return c.Get(ctx, options.PropertyFilterSpec, options.RetrieveOptions)
	}
	watchFunc := func(ctx context.Context, options ListOptions) (watch.Watcher, error) {
		return c.Watch(ctx, options.PropertyFilterSpec)
	}
	return &ListWatch{ListFunc: listFunc, WatchFunc: watchFunc}
}

// List a set of resources
func (lw *ListWatch) List(
	ctx context.Context,
	options ListOptions) ([]mo.Reference, error) {

	// ListWatch is used in Reflector, which already supports pagination.
	// Don't paginate here to avoid duplication.
	return lw.ListFunc(ctx, options)
}

// Watch a set of resources.
func (lw *ListWatch) Watch(
	ctx context.Context,
	options ListOptions) (watch.Watcher, error) {

	return lw.WatchFunc(ctx, options)
}
