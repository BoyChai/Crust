package Crust

type Shell struct {
	//命令行前缀
	ShellTitle string
	//前缀和命令的分隔符 默认为"> "
	ShellDivision string
	//命令查找方式
	ExplainType string
	//是否使用默认的help模板
	DefaultHelpTemplate bool
	//help模板指定
	HelpTemplate func()
	//是否开启时输出help信息
	StartOutputHelp bool
	//help输出的列数
	HelpColumn int
}

type Command struct {
	ID    int    //命令ID
	Name  string //命令名称
	Info  string //命令详细信息
	Alias string //命令别名
}

type FunctionBinDing struct {
	//命令ID
	CommandID int
	//命令执行函数
	Function func()
}
