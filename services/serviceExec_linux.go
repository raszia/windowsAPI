//go:build linux
// +build linux

package services

import (
	"windows/command"
)

func (req *ReqStruct) execute() error {
	return command.CmdCommand(command.CMDServiceLinux).ArgAdd(req.ServiceName).ArgAdd(req.ServiceAction).Run()
}
