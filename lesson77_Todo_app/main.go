package main

import (
	"fmt"
	"lesson77_Todo_app/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	//ユーザーを登録して、そのユーザーがToDoリストを作成するアプリ

	//init関数を呼び出したいのでこのようにやってるだけ、意味はない
	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@example.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	/*
		//取得処理
		u, _ := models.GetUser(1)
		fmt.Println(u)

		//更新処理
		u.Name = "Test2"
		u.Email = "test2@example.com"
		u.UpdateUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)

		//削除処理
		u.UserDeleteUser()
		u, _ = models.GetUser(1)
		fmt.Println(u)
	*/

	/*
		//新しく作成したユーザーはIDが2番に増分されているので2を指定する
		user, _ := models.GetUser(2)
		user.CreateTodo("First Todo")
	*/

	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/

	user, _ := models.GetUser(2)
	user.CreateTodo("Second Todo")

	todos, _ := models.GetTodos()
	for _, v := range todos {
		fmt.Println(v)
	}

}
