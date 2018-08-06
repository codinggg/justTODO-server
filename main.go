package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
)

type TodoData struct {
	Time string `json:"time"`
	Info string `json:"info"`
}

func main() {
	app := iris.New()

	app.Get("/get", handleGet)

	app.Post("/post", handlePost)

	app.Run(iris.Addr(":8808"), iris.WithCharset("UTF-8"), iris.WithoutVersionChecker)
}

func handlePost(ctx iris.Context) {
	var todoData TodoData
	err := ctx.ReadJSON(&todoData)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
		return
	}
	ctx.Writef("-Received: %#+v\n", todoData)

	rtn, errs := json.Marshal(todoData)
	if errs != nil {
		fmt.Println("err with marshal")
		return
	}
	str1 := string(rtn[:])
	fmt.Println(str1)
}

func handleGet(ctx iris.Context) {
	todoList := []TodoData{
		{"2018-05-06 12:00:00", "test1"},
		{"2018-05-06 13:00:00", "test2"},
		{"2018-05-06 14:00:00", "test3"},
		{"2018-05-06 15:00:00", "test4"},
	}
	ctx.JSON(todoList)
}
