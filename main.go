package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main()  {
	var root, port string
	root = "./"
	port = "80"
	f,err := os.Open("../web.conf")
	if err != nil{
		fmt.Println("配置文件未找到")
		fmt.Println("默认根目录:",root)
		fmt.Println("默认监听端口:",port)
	}else{
		con,_ := ioutil.ReadAll(f)
		fmt.Println(string(con))
		root = jsoniter.Get(con,"root").ToString()
		fmt.Println("根目录:",root)
		port = jsoniter.Get(con,"port").ToString()
		fmt.Println("监听端口:",port)
	}

	router := httprouter.New()
	router.ServeFiles("/*filepath",http.Dir(root))

	log.Fatal(http.ListenAndServe(":"+port,router))



}
