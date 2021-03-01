package service

import (
	"easycron/common"
	"easycron/master/model"
)

type JobService struct{
}

func (service *JobService) SaveJob(name string, command string, expr string) (resp *common.Response){

	job := &model.JobModel{
		Name: 		name,
		Command: 	command,
		Expr: 		expr,
	}
	
	oldJob, err := job.Save()

	resp = &common.Response{
		Code:		0,
		Msg: 		"",
		Err: 		err,
		Data: 		oldJob,
	}

	return
}

func (service *JobService) DeleteJob(name string)  (resp *common.Response) {

	job := &model.JobModel{
		Name: name,
		Command: "",
		Expr: "",
	}

	oldJob, err := job.Delete()

	resp = &common.Response{
		Code:		0,
		Msg: 		"",
		Err: 		err,
		Data: 		oldJob,
	}

	return
}

func (service *JobService) UpdateJob(name string, command string, expr string) (resp *common.Response){

	resp = service.SaveJob(name, command, expr)

	return
}

func (service *JobService) ListJobs(name string)  (resp *common.Response) {

	job := &model.JobModel{
		Name: name,
	}
	
	jobs, err := job.List()

	resp = &common.Response{
		Code:		0,
		Msg: 		"",
		Err: 		err,
		Data: 		jobs,
	}
	return
}

func (service *JobService) KillJob(name string)  (resp *common.Response) {

	job := &model.JobModel{
		Name: name,
		Command: "",
		Expr: "",
	}
	
	err := job.Kill()

	resp = &common.Response{
		Code:		0,
		Msg: 		"",
		Err: 		err,
		Data: 		nil,
	}

	return
}