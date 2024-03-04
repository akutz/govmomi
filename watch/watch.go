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

package watch

import "github.com/vmware/govmomi/vim25/types"

// Watcher can be implemented by anything that knows how to watch and report
// changes.
type Watcher interface {
	// Stop stops watching. Will close the channel returned by ResultChan()
	// Releases any resources used by the watch.
	Stop()

	// ResultChan returns a chan which will receive all the events. If an error
	// occurs or Stop() is called, the implementation will close this channel
	// and release any resources used by the watch.
	ResultChan() <-chan Event
}

// EventType defines the possible types of events.
type EventType string

const (
	Added    EventType = "ADDED"
	Modified EventType = "MODIFIED"
	Deleted  EventType = "DELETED"
)

// Event represents a single event to a watched resource.
type Event struct {
	Type EventType

	// Object is:
	//  * If Type is Added or Modified: the new state of the object.
	//  * If Type is Deleted: the state of the object immediately before
	//    deletion.
	Object types.ObjectUpdate
}
