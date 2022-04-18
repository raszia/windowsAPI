package command

type cmdCommandStruct struct {
	command string
	args    []string
}

func CmdCommand(theCommand string) *cmdCommandStruct {

	return &cmdCommandStruct{
		command: theCommand,
	}

}

func (cmd *cmdCommandStruct) ArgAdd(arg string) *cmdCommandStruct {
	cmd.args = append(cmd.args, arg)
	return cmd
}
