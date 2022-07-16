package Crust

import (
	"fmt"
	"strconv"
)

// help 输出帮助信息
func (s *Shell) help() {
	if ShellConfig.DefaultHelpTemplate {
		defaultHelpInfo(CMD)
	} else {
		ShellConfig.HelpTemplate()
	}
}

//
func defaultHelpInfo(command []Command) {
	if len(command) == 1 {
		fmt.Printf(strconv.Itoa(command[0].ID) + ") " + command[0].Name + "\n")
		return
	}
	for i := 0; i < len(command); i++ {
		fmt.Printf(strconv.Itoa(command[i].ID) + ") " + command[i].Name)
		if i == len(command) {
			fmt.Println("")
			//判断输出几段之后换行↓
		} else if (i + 1/ShellConfig.HelpColumn) == 0 {
			fmt.Printf("  ")
		} else {
			fmt.Println("")
		}
	}
}
