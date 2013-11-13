package worker


import (
  "mas/logger"
  "sync"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


// WorkerPool implementation.
type WorkerPool struct {
  m_NbWorker int
  m_WaitGroup *sync.WaitGroup
  m_ChannelIn chan IJob
}


// NewWorkerPool will create an instance of a worker pool.
// p_NbWorker number of worker to spawn.
// It returns a pointer to the pool.
func NewWorkerPool(p_NbWorker int) *WorkerPool {
  workerPool := WorkerPool{}
  workerPool.m_NbWorker = p_NbWorker
  workerPool.m_WaitGroup = new(sync.WaitGroup)
  workerPool.m_ChannelIn = make(chan IJob)

  for i := 0; i < workerPool.m_NbWorker; i++ {
    workerPool.m_WaitGroup.Add(1)
    go Worker(i, workerPool.m_WaitGroup, workerPool.m_ChannelIn)
  }

  return &workerPool
}


// Do give a job to one of the worker.
// p_Job the job to be executed.
func (w *WorkerPool) Do(p_Job IJob) {
  w.m_ChannelIn <- p_Job
}


// Wait will close and wait until all workers are done.
func (w *WorkerPool) Wait() {
  close(w.m_ChannelIn)
  w.m_WaitGroup.Wait()
  s_Logger.Debug("All workers end gracefully.")
}


// Worker itself.
// p_Id id of the worker.
// p_WaitGroup pointer to the WaitGroup instance.
// p_ChannelIn the channel where to jobs come from.
func Worker(p_Id int, p_WaitGroup *sync.WaitGroup, p_ChannelIn chan IJob) {
  defer p_WaitGroup.Done()
  s_Logger.Debug("Worker", p_Id, "started")

  for job := range p_ChannelIn {
    job.Do()
  }
}
