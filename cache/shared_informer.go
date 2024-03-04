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
	"time"
)

// SharedInformer provides eventually consistent linkage of its clients to the
// authoritative state of a given collection of managed entities. An object is
// identified by its managed object reference. One SharedInformer provides
// linkage to objects of a particular entity type. The linked object collection
// of a SharedInformer may be further restricted to a given parent entity, ex.
// a Folder, or via a PropertyFilter's ObjectSpec.
//
// The authoritative state of an object exists in vSphere, and an object goes
// through a strict sequence of states. An object state is either (1) present
// or (2) absent.
//
// A SharedInformer maintains a local cache --- exposed by GetStore(), by
// GetIndexer() in the case of an indexed informer, and possibly by machinery
// involved in creating and/or accessing the informer --- of the state of each
// relevant object. This cache is eventually consistent with the authoritative
// state. This means that, unless prevented by persistent communication
// problems, if ever a particular object ID X is authoritatively associated
// with a state S then for every SharedInformer I whose collection includes
// (X, S) eventually either (1) I's cache associates X with S or a later state
// of X, (2) I is stopped, or (3) the authoritative state service for X
// terminates. To be formally complete, we say that the absent state meets any
// restriction by a PropertyFilter's ObjectSpec.
//
// For a given informer and relevant object ID X, the sequence of states that
// appears in the informer's cache is a subsequence of the states
// authoritatively associated with X. That is, some states might never appear
// in the cache but ordering among the appearing states is correct. Note,
// however, that there is no promise about ordering between states seen for
// different objects.
//
// The local cache starts out empty, and gets populated and updated during
// `Run()`.
//
// As a simple example, if a collection of objects is henceforth unchanging, a
// SharedInformer is created that links to that collection, and that
// SharedInformer is `Run()` then that SharedInformer's cache eventually holds
// an exact copy of that collection (unless it is stopped too soon, the
// authoritative state service ends, or communication problems between the two
// persistently prevent achievement).
//
// As another simple example, if the local cache ever holds a non-absent state
// for some object ID and the object is eventually removed from the
// authoritative state, then eventually the object is removed from the local
// cache (unless the SharedInformer is stopped too soon, the authoritative
// state service ends, or communication problems persistently prevent the
// desired result).
//
// The keys in the Store are ManagedObjectReferences. Clients can use
// `KeyFunc(obj)` to extract the key for a given object, and `SplitKey(key)` to
// split a key into its constituent parts.
//
// Every query against the local cache is answered entirely from one snapshot of
// the cache's state. Thus, the result of a `List` call will not contain entries
// with the same namespace and name.
//
// A client is identified here by a ResourceEventHandler. For every update to
// the SharedInformer's local cache and for every client added before `Run()`,
// eventually either the SharedInformer is stopped or the client is notified of
// the update. A client added after `Run()` starts gets a startup batch of
// notifications of additions of the objects existing in the cache at the time
// that client was added; also, for every update to the SharedInformer's local
// cache after that client was added, eventually either the SharedInformer is
// stopped or that client is notified of that update.
//
// Client notifications happen after the corresponding cache update and, in the
// case of a SharedIndexInformer, after the corresponding index updates.  It is
// possible that additional cache and index updates happen before such a
// prescribed notification. For a given SharedInformer and client, the
// notifications are delivered sequentially. For a given SharedInformer, client,
// and object ID, the notifications are delivered in order.
//
// A client must process each notification promptly; a SharedInformer is not
// designed to deal well with a large backlog of notifications to deliver.
// Lengthy processing should be offloaded.
//
// A delete notification exposes the last locally known non-absent state, except
// that its ResourceVersion is replaced with a ResourceVersion in which the
// object is actually absent.
type SharedInformer interface {
	// AddEventHandler adds an event handler to the shared informer using
	// the shared informer's resync period.  Events to a single handler are
	// delivered sequentially, but there is no coordination between
	// different handlers.
	//
	// It returns a registration handle for the handler that can be used to
	// remove the handler again, or to tell if the handler is synced (has
	// seen every item in the initial list).
	AddEventHandler(
		handler ResourceEventHandler) (ResourceEventHandlerRegistration, error)

	// AddEventHandlerWithResyncPeriod adds an event handler to the shared
	// informer with the requested resync period; zero means this handler does
	// not care about resyncs.
	//
	// The resync operation consists of delivering to the handler an update
	// notification for every object in the informer's local cache; it does not
	// add any interactions with the authoritative storage. Some informers do
	// not resync at all, not even for handlers added with a non-zero
	// resyncPeriod.
	//
	// For an informer that does resyncs, and for each handler that requests
	// resyncs, that informer develops a nominal resync period that is no
	// shorter than the requested period but may be longer.The actual time
	// between any two resyncs may be longer than the nominal period because the
	// implementation takes time to do work and there may be competing load and
	// scheduling noise.
	//
	// It returns a registration handle for the handler that can be used to
	// remove the handler again and an error if the handler cannot be added.
	AddEventHandlerWithResyncPeriod(
		handler ResourceEventHandler,
		resyncPeriod time.Duration) (ResourceEventHandlerRegistration, error)

	// RemoveEventHandler removes a formerly added event handler given by
	// its registration handle.
	// This function is guaranteed to be idempotent, and thread-safe.
	RemoveEventHandler(handle ResourceEventHandlerRegistration) error

	// GetStore returns the informer's local cache as a Store.
	GetStore() Store

	// Run starts and runs the shared informer, returning after it stops.
	// The informer will be stopped when stopCh is closed.
	Run(stopCh <-chan struct{})

	// HasSynced returns true if the shared informer's store has been
	// informed by at least one full LIST of the authoritative state
	// of the informer's object collection.  This is unrelated to "resync".
	//
	// Note that this doesn't tell you if an individual handler is synced!!
	// For that, please call HasSynced on the handle returned by
	// AddEventHandler.
	HasSynced() bool

	// LastSyncResourceVersion is the resource version observed when last synced
	// with the underlying store. The value returned is not synchronized with
	// access to the underlying store and is not thread-safe.
	LastSyncResourceVersion() string

	// The WatchErrorHandler is called whenever ListAndWatch drops the
	// connection with an error. After calling this handler, the informer
	// will backoff and retry.
	//
	// The default implementation looks at the error type and tries to log
	// the error message at an appropriate level.
	//
	// There's only one handler, so if you call this multiple times, last one
	// wins; calling after the informer has been started returns an error.
	//
	// The handler is intended for visibility, not to e.g. pause the consumers.
	// The handler should return quickly - any expensive processing should be
	// offloaded.
	SetWatchErrorHandler(handler WatchErrorHandler) error

	// The TransformFunc is called for each object which is about to be stored.
	//
	// This function is intended for you to take the opportunity to
	// remove, transform, or normalize fields. One use case is to strip unused
	// metadata fields out of objects to save on RAM cost.
	//
	// Must be set before starting the informer.
	//
	// Please see the comment on TransformFunc for more details.
	SetTransform(handler TransformFunc) error

	// IsStopped reports whether the informer has already been stopped.
	// Adding event handlers to already stopped informers is not possible.
	// An informer already stopped will never be started again.
	IsStopped() bool
}

// Opaque interface representing the registration of ResourceEventHandler for
// a SharedInformer. Must be supplied back to the same SharedInformer's
// `RemoveEventHandler` to unregister the handlers.
//
// Also used to tell if the handler is synced (has had all items in the initial
// list delivered).
type ResourceEventHandlerRegistration interface {
	// HasSynced reports if both the parent has synced and all pre-sync
	// events have been delivered.
	HasSynced() bool
}

// TransformFunc allows for transforming an object before it will be processed.
// TransformFunc (similarly to ResourceEventHandler functions) should be able
// to correctly handle the tombstone of type cache.DeletedFinalStateUnknown.
//
// The most common usage pattern is to clean-up some parts of the object to
// reduce component memory usage if a given component doesn't care about them.
//
// Note that TransformFunc is called while inserting objects into the
// notification queue and is therefore extremely performance sensitive; please
// do not do anything that will take a long time.
type TransformFunc func(interface{}) (interface{}, error)
