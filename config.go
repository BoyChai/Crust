package Crust

import (
	"errors"
	"fmt"
)

var ShellConfig Shell

func init() {
	ShellConfig.ShellDivision = "> "
	// 默认命令匹配方式为any
	ShellConfig.ExplainType = "any"
	// 是否使用默认模板
	ShellConfig.DefaultHelpTemplate = true
	// 默认help输出的列数
	ShellConfig.HelpColumn = 2
	// 默认启动时不输出help
	ShellConfig.StartOutputHelp = false
}

// SetShellTitle 修改命令行前缀
func (s *Shell) SetShellTitle(shellTitle string) {
	ShellConfig.ShellTitle = shellTitle
}

// SetShellDivision 修改命令行前缀
func (s *Shell) SetShellDivision(shellDivision string) {
	ShellConfig.ShellDivision = shellDivision
}

// SetExplainType 设置处理器的模式
func (s *Shell) SetExplainType(explainType string) error {
	if explainType == "id" || explainType == "name" || explainType == "any" {
		ShellConfig.ExplainType = explainType
		return nil
	}
	return errors.New("the value of explainType can only be 'id' or 'name' or 'any'")

}

// SetStartOutputHelp 设置是否启动时输出help信息
func (s *Shell) SetStartOutputHelp(startOutputHelp bool) {
	ShellConfig.StartOutputHelp = startOutputHelp
}

//SetHelpColumn 设置输出列数
func (s *Shell) SetHelpColumn(helpColumn int) {
	ShellConfig.HelpColumn = helpColumn

}

// SetDefaultHelpTemplate 设置是否使用默认help模板
func (s *Shell) SetDefaultHelpTemplate(defaultHelpTemplate bool) {
	ShellConfig.DefaultHelpTemplate = defaultHelpTemplate
}

// SetHelpTemplate 自定义help模板
func (s *Shell) SetHelpTemplate(helpTemplate string) {
	ShellConfig.HelpTemplate = func() {
		fmt.Println(helpTemplate)
	}
}
