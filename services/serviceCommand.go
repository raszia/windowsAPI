package services

import (
	"windows/command"
)

const (
	serviceActionStop    = "stop"
	serviceActionSart    = "start"
	serviceActionRestart = "restart"
)

func (req *ReqStruct) execute() error {
	return command.CmdCommand(command.CMDnet).ArgAdd(req.ServiceName).ArgAdd(req.ServiceAction).Run()
}
