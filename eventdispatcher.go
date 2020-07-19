package main

import (
	"reflect"
	"sync"
)

type Event interface {
	Copy() Event
}

type EventDispatcher struct {
	lock      *sync.RWMutex
	listeners map[reflect.Type][]reflect.Value
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		lock:      &sync.RWMutex{},
		listeners: make(map[reflect.Type][]reflect.Value),
	}
}

func (d *EventDispatcher) RegisterEvent(event Event) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	eventType := reflect.TypeOf(event).Elem()
	if _, ok := d.listeners[eventType]; ok {
		return false
	}

	d.listeners[eventType] = []reflect.Value{}
	return true
}

func (d *EventDispatcher) RegisterListener(channel interface{}) bool {
	d.lock.Lock()
	defer d.lock.Unlock()

	channelValue := reflect.ValueOf(channel)
	channelType  := channelValue.Type()

	if channelType.Kind() != reflect.Chan {
		panic("attempted to register non-channel listener")
	}

	eventType := channelType.Elem()
	if arr, ok := d.listeners[eventType]; ok {
		d.listeners[eventType] = append(arr, channelValue)
		return true
	}
	return false
}

func (d *EventDispatcher) Dispatch(event Event) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()

	eventType := reflect.TypeOf(event)
	listeners, ok := d.listeners[eventType]
	if !ok {
		return false	
	}
	
	for _, listener := range listeners {
		listener.TrySend(reflect.ValueOf(event.Copy()))
	}
	return true
}