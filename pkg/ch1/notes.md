###Go notes


Go verbs:
```
%d
%x, %o, %b
%f, %g, %e
%t
%c
%s
%q
%v
%T
%%
```


Go switch:
go switch 的 case分支默认不会击穿, 如果想击穿, 使用 fallthrough 语句.
```
switch coinflip(){
case "heads":
	heads++
case "tails":
	tails++
default:
	fmt.Println("landed on edge!")
}
```
```
switch {
    case x > 0:
		return +1
	case x < 0:
		return -1
	default:
		return 0
}
```
```
switch simpleStatement; condition {
    case ...
}
```

命名类型:
```
type Point struct{
	X, Y int
}
var p Point
```
