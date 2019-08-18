package task

import (
	"dcs-gocron/common"
	"strings"
)

type Job struct {
	Name string
	Command string
	CronExpr string
}
//
//func (job *Job)UnPack(data []byte) (err error) {
//	err = json.Unmarshal(data, &job)
//	return
//}


// 获取真实任务名: /cron/jobs/job20  -> job20
func (job *Job)ExtractRealName() string{
	return strings.TrimPrefix(job.Name, common.JOBDIR)
}
