package services

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

const (
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
	cmd.command = "dnscmd"
	return cmd
}

func (cmd *serviceCommandStruct) Run() error {
	switch cmd.action {
	case serviceActionStop:
	case serviceActionSart:
	case serviceActionRestart:
	default:
		return errors.New("bad service action: " + cmd.action)
	}
	fmt.Println(cmd.command, strings.Join(cmd.args, " "))
	x := exec.Command(cmd.command, cmd.args...)

	var buff, errBuf bytes.Buffer
	x.Stdout = &buff
	x.Stderr = &errBuf
	err := x.Run()
	return err

}
