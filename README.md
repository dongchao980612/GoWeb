# Go Web 编程快速入门

作者：软件工艺师(B站UP主)
代码：dongchao  
发布时间：2024-12-4
更新时间：2021-12-4

## 1、课程准备和一个demo

```go
package main

import "net/http"

func main() {
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
 })
 http.ListenAndServe("localhost:8080", nil) // defalitServeMux
 // 网络地址为localhost:8080
 // defalitServeMux为nil，即使用默认的路由器
}
```

## 2、Handle请求

```go
// 方法一：
http.ListenAndServe("localhost:8080", nil) // defalitServeMux

// 方法二：
server := http.Server{
    Addr:    "localhost:8080",
    Handler: nil, // defalitServeMux
}
server.ListenAndServe()
```

示例：

```go
package main

import (
 "net/http"
)

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("Hello page!"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("About page!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
 w.Write([]byte("Welcome page!"))
}

func main() {

 hh := helloHandler{}
 ah := aboutHandler{}
 server := http.Server{
  Addr:    "localhost:8080",
  Handler: nil, // DefaultServeMux
 }
 http.Handle("/hello", &hh) // 注册路由, 路径为/
 http.Handle("/about", &ah) // 注册路由, 路径为/about

 http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Home page!"))
 })

 // http.HandleFunc("/welcome", welcome) // mux.handle(pattern, HandlerFunc(handler))
 http.Handle("/welcome", http.HandlerFunc(welcome))

 server.ListenAndServe()

}
```

## 3、内置handler

> 访问<http://localhost:8080/index.html，将会返回www/index.html的内容>  

```go
// 方法一：
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    http.ServeFile(w, r, filepath.Join("www", r.URL.Path))
})

http.ListenAndServe(":8080", nil)

// 方法一：
http.ListenAndServe(":8080", http.FileServer(http.Dir("www")))
 
```

## 4、Request对象

```go
http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintln(w, r.URL.Path)+
    fmt.Fprintln(w, r.Header)
    fmt.Fprintln(w, r.Header["Accept-Encoding"])
    fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
})
http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
    length := r.ContentLength
    body := make([]byte, length)
    r.Body.Read(body)
    fmt.Fprintln(w, string(body))
})
http.ListenAndServe(":8080", nil)
```

## 查询参数

```go
http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
    url := r.URL
    query := url.Query()

    id := query.Get("id")
    name := query.Get("name")

    log.Printf("id: %s, name: %s", id, name)

})

http.ListenAndServe(":8080", nil)
```

## 5、Form表单

```go
fmt.Fprintln(w, "application/x-www-form-urlencoded")
r.ParseForm()
fmt.Fprintln(w, "Form", r.Form)
fmt.Fprintln(w, "PostForm", r.PostForm)
fmt.Fprintln(w, "multipart/form-data")
fmt.Fprintln(w, "MultipartForm", r.MultipartForm)
r.ParseMultipartForm(1024)
fmt.Fprintln(w, r.MultipartForm)

fmt.Fprintln(w, "application/x-www-form-urlencoded")
fmt.Fprintln(w, r.FormValue("username"))
fmt.Fprintln(w, r.PostFormValue("username"))

```
