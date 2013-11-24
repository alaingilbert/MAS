package worker

import (
  "mas/logger"
  "sync"
)

var sLogger = logger.NewLogger(logger.DEBUG)

// WorkerPool implementation.
type WorkerPool struct {
  mNbWorker  int
  mWaitGroup *sync.WaitGroup
  mChannelIn chan IJob
}

// NewWorkerPool will create an instance of a worker pool.
// pNbWorker number of worker to spawn.
// It returns a pointer to the pool.
func NewWorkerPool(pNbWorker int) *WorkerPool {
  workerPool := WorkerPool{}
  workerPool.mNbWorker = pNbWorker
  workerPool.mWaitGroup = new(sync.WaitGroup)
  workerPool.mChannelIn = make(chan IJob)

  for i := 0; i < workerPool.mNbWorker; i++ {
    workerPool.mWaitGroup.Add(1)
    go Worker(i, workerPool.mWaitGroup, workerPool.mChannelIn)
  }

  return &workerPool
}

// Do give a job to one of the worker.
// pJob the job to be executed.
func (w *WorkerPool) Do(pJob IJob) {
  w.mChannelIn <- pJob
}

// Wait will close and wait until all workers are done.
func (w *WorkerPool) Wait() {
  close(w.mChannelIn)
  w.mWaitGroup.Wait()
  sLogger.Debug("All workers end gracefully.")
}

// Worker itself.
// pID id of the worker.
// pWaitGroup pointer to the WaitGroup instance.
// pChannelIn the channel where to jobs come from.
func Worker(pID int, pWaitGroup *sync.WaitGroup, pChannelIn chan IJob) {
  defer pWaitGroup.Done()
  sLogger.Debug("Worker", pID, "started")

  for job := range pChannelIn {
    job.Do()
  }
}
