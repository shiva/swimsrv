package main

import (
    "code.google.com/p/gorest"
    "net/http"
)

func main() {
    gorest.RegisterService(new(HelloService)) //Register our service
    http.Handle("/", gorest.Handle())
    http.ListenAndServe(":8000", nil)
}

//Service Definition
type HelloService struct {
    gorest.RestService `root:"/tutorial/"`
    helloWorld         gorest.EndPoint `method:"GET" path:"/hello-world/" output:"string"`
    sayHello           gorest.EndPoint `method:"GET" path:"/hello/{name:string}" output:"string"`
    getJSON            gorest.EndPoint `method:"GET" path:"/get-json/" output:"ItemStore"
}

func (serv HelloService) HelloWorld() string {
    return "Hello World"
}

func (serv HelloService) SayHello(name string) string {
    return "Hello " + name
}

type ItemStore struct {
    Items []Item
}

func (serv HelloService) GetJSON() ItemStore {
    return
}
