package task

type JobEvent struct{
	EventType 	int // SAVE DELETE
	EventJob 	*Job
}

func NewJobEvent(eventType int, job *Job)(jobEvent *JobEvent) {
	return &JobEvent{
		EventType: eventType,
		EventJob: job,
	}
}
