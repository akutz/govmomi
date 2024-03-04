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
	"fmt"

	"github.com/vmware/govmomi/sets"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

// Indexer extends Store with multiple indices and restricts each accumulator to
// simply hold the current object (and be empty after Delete).
//
// There are three kinds of strings here:
//  1. a storage key, as defined in the Store interface,
//  2. a name of an index, and
//  3. an "indexed value", which is produced by an IndexFunc and
//     can be a field value or any other string computed from the object.
type Indexer interface {
	Store
	// Index returns the stored objects whose set of indexed values intersects
	// the set of indexed values of the given object, for the named index
	Index(indexName string, obj interface{}) ([]interface{}, error)
	// IndexKeys returns the storage keys of the stored objects whose set of
	// indexed values for the named index includes the given indexed value
	IndexKeys(indexName, indexedValue string) ([]string, error)
	// ListIndexFuncValues returns all the indexed values of the given index
	ListIndexFuncValues(indexName string) []string
	// ByIndex returns the stored objects whose set of indexed values for the
	// named index includes the given indexed value
	ByIndex(indexName, indexedValue string) ([]interface{}, error)
	// GetIndexers return the indexers
	GetIndexers() Indexers
	// AddIndexers adds more indexers to this store. This supports adding
	// indexers after the store already has items.
	AddIndexers(newIndexers Indexers) error
}

// IndexFunc knows how to compute the set of indexed values for an object.
type IndexFunc func(obj interface{}) ([]string, error)

// IndexFuncToKeyFuncAdapter adapts an indexFunc to a keyFunc. This is only
// useful if your index function returns unique values for every object. This
// conversion can create errors when more than one key is found. You should
// prefer to make proper key and index functions.
func IndexFuncToKeyFuncAdapter(indexFunc IndexFunc) KeyFunc {
	return func(obj interface{}) (string, error) {
		indexKeys, err := indexFunc(obj)
		if err != nil {
			return "", err
		}
		if len(indexKeys) > 1 {
			return "", fmt.Errorf("too many keys: %v", indexKeys)
		}
		if len(indexKeys) == 0 {
			return "", fmt.Errorf("unexpected empty indexKeys")
		}
		return indexKeys[0], nil
	}
}

var emptyReference types.ManagedObjectReference

// ParentIndexFunc is a default index function that indexes based on an object's
// parent.
func ParentIndexFunc(obj mo.Reference) (types.ManagedObjectReference, error) {
	ref := obj.Reference()
	if ref == emptyReference {
		return emptyReference, fmt.Errorf("object has no reference")
	}
	me, ok := obj.(mo.IsManagedEntity)
	if !ok {
		return emptyReference, fmt.Errorf("object is not ManagedEntity")
	}
	parent := me.GetManagedEntity().Parent
	if parent == nil {
		return emptyReference, fmt.Errorf("object has no parent")
	}
	return *parent, nil
}

// Index maps the indexed value to a set of keys in the store that match on that value
type Index map[string]sets.Set[string]

// Indexers maps a name to an IndexFunc
type Indexers map[string]IndexFunc

// Indices maps a name to an Index
type Indices map[string]Index
