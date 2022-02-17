[toc]

# break

break语句可以结束for、switch和select的代码块。break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和select的代码块上。

```go
01   package  main
02
03   import  "fmt"
04
05   func  main()  {
06
07   	OuterLoop:
08         for  i  :=  0;  i  <  2;  i++  {
09               for  j  :=  0;  j  <  5;  j++  {
10                     switch  j  {
11                     case  2:
12                           fmt.Println(i,  j)
13                           break  OuterLoop
14                     case  3:
15                           fmt.Println(i,  j)
16                           break  OuterLoop
17                     }
18               }
19         }
20   }
```

