package controllers

import (
	"log"
	"net/http"
)

//ハンドラー
/*
以下の様な引数をとる関数はハンドラーとして登録できる
*/

func top(w http.ResponseWriter, r *http.Request) {
	/*
		t, err := template.ParseFiles("app/views/templates/top.html")
		if err != nil {
			log.Fatalln(err)
		}

			//htmlのテンプレートパッケージ
			//引数で渡したテンプレートファイルを解析し、
			//テンプレートの構造体を生成する
			//top.htmlのファイル名をテンプレート名とする新しいテンプレートを生成

		t.Execute(w, "hello")

			//tのテンプレートファイルを実行
			//第1引数はResponseWriter、第2引数はデータを渡す
	*/
	_, err := session(w, r) //cookieを取得
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top") //セッションが存在しない場合topにアクセス
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r) //クッキーを取得し、セッションと一致するものがあればログインしている
	if err != nil {
		http.Redirect(w, r, "/", 302) //topページにリダイレクト
	} else {
		user, err := sess.GetUserBySession() //sessionのUserIDと一致するユーザーを取得
		if err != nil {
			log.Println(err)
		}
		//ユーザーが作成したToDo一覧を取得
		todos, _ := user.GetTodosByUser()
		user.Todos = todos
		//userの情報をテンプレートに渡したいので、第2引数のデータをuserに指定
		generateHTML(w, user, "layout", "private_navbar", "index") //ログインしている場合

	}

}

func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		err = r.ParseForm() //Formを取得
		if err != nil {
			log.Println(err)
		}
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}
}
