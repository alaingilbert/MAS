package worker


import (
  "mas/logger"
  "sync"
)


var s_Logger logger.Logger = logger.NewLogger(logger.DEBUG)


type WorkerPool struct {
  m_NbWorker int
  m_WaitGroup *sync.WaitGroup
  m_ChannelIn chan Job
}


func NewWorkerPool(p_NbWorker int) *WorkerPool {
  workerPool := WorkerPool{}
  workerPool.m_NbWorker = p_NbWorker
  workerPool.m_WaitGroup = new(sync.WaitGroup)

  for i := 0; i < workerPool.m_NbWorker; i++ {
    workerPool.m_WaitGroup.Add(1)
    go Worker(i, workerPool.m_WaitGroup, workerPool.m_ChannelIn)
  }

  return &workerPool
}


func (w *WorkerPool) Do(p_Job Job) {
  w.m_ChannelIn <- p_Job
}


func (w *WorkerPool) Wait() {
  close(w.m_ChannelIn)
  w.m_WaitGroup.Wait()
  s_Logger.Debug("All workers end gracefully.")
}


type Job interface {
  Do()
}


func Worker(p_Id int, p_WaitGroup *sync.WaitGroup, p_ChannelIn chan Job) {
  defer p_WaitGroup.Done()
  s_Logger.Debug("Worker", p_Id, "started")

  for job := range p_ChannelIn {
    s_Logger.Debug(job)
  }
}
