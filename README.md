#### AC自动机

一个golang写的AC自动机敏感词匹配库

安装

```shell
go get -u -v github.com/lycblank/acmachine
```

使用

```golang
import (
    "github.com/lycblank/acmachine"
)
func main() {
    ac := acmachine.NewMachine(acmachine.SplitString, acmachine.CombineString)
    ac.AddPattern("彩票")
    ac.AddPattern("博彩")
    ac.AddPattern("广告")
    ac.Build()
    rets := ac.Match("我中了一个彩票")
    for _, ret := range rets {
        fmt.Println(ret.Pattern)
    }
}
```

运行结果
```shell
彩票
```