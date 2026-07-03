package consumers

type SchedulerStrategy interface {
	Name() string
	getNext() (int, error)
}
