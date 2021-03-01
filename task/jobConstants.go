package task

type Constans struct{
	JOB_SAVE_DIR 		string
	JOB_KILL_DIR 		string
	JOB_EVENT_SAVE 		int
	JOB_EVENT_DELETE 	int
}

var GC *Constans= &Constans{
	"/task/jobs/",
	"/task/kill/",
	1,
	2,
}

