package main

type Job interface {
    Do() error
}

type Worker struct {
    JobChannel JobChan
    quit       chan bool
}

var (
    JobQueue          JobChan
    WorkerPool        WorkerChan
)
func main() 
	// define job channel
	type JobChan chan Job

	// define worker channer
	type WorkerChan chan JobChan
	const MAX_QUEUE_SIZE = 100
	var queue = make(chan job,MAX_QUEUE_SIZE)
	
	job := &Job{request}

	queue <- job
}
func (w *Worker) Start() {
    go func() {
        for {
            // regist current job channel to worker pool
            WorkerPool <- w.JobChannel
            select {
            case job := <-w.JobChannel:
                if err := job.Do(); err != nil {
                    fmt.printf("excute job failed with err: %v", err)
                }
            // recieve quit event, stop worker
            case <-w.quit:
                return
            }
        }
    }()
}