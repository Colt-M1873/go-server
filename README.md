# Go simple web server

![Mega.nz](https://img.shields.io/badge/Made%20with-Go-1f425f.svg?style=flat&logo=Go&logoColor=white&color=3ba9d5) ![HTML5](https://img.shields.io/badge/html5-%23E34F26.svg?style=flat&logo=html5&logoColor=white)

用go标准库 `net/http` 实现一个简陋（40行）的webserver

无额外依赖

File tree

```tree /F
...\GOWORKSPACE\SRC\GO-SERVER
│  go-server.exe 
│  go.mod 
│  main.go 
│  README.md
│  
└─static
        form.html
        index.html
```

url路径与对应功能如下图

```asciiflow
       ┌─►  "/"   ───────► index.html
       │
server─┼─► "/hello" ─────► hello func
       │
       └─► "/form" ─► form func ─► form.html
```

## `/`

因为一切的基础`/`被handle为了一个路径在`"./static"`的`http.FileServer`，在进入 `/`路径的时候默认提供`index.html`页面

## `/hello`

由`helloHandler`显示一个httpresponse，内容为

> "hello from go mini server!"

## `/form`

直接进入`/form`路径会发现页面上只显示了http.ResposeWriter返回的文字内容，而没有显示html页面

## `/form.html`

因为页面显示在在`form.html`路径下，`form.html`提供的是html页面，`form`对应的是formHandler函数

## 工作过程

如`form.html`中所写

```html
<form method="POST" action="/form">
```

当你在`form.html`页面中输入name和address，并点击submit按键之后，html页面的`<form>` 元素会调用POST方法，把数据给到`/form`这个路径

而在main.go中

```golang
http.HandleFunc("/form", formHandler)
```

`/form`这个路径由`formHandler`接管，`formHandler`将请求中的name和address参数通过`http.ResponseWriter`显示到页面里

## 关于logging

golang 标准库`net/http`里自带的FileServer，即`fs.go`，一共900多行，其中并没有默认的logging来显示客户端发来的请求和时间

用惯了框架，自带的FileServer默认没有提供logging功能，很不习惯，在Github上找到了一个专门的HTTP middleware（也就是handlers）的项目，包括logging和其他的一些功能，应该很实用 ，项目地址： [gorilla/handlers](https://github.com/gorilla/handlers)

以下的代码是一个手写的简单`loggingHandler`，就是用默认的log加上request的方法和内容来进行输出，只要将这个`loggingHandler`套在`FileServer`的外层即可，代码片段来自谷歌的[golang-nuts小组](https://groups.google.com/g/golang-nuts)里2013年的一条[讨论:Logging with builtin http FileServer](https://groups.google.com/g/golang-nuts/c/D6yevo6VyyM?pli=1)

```go
package main

import (
	"log"
	"net/http"
)

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":8080", loggingHandler(http.FileServer(http.Dir("."))))
}

```

## 参考

[Learn Go Programming by Building 11 Projects: the first Project](https://www.youtube.com/watch?v=jFfo23yIWac&t=2800s&ab_channel=freeCodeCamp.org)

[Go Standard Library `net/http`](https://pkg.go.dev/net/http#section-documentation)

ASCII flowchart: [asciiflow](https://asciiflow.com/)

Golang HTTP middleware: [gorilla/handlers](https://github.com/gorilla/handlers)
