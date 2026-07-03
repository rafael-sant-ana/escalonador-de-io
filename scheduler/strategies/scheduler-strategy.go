package strategies

type SchedulerStrategy interface {
	Name() string
	GetNext([]int, int) (int, []int, error)
}
