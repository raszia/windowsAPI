//go:build windows
// +build windows

package services

import (
	"windows/command"
)

func (req *ReqStruct) execute() error {
	return command.CmdCommand(command.CMDnet).ArgAdd(req.ServiceName).ArgAdd(req.ServiceAction).Run()
}
