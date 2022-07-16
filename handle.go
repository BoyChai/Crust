package Crust

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// RunShell 启动一个shell

func (s *Shell) RunShell() {
	if ShellConfig.StartOutputHelp {
		s.help()
	}
	s.processor()
}

// processor 命令行处理器
func (s *Shell) processor() {
	s.print()
	action, err := s.receiver()
	err = s.explain(action)
	if err != nil {
		fmt.Println(err)
	}
	s.processor()

}

// 命令行输出器
func (s *Shell) print() {
	fmt.Print(ShellConfig.ShellTitle)
	fmt.Printf(ShellConfig.ShellDivision)
}

// 命令行接收器
func (s *Shell) receiver() (string, error) {
	var action string
	_, err := fmt.Scanln(&action)
	return action, err
}

// 命令行解释器
func (s *Shell) explain(action string) error {
	cmd := CMD
	explainType := ShellConfig.ExplainType
	// 解释退出指令
	if action == "" {
		return nil
	} else if action == "q" || action == "quit" || action == "exit" {
		os.Exit(0)
	} else if action == "help" {
		s.help()
	} else {
		// 解释器进行命令匹配
		switch explainType {
		case "id":
			for i := 0; i < len(cmd); i++ {
				if action == strconv.Itoa(cmd[i].ID) {
					s.implement(i)
					return nil
				}
			}

		case "name":
			for i := 0; i < cap(cmd); i++ {
				if action == cmd[i].Name {
					s.implement(i)
					return nil
				}
				if action == cmd[i].Alias {
					s.implement(i)
					return nil
				}
			}
		case "any":
			for i := 0; i < cap(cmd); i++ {
				if action == strconv.Itoa(cmd[i].ID) {
					s.implement(i)
					return nil
				}
				if action == cmd[i].Name {
					s.implement(i)
					return nil
				}
				if action == cmd[i].Alias {
					s.implement(i)
					return nil
				}
			}
		default:
			return errors.New("command not found")
		}
	}

	return nil
}

// 命令行执行器
func (s *Shell) implement(capId int) {
	//从命令绑定库查询指令之后执行
	for i := 0; i < cap(CMDFunction); i++ {
		if CMD[capId].ID == CMDFunction[i].CommandID {
			CMDFunction[i].Function()
		}
	}
}
