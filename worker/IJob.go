package worker

// IJob interface that any job must implement to be executed by a worker.
type IJob interface {
  Do()
}
