package task

const (
	JOB_EVENT_SAVE = 1
	JOB_EVENT_DELET = 2
)

// 变化事件： SAVE, DELETE
type JobEvent struct {
	EventType int
	Job *Job
}

func NewJobEvent(eventType int, job *Job) *JobEvent{
	e := &JobEvent{
		EventType: eventType,
		Job:       job,
	}
	return e
}
