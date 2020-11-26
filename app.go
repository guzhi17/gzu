package gzu


import (
	"log"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	GlobalChExit = make(chan struct{})

	closeChOnce int32 = 0

	once sync.Once
)


func AppClose(){
	if atomic.AddInt32(&closeChOnce, 1) == 1{
		close(GlobalChExit)
	}
}

func AppIsClosed() bool{
	return atomic.LoadInt32(&closeChOnce)>0
}

func appDoInit(){

	sigs := make(chan os.Signal,1)
	//syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	//done := make(chan bool,1)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		AppEvent(E_EXIT, nil, time.Second)
		AppClose()
		log.Println(sig)
		//done <- true
		time.Sleep(time.Second*8)
		os.Exit(2)
	}()
}


type AppRunnable interface {
	AppRun()
}

type AppSync func()

func AppInit()  {
	AppEvent(E_INIT, nil, -1)
	log.Println("App Initialized!!!")
}

func AppRun(runners...interface{})  {

	once.Do(func() {
		AppInit()
		appDoInit()
	})

	//we are going to modify the
	for _, runner := range runners{
		switch r := runner.(type) {
		case AppRunnable:
			go r.AppRun()
		case AppSync:
			r()
		case func():
			go r()
		default:
			log.Println("not a runnable type", r)
		}
	}
	//<- GlobalChExit:
	for {
		select {
		case <- GlobalChExit:
			goto EXIT
		}
	}
EXIT:
	AppClose()
}



//--------------------------------------------------------------------------

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
func AppEventRegister(e interface{}, v...IAppEventHandler){
	internalEventsHandlers.Register(e, v...)
}
func AppEventRegisterPrepend(e interface{}, v...IAppEventHandler){
	internalEventsHandlers.RegisterPrepend(e, v...)
}

//remove first
func AppEventUnRegister(e interface{}, v IAppEventHandler)bool{
	return internalEventsHandlers.UnRegister(e, v)
}


func AppEvent(e, d interface{}, timeout time.Duration) {
	handlers := internalEventsHandlers.Clone(e)
	if len(handlers) < 1{
		//log.Println("no handler for ", e)
		return
	}
	//log.Println("handler for ", e, timeout)
	for _, v := range handlers{
		if timeout == 0{
			//no wait
			go appEventHandle(v, e, d)
		}else if timeout < 0{
			//wait forever
			appEventHandle(v, e, d)
		}else{
			c := make(chan bool)
			go func() {
				defer func() {
					close(c)
				}()
				appEventHandle(v, e, d)
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

func appEventHandle(h IAppEventHandler, e, d interface{})  {
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


func AppOnExit(exit func())  {
	//filo
	AppEventRegisterPrepend(E_EXIT, &HandlerFunc0{f:exit})
}

func AppOnInit(c func())  {
	//fifo
	AppEventRegister(E_INIT, &HandlerFunc0{f:c})
}

func AppOnConfigChanged(f func(interface{}))  {
	AppEventRegister(E_CONFIG_CHANGED, &HandlerFunc1{f:f})
}
//------------------------------------------------------------