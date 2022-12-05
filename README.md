## PublicClipboard-Server

使用 gin 实现剪切板服务后端

### 相关仓库
[PublicClipboard-Client](https://github.com/sxz799/PublicClipboard-Client)

[PublicClipboard-Web](https://github.com/sxz799/PublicClipboard-Web)


### 编译命令

#### mac/linux
```
go build -ldflags="-s -w" -o ClipboardServer main.go
```
#### windows

```
go build -ldflags="-s -w" -o ClipboardServer.exe main.go
```

### mac下交叉编译

#### 编译Linux可执行文件

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ClipboardServer main.go
```
#### 编译windows可执行文件

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ClipboardServer.exe main.go
```


