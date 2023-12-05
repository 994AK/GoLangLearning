# 入门开始

## hello
### 怎么创建一个`hello.go`，并且运行起来？

```text
1.创建初始化go模块
> go mod init example/hello
   
2.编写一个go程序,创建一个hello.go
package main

import "fmt"

func main() {
	fmt.Println("Hello")
}

3.引用外部包,以quit为例

package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Hello")
	fmt.Println(quote.Go())
}

执行:

> go mod tidy 
安装依赖

> go run .
执行

```

### 创建模块
```text
创建一个 greetings 模块,返回一个函数

> mkdir greetings
> go mod init example/greetings

创建文件 greetings.go

package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. ", name)

	return message
}

在使用的模块中
go mod edit -replace 包名 指向位置

我greetings在上一级
<home>
- greetings
- hello
  - hello.go
  - go.mod
> go mod edit -replace example/
greetings=../greetings
> go mod tidy

我在hello.go使用的,所以需要引入包
hello.go

package main

import (
	"example/greetings"
	"fmt"
)

import "rsc.io/quote"

func main() {

	message := greetings.Hello("994AK")
	fmt.Println("Hello")
	fmt.Println(quote.Go())
	fmt.Println(message)

}
```