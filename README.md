# go-instrument

通过ast包提供的能力对go源码进行增强，需要把编译出的可执行文件放在目标项目目录下执行（本项目主要用于学习go语言的ast相关技术总结，有功能需要可以参考其中的实现思路自行开发）

## 本地编译
```bash
go build -o ./bin/go-instrument ./cmd/
```