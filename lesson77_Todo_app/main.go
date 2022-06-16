package main

import (
	"fmt"

	"lesson77_Todo_app/app/controllers"
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
		u.Name = "test2"
		u.Email = "test2@example.com"
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

	/*
		user, _ := models.GetUser(3)
		user.CreateTodo("Third Todo")
	*/

	/*
		todos, _ := models.GetTodos()
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		user2, _ := models.GetUser(3)
		todos, _ := user2.GetTodosByUser()
		//取得したtodoをfor文で回す
		for _, v := range todos {
			fmt.Println(v)
		}
	*/

	/*
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	/*
		t, _ := models.GetTodo(3)
		t.DeleteTodo()
	*/

	//サーバーの立ち上げ
	controllers.StartMainServer()

	/*
		//入力されたEメールアドレスからDBのユーザーを取得
		user, _ := models.GetUserByEmail("test@example.com")
		fmt.Println(user)

		//セッションの作成
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(session)

		//セッションの確認
		valid, _ := session.CheckSession()
		fmt.Println(valid)
	*/

}
