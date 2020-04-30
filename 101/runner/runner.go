package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//Runner在给定的超时时间内执行一组任务，并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	//interrupt通道报告从操作系统发送的信号
	interrupt chan os.Signal

	//complete通道报告处理任务已经完成
	complete chan error

	//timeout报告处理任务已经超实惠
	timeout <-chan time.Time

	//tasks持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

//任务执行超时时返回
var ErrorTimeOut = errors.New("接收到时间超时")

//收到操作系统的事件的时候返回
var ErrorInterrupt = errors.New("操作系统中断")

func New(duration time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(duration),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrorTimeOut

	}

}

func (r *Runner) run() error {
	//见过遍历变量的，但是第一次见遍历函数的，这样的话，只能是执行函数之前判断，万一我在函数执行过程中取消了呢？
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrorInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	//如果是有缓存的通道的话，假如通道里面没有东西，这个gorountine会不会等待，还是像下面这个代码一样，直接走了
	//会的，会直接阻塞
	case <-r.interrupt:
		//会保证 r 不会接受任何信号
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
