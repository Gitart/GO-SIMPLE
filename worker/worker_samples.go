package main

import (
	"fmt"
	"sync"
	"time"
)

/*
How to use.
Example:

const Workers = 3

type Task struct {
	req *http.Request
}

func (t *Task) Exec() {
	res, err := http.DefaultClient.Do(t.req)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	io.Copy(io.Discard, res.Body)
	defer res.Body.Close()
}

func main() {

	disp := NewDefaultDispatcher(Workers)
	done := disp.Dispatch()
	go func(d *Dispatcher) {
		close(done)
		d.Wait()
	}(disp)

	var reqID uint64 = 1
	for {
		time.Sleep(500 * time.Millisecond)
		
		req, err := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/in?q="+fmt.Sprintf("%d", reqID), nil)
		if err != nil {
			log.Printf("%v\n", err)
			return
		}

		disp.JobQueue <- &Task{req: req}

		reqID++
	}
}

*/
package main

import (
	"errors"
	"log"
	"runtime"
	"sync"
	"time"
)

var (
	ErrInitWorkers  = errors.New("[Dispatcher] You must call func 'InitWorkers'")
	ErrInitJobQueue = errors.New("[Dispatcher] You must call func 'InitJobQueue'")
)

const (
	WarnAlreadyInit = "[Dispatcher] Warning. Already init"
)

var DispatcherWatchdog = 10 * time.Second

// Jober интерфейс задания
type Jober interface {
	Exec()
}

type (
	worker struct {
		registerFunc func(chan Jober) // Обработчик регистрации воркера
		recoveryFunc func(*worker)    // Обработчик паники
		quit         chan struct{}
		jobChannel   chan Jober
	}

	Dispatcher struct {
		JobQueue    chan Jober // Канал для задач
		workers     map[*worker]struct{}
		workersPool chan chan Jober // Канал для воркеров
		wg          sync.WaitGroup
	}
)

func NewDefaultDispatcher(num int) *Dispatcher {
	var d Dispatcher
	d.InitJobQueue(num)
	d.InitWorkers(num)
	return &d
}

// Регистрирует воркер на получение заданий
func (d *Dispatcher) registerWorker(c chan Jober) {
	d.workersPool <- c
}

// Отслеживает паники воркеров
func (d *Dispatcher) recoverWorker(w *worker) {
	d.StopWorker(w)

	if err := recover(); err != nil {
		log.Printf("[Dispatcher] worker panic: %v\n", err)

		// Пересоздаем воркер в случае паники, чтобы пул воркеров не опустел
		d.StartWorker()
	}
}

// InitWorkers инициализирует пул воркеров и осуществляет запуск воркеров
func (d *Dispatcher) InitWorkers(n int) {
	if d.workers != nil {
		log.Printf(WarnAlreadyInit)
		return
	}

	d.workersPool = make(chan chan Jober, n)

	d.workers = make(map[*worker]struct{}, n)

	for i := 0; i < n; i++ {
		d.StartWorker()
	}
}

// InitJobQueue создает очередь заданий для воркеров
func (d *Dispatcher) InitJobQueue(n int) chan Jober {
	if d.JobQueue == nil {
		d.JobQueue = make(chan Jober, n)
	}
	return d.JobQueue
}

// StartWorker запускает воркер
func (d *Dispatcher) StartWorker() {
	if d.workers == nil {
		panic(ErrInitWorkers)
	}

	worker := worker{
		recoveryFunc: d.recoverWorker,
		registerFunc: d.registerWorker,
		quit:         make(chan struct{}),
		jobChannel:   make(chan Jober),
	}

	d.workers[&worker] = struct{}{}
	worker.start()
	d.wg.Add(1)
}

// StopWorker останавливает воркер
func (d *Dispatcher) StopWorker(w *worker) {
	if d.workers == nil {
		panic(ErrInitWorkers)
	}

	if _, ok := d.workers[w]; ok {
		d.wg.Done()
		w.stop()
		delete(d.workers, w)
	}
}

// DestroyWorkers уничтожает все воркеры
func (d *Dispatcher) DestroyWorkers() {
	for w := range d.workers {
		go d.StopWorker(w)
	}
}

// Dispatch запускает основной цикл управления.
// Канал enough следует закрыть снаружи, чтобы дать понять, что все закончено и воркеры больше не нужны.
func (d *Dispatcher) Dispatch() chan struct{} {
	if d.JobQueue == nil {
		panic(ErrInitJobQueue)
	}

	if d.workers == nil {
		panic(ErrInitWorkers)
	}

	var enough = make(chan struct{})

	go func() {
		var job Jober
		var jobChan chan Jober
		var ok bool

		defer d.DestroyWorkers()

		watchdog := time.NewTicker(DispatcherWatchdog)
		defer watchdog.Stop()

		for {
			select {
			// Расталкиваем задачи по воркерам
			case job, ok = <-d.JobQueue:
				if !ok {
					return
				}

				jobChan, ok = <-d.workersPool
				jobChan <- job

			case <-watchdog.C:
				// Все воркеры в ожидании
				if len(d.workersPool) == len(d.workers) {
					select {
					case <-enough:
						runtime.Goexit()
					default:
					}
				}
			}
		}
	}()

	return enough
}

// Wait дожидается окончания диспатчера
func (d *Dispatcher) Wait() {
	d.wg.Wait()
}

// Start запускает воркер на исполнение
func (w *worker) start() {
	go func() {

		defer close(w.jobChannel)
		defer w.recoveryFunc(w)

		for {
			select {
			case <-w.quit:
				runtime.Goexit()
			default:
			}

			// Регистрируемся на получение заданий
			w.registerFunc(w.jobChannel)

			job := <-w.jobChannel
			job.Exec()
		}
	}()
}

// Stop останавливает работу воркера
func (w *worker) stop() {
	w.quit <- struct{}{}
	close(w.quit)
}

