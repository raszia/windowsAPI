//go:build linux
// +build linux

package services

import (
	"windows/command"
)

const CMDServiceLinux = "service"

func (req *ReqStruct) execute() error {
	return command.CmdCommand(CMDServiceLinux).ArgAdd(req.ServiceName).ArgAdd(req.ServiceAction).Run()
}
