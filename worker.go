package main

import "fmt"

type worker struct {
	queue chan *job
}

func newWorker(size int) *worker {
	var queue chan *job
	if size <= 0 {
		queue = make(chan *job)
	} else {
		queue = make(chan *job, size)
	}
	w := &worker{queue: queue}
	w.run()
	return w
}

func (w *worker) run() {
	go func() {
		for {
			select {
			case job := <-w.queue:
				go job.Run()
			}
		}
	}()
}

func (w *worker) AddJob(fn func() error) JobStatus {
	job := newJob(fn)
	w.queue <- job
	return job
}

type Status int

const (
	Waiting Status = 1 << iota
	Running
	Success
	Error
)

func newJob(fn func() error) *job {
	return &job{
		jobStatus: &jobStatus{
			errChan:     make(chan error, 1),
			successChan: make(chan bool, 1),
		},
		fn: fn,
	}
}

type Job interface {
	Run()
}

type job struct {
	*jobStatus
	fn func() error
}

func (j *job) Run() {
	defer func() {
		if r := recover(); r != nil {
			var err error
			if recErr, ok := r.(error); ok {
				err = recErr
			} else {
				err = fmt.Errorf("%v", r)
			}
			j.jobStatus.setError(err)
		}
	}()
	err := j.fn()
	if err != nil {
		j.jobStatus.setError(err)
	} else {
		j.jobStatus.setSuccess(true)
	}
}

type JobStatus interface {
	// Success blocks until the job is done and returns whether the run was successful or not
	// If Success returns true, error is always nil
	// If Success returns false, the returned error can be non-nil
	Success() (bool, error)
}

type jobStatus struct {
	errChan     chan error
	err         error
	successChan chan bool
	success     bool
}

func (j *jobStatus) setSuccess(success bool) {
	j.successChan <- success
}

func (j *jobStatus) setError(err error) {
	j.successChan <- false
	j.errChan <- err
}

// Success blocks until the job is done and returns whether the run was successful or not
// If Success returns true, error is always nil
// If Success returns false, the returned error can be non-nil
func (j *jobStatus) Success() (bool, error) {
	if len(j.successChan) != 0 {
		j.success = <-j.successChan
	}
	if len(j.errChan) != 0 {
		j.err = <-j.errChan
	}
	return j.success, j.err
}
