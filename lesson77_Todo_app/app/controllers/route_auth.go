package controllers

import (
	"lesson77_Todo_app/app/models"
	"log"
	"net/http"
)

//signupのハンドラを作成
func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { //GETメソッドの時は、テンプレートファイルの出力
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", 302)
		}

		//入力フォームにて値を入れて、登録を押すとPOSTメソッドで処理が行われる
	} else if r.Method == "POST" {
		//入力フォームで入力された値をもとに新しいユーザーを作成する
		err := r.ParseForm() //入力フォームの解析
		if err != nil {
			log.Println(err)
		}
		user := models.User{ //入力された値をUserのstructとして各フィールドで受け取る
			Name:     r.PostFormValue("name"), //signup.htmlで指定されている属性から値を取得
			Email:    r.PostFormValue("email"),
			PassWord: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil { //データベースにユーザーを登録
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302) //登録が成功したらtopページにリダイレクトする
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}

}

//ユーザーの認証
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302) //処理に失敗したらログインページにリダイレクト
	}
	if user.PassWord == models.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}

		//上で作成されたセッションをもとにして、cookieを作成してブラウザにログイン情報を保存する
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		http.Redirect(w, r, "/", 302)
	} else { //パスワードが一致しない場合
		http.Redirect(w, r, "/login", 302)

	}
	/*
		パスワードの判定
		パスワードは暗号化されているので、入力されたパスワードを暗号化し比較する
	*/

}

//ログアウトのハンドラ
func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value} //sessionのストラクトを作って、UUIDに取得したcookieのvalueを渡す
		session.DeleteSessionByUUID()
	}
	http.Redirect(w, r, "/login", 302)
}
