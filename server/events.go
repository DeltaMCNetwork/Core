package server

import "reflect"

type InjectionManager struct {
	injectors map[string][]*injecter
}

type injecter struct {
	tname  string
	action reflect.Value
}

type IEvent interface {
}

func CreateInjectionManager() *InjectionManager {
	return &InjectionManager{
		injectors: make(map[string][]*injecter, 0),
	}
}

func (manager *InjectionManager) Register(event interface{}) {
	firstArg := reflect.TypeOf(event).In(0)
	name := firstArg.Elem().Name()

	manager.injectors[name] = append(manager.injectors[name], &injecter{
		tname:  name,
		action: reflect.ValueOf(event),
	})
}

func (manager *InjectionManager) Post(event IEvent) {
	if event == nil {
		Error("Event was nil!")
	}

	name := reflect.ValueOf(event).Elem().Type().Name()
	injections := manager.injectors[name]
	values := make([]reflect.Value, 1)
	values[0] = reflect.ValueOf(event)

	for _, inj := range injections {
		inj.action.Call(values)
	}
}

//test event

type InitializationEvent struct {
	IEvent
	Time int
}

type ServerTickEvent struct {
	Count int
}
