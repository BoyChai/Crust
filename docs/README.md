# 快速开始

## 要求

- go 1.18以上

## 安装

```
$ go get -u github.com/BoyChai/Crust
```

## 使用

首先，创建一个名为`example.go`：

```
$ touch example.go
```

接下来，将以下代码放入`example.go`：

```
package main

import "github.com/BoyChai/Crust"

func main() {
	var s Crust.Shell
	s.RunShell()
}
```

而且，您可以通过以下方式运行代码`go run example.go`：

```
# 运行之后就创建好了一个Crust终端
$ go run .\example.go                                   
> 
> 
> 
```



#  基础

## 创建命令

通过下代码创建一个命令并导入到Crust里：

```
package main

import (
	"fmt"
	"github.com/BoyChai/Crust"
)

func main() {
	
	cmd1 := Crust.Command{
		ID:    1,
		Name:  "cmd1",
		Info:  "我是一条命令",
		Alias: "cmd",
	}
	
	Crust.NewCommand(cmd1, func() {
		fmt.Println("我是一条命令")
	})
	
	var s Crust.Shell
	s.RunShell()
}
```

运行结果如下：

```
# 进入Crust之后可以输入命令的ID来执行对应绑定的命令
$ go run .\example.go
> 
> 1
我是一条命令
>
```

## 输出帮助信息

默认Crust提供了输出信息的指令，输入help即可进行输出，代码如下：

```
package main

import (
	"fmt"
	"github.com/BoyChai/Crust"
)

func main() {

	cmd1 := Crust.Command{
		ID:    1,
		Name:  "cmd1",
		Info:  "我是一条命令",
		Alias: "cmd",
	}
	cmd2 := Crust.Command{
		ID:    2,
		Name:  "cmd2",
		Info:  "我是一条命令",
		Alias: "cmd2",
	}

	Crust.NewCommand(cmd1, func() {
		fmt.Println("我是一条命令")
	})
	Crust.NewCommand(cmd2, func() {
		fmt.Println("我是一条命令")
	})

	var s Crust.Shell
	s.RunShell()
}

```

运行时输入help打印帮助信息，如下：

```
$ go run .\example.go
> help
1) cmd1  2) cmd2
> 
```

## 前缀和分割

- Crust的组成

"前缀 分隔符 输入"

- 修改前缀

Crust的默认前缀为空，修改前缀可以通过下面代码：

```
package main

import (
	"github.com/BoyChai/Crust"
)

func main() {
	var s Crust.Shell
	s.SetShellTitle("Curst ")
	s.RunShell()
}
```

运行结果如下：

```
$ go run .\example.go
Curst >
Curst > 
Curst > 
```

- 修改分割符

修改分割符可以通过一下代码进行修改：

```
package main

import (
	"github.com/BoyChai/Crust"
)

func main() {
	var s Crust.Shell
	s.SetShellDivision("$ ")
	s.RunShell()
}
```

运行结果如下：

```
go run .\example.go
$ 
$ 
$ 
```

# 进阶

## 命令匹配策略

func (s *Shell) SetExplainType(explainType string)

命令匹配策略分id、name、any三种，id的匹配方式只匹配命令的id，name的匹配方式是名称和别名，any为id和name的结合都会进行匹配。默认的匹配策略是any，设置函数为

> func (s *Shell) SetExplainType(explainType string)

## 启动输出帮助信息

默认启动时是不输出帮助信息的，想要启动时输出可通过下面函数来实现：

> func (s *Shell) SetStartOutputHelp(startOutputHelp bool)

例如以下代码：

```
package main

import (
	"fmt"
	"github.com/BoyChai/Crust"
)

func main() {

	cmd1 := Crust.Command{
		ID:    1,
		Name:  "cmd1",
		Info:  "我是一条命令",
		Alias: "cmd",
	}
	cmd2 := Crust.Command{
		ID:    2,
		Name:  "cmd2",
		Info:  "我是一条命令",
		Alias: "cmd2",
	}

	Crust.NewCommand(cmd1, func() {
		fmt.Println("我是一条命令")
	})
	Crust.NewCommand(cmd2, func() {
		fmt.Println("我是一条命令")
	})

	var s Crust.Shell
	s.SetStartOutputHelp(true)
	s.RunShell()
}
```

程序启动效果如下：

```
$ go run .\example.go                
1) cmd1  2) cmd2
>
```

## 指定帮助信息输出格式

默认的帮助信息输出只会2个一行，可以通过以下函数来修改：

> func (s *Shell) SetHelpColumn(helpColumn int)

例如以下代码：

```
func main() {

	cmd1 := Crust.Command{
		ID:    1,
		Name:  "cmd1",
		Info:  "我是一条命令",
		Alias: "cmd",
	}
	cmd2 := Crust.Command{
		ID:    2,
		Name:  "cmd2",
		Info:  "我是一条命令",
		Alias: "cmd2",
	}

	Crust.NewCommand(cmd1, func() {
		fmt.Println("我是一条命令")
	})
	Crust.NewCommand(cmd2, func() {
		fmt.Println("我是一条命令")
	})

	var s Crust.Shell
	s.SetHelpColumn(1)
	s.RunShell()
}
```

输出信息如下：

```
$ go run .\example.go
> help
1) cmd1
2) cmd2
>
```



## 自定义帮助信息

帮助信息可以自己进行定义，默认是通过下面方法实现的

> func defaultHelpInfo(command []Command)

其中这个command是一个数组，里面存储着定义的全部命令，自己也可以定义一个输出信息。

我们可以通过一下方法来禁用默认的帮助信息模板

> func (s *Shell) SetDefaultHelpTemplate(defaultHelpTemplate bool)

当我们定义好了一个自己的帮助信息文档可以通过下面的方法来导入

> func (s *Shell) SetHelpTemplate(helpTemplate string)

例如下面的代码：

```
package main

import (
	"fmt"
	"github.com/BoyChai/Crust"
)

func main() {
	var s Crust.Shell
	s.SetDefaultHelpTemplate(false)
	s.SetHelpTemplate("我是自定help帮助信息")
	s.RunShell()
}
```

输出信息如下：

```
$ go run .\example.go
> help
我是自定义帮助信息
> 
```



# 其他

## 获取全部命令集

> func GetAllCommand() []Command

## 获取全部命令绑定函数

> func GetAllCMDFunction() []FunctionBinDing