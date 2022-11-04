package configdata

type Flag string

type CmdData struct {
	name Flag
}

type AllCommandData interface {
	Flags() []Flag
}

type allCmdData struct {
	cmdData []CmdData
}

func (f allCmdData) Flags() (result []Flag) {
    for _, item := range f.cmdData {
        result = append(result, item.name)
    }

    return
}

func AvailableCommands() (result AllCommandData) {
	result = allCmdData{
		[]CmdData{
			{
				name: "--help",
			},
			{
				name: "--imgpath",
			},
		},
	}

	return
}
