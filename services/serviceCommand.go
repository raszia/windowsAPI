package services

import (
	"windows/utility"
)

const (
	serviceCMD           = "net"
	serviceActionStop    = "stop"
	serviceActionSart    = "start"
	serviceActionRestart = "restart"
)

type serviceCommandStruct struct {
	command string
	action  string
	args    []string
}

func (cmd *serviceCommandStruct) serviceName(name string) *serviceCommandStruct {
	cmd.args = append(cmd.args, name)
	return cmd
}

func (cmd *serviceCommandStruct) serviceAction(name string) *serviceCommandStruct {
	cmd.args = append(cmd.args, name)
	cmd.action = name
	return cmd
}

func serviceCommand() *serviceCommandStruct {

	cmd := &serviceCommandStruct{}
	cmd.command = serviceCMD
	return cmd
}

func (cmd *serviceCommandStruct) Run() error {

	return utility.RunCmd(cmd.command, cmd.args...)

}

func (req *ReqStruct) execute() error {
	return serviceCommand().serviceName(req.ServiceName).serviceAction(req.ServiceAction).Run()

}
