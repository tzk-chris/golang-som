[toc]

# switch

## fallthrough



在Go语言中case是一个独立的代码块，执行完毕后不会像C语言那样紧接着下一个case执行。但是为了兼容一些移植代码，依然加入了fallthrough关键字来实现这一功能，代码如下：

```go
var s = "hello"
switch {
case s == "hello":
    fmt.Println("hello")
    fallthrough
case s != "world":
    fmt.Println("world")
}
```

