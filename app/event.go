package app

import (
	"log"
	"sync"
	"time"
)


const(
	E_NONE = iota
	E_EXIT
	E_INIT
	E_CONFIG_CHANGED
)

type IAppEventHandler interface {
	Handle(e interface{}, d interface{})
}

type EventHandlers []IAppEventHandler

type EventsHandlers struct {
	sync.RWMutex
	handlers map[interface{}]EventHandlers
}

var(
	internalEventsHandlers = &EventsHandlers{
		handlers: map[interface{}]EventHandlers{},
	}
)

//panic if name the same
func (h *EventsHandlers)Register(e interface{}, v...IAppEventHandler){
	h.Lock()
	defer h.Unlock()
	h.handlers[e] = append(h.handlers[e], v...)
}

func (h *EventsHandlers)RegisterPrepend(e interface{}, v...IAppEventHandler){
	h.Lock()
	defer h.Unlock()
	h.handlers[e] = append(v, h.handlers[e]...)
}

//remove first
func (h *EventsHandlers)UnRegister(e interface{}, v IAppEventHandler)bool{
	h.Lock()
	defer h.Unlock()
	if old, ok := h.handlers[e]; ok{
		for i, eh := range old{
			if eh == v{
				h.handlers[e] = append(old[:i], old[i+1:]...)
				return true
			}
		}
	}
	return false
}

//no copy, so its not
func (h *EventsHandlers)Get(e interface{})[]IAppEventHandler{
	h.RLock()
	defer h.RUnlock()
	if old, ok := h.handlers[e]; ok{
		return old
	}
	return nil
}

//no copy, so its not
func (h *EventsHandlers)Clone(e interface{})(r []IAppEventHandler){
	h.RLock()
	defer h.RUnlock()
	if old, ok := h.handlers[e]; ok{
		if len(old) > 0{
			r = make([]IAppEventHandler, len(old))
			copy(r, old)
			return r
		}
	}
	return nil
}

//panic if name the same
func Register(e interface{}, v...IAppEventHandler){
	internalEventsHandlers.Register(e, v...)
}
func RegisterPrepend(e interface{}, v...IAppEventHandler){
	internalEventsHandlers.RegisterPrepend(e, v...)
}

//remove first
func UnRegister(e interface{}, v IAppEventHandler)bool{
	return internalEventsHandlers.UnRegister(e, v)
}


func Event(e, d interface{}, timeout time.Duration) {
	handlers := internalEventsHandlers.Clone(e)
	if len(handlers) < 1{
		//log.Println("no handler for ", e)
		return
	}
	//log.Println("handler for ", e, timeout)
	for _, v := range handlers{
		if timeout == 0{
			//no wait
			go handle(v, e, d)
		}else if timeout < 0{
			//wait forever
			handle(v, e, d)
		}else{
			c := make(chan bool)
			go func() {
				defer func() {
					close(c)
				}()
				handle(v, e, d)
			}()
			select {
			case <-time.After(timeout):
				log.Println("event handler expire")
			case <- c:
				continue
			}
		}
	}
}

func handle(h IAppEventHandler, e, d interface{})  {
	defer func() {
		recover()
	}()
	h.Handle(e, d)
}


type HandlerFunc0 struct {
	f func()
}

func (h HandlerFunc0)Handle(e, d interface{})  {
	h.f()
}


type HandlerFunc1 struct {
	f func(d interface{})
}

func (h HandlerFunc1)Handle(e, d interface{})  {
	h.f(d)
}


func OnExit(exit func())  {
	//filo
	RegisterPrepend(E_EXIT, &HandlerFunc0{f:exit})
}

func OnInit(c func())  {
	//fifo
	Register(E_INIT, &HandlerFunc0{f:c})
}

func OnConfigChanged(f func(interface{}))  {
	Register(E_CONFIG_CHANGED, &HandlerFunc1{f:f})
}