package Crust

import (
	"errors"
	"strconv"
)

// CMD 初始化一个命令库
var CMD []Command

// CMDFunction 初始化一个方法库
var CMDFunction []FunctionBinDing

// NewCommand 创建一个命令
func NewCommand(command Command, Function func()) error {
	//ID判断不能为空
	if strconv.Itoa(command.ID) == "" {
		return errors.New("terminal id cannot be empty")
	}

	//创建命令绑定变量
	var commandFunc FunctionBinDing
	commandFunc.CommandID = command.ID
	commandFunc.Function = Function

	//将当前命令添加到命令库
	CMD = append(CMD, command)

	//将当前命令绑定变量添加到绑定库
	CMDFunction = append(CMDFunction, commandFunc)

	return nil
}

// GetAllCommand 获取命令库
func GetAllCommand() []Command {
	return CMD
}

// GetAllCMDFunction 获取命令绑定库
func GetAllCMDFunction() []FunctionBinDing {
	return CMDFunction
}
