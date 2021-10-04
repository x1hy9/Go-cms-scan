# Go-cms-scan
第一次用GO写的工具，指纹采集于GOBY


# 使用方法

```
Usage: 扫描器 [-f]|[-u]

欢迎使用扫描器

Options:
  -v, --version   Show the version and exit
  -f, --file      文件名称
  -u, --url       单个URL
  ```
  

# 编译
```
go build cmd/onefinger/main.go
```
 
# 用例
```
go run .\main.go -f 批量文件名

go run .\mian.go -u 单个url
```
