package app

import (
	"github.com/guzhi17/gzu/atom"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

var (
	GlobalChExit = make(chan struct{})

	closeChOnce int32 = 0

	once atom.Once
	//once sync.Once
)


func Close(){
	if atomic.AddInt32(&closeChOnce, 1) == 1{
		close(GlobalChExit)
	}
}

func IsClosed() bool{
	return atomic.LoadInt32(&closeChOnce)>0
}

func doInitApp(){

	sigs := make(chan os.Signal,1)
	//syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	//done := make(chan bool,1)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		Event(E_EXIT, nil, time.Second)
		Close()
		log.Println(sig)
		//done <- true
		time.Sleep(time.Second*8)
		os.Exit(2)
	}()
}


type Runnable interface {
	Run()
}

type Sync func()

func Init()  {
	Event(E_INIT, nil, -1)
	log.Println("App Initialized!!!")
}

func Run(runners...interface{})  {

	once.Do(func() {
		Init()
		doInitApp()
	})

	//we are going to modify the
	for _, runner := range runners{
		switch r := runner.(type) {
		case Runnable:
			go r.Run()
		case Sync:
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
	Close()
}