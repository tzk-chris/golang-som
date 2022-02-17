[toc]

# goto

goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。



## 常规退出

```go
01   package main
02
03   import "fmt"
04
05   func main() {
06
07        var break Again bool
08
09       // 外循环
10        for x := 0; x < 10; x++ {
11
12            // 内循环
13              for y := 0; y < 10; y++ {
14
15                 // 满足某个条件时，退出循环
16                    if y == 2 {
17
18                     // 设置退出标记
19                         break Again = true
20
21                     // 退出本次循环
22                         break
23                    }
24
25              }
26
27            // 根据标记，还需要退出一次循环
28              if break Again {
29                    break
30              }
31
32        }
33
34        fmt.Println("done")
35   }
```



## goto进行优化

```go
01   package main
02
03   import "fmt"
04
05   func main() {
06
07        for x := 0; x < 10; x++ {
08
09              for y := 0; y < 10; y++ {
10
11                    if y == 2 {
12                     // 跳转到标签
13                         goto break Here
14                    }
15
16              }
17        }
18
19       // 手动返回，避免执行进入标签
20        return
21
22       // 标签
23   break Here:
24        fmt.Println("done")
25   }
```



## 同一错误处理

```go
err := first Check Error()
if err != nil {
  fmt.Println(err)
  exit Process()
  return
}

err := second Check Error()
if err != nil {
  fmt.PrintLn(err)
  exit Process()
  return
}

fmt.Println("done")
```

在上面代码中，加粗部分都是重复的错误处理代码。后期陆续在这些代码中如果添加更多的判断，就需要在每一块雷同代码中依次修改，极易造成疏忽和错误。

## goto优化

```go
err := first Check Error()
if err != nil {
  goto on Exit
}

err := second Check Error()
if err != nil {
  goto on Exit
}

fmt.Println("done")
return

on Exit:
	fmt.println(err)
	exit Process()
```

● 第3行和第9行，发生错误时，跳转错误标签on Exit。

● 第17行和第18行，汇总所有流程进行错误打印并退出进程。