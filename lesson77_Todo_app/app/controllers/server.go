package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"lesson77_Todo_app/app/models"
	"lesson77_Todo_app/config"
)

//ハンドラ関数でテンプレートを渡して表示
func generateHTML(w http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames { //引数でとったfilenamesをfileに入れる
		files = append(files, fmt.Sprintf("app/views/templates/%s.html", file))
	}
	/*
		filenamesのファイルを取り出し、Sprintfで呼び出したファイルパスに入れて、
		filesに格納する
	*/
	//defineで宣言したファイルを読み込む場合はExecuteTemplateで明示的に宣言する必要あり
	templates := template.Must(template.ParseFiles(files...)) //エラーが起きたらパニック状態になる
	templates.ExecuteTemplate(w, "layout", data)
}

func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie") //リクエストからクッキーを取得、関数authenticateでクッキーのネームとして渡した"_cookie"を指定する
	if err == nil {
		sess = models.Session{UUID: cookie.Value} //sessionのストラクトを作る、UUIDをcookieで取得したValueとする
		if ok, _ := sess.CheckSession(); !ok {    //sessionをチェック
			err = fmt.Errorf("Invalid session")
		}
	}
	return sess, err
}

func StartMainServer() error {
	//css、jsファイルを読み込む
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))
	/*
		staticの改装の下にcss、jsファイルが必要になってくるがstaticフォルダがないので、
		http.StripPrefix("/static/", files)でstaticを取り除いてfilesを指定
	*/

	//URLを登録してhtmlファイルを返して表示する
	http.HandleFunc("/", top) // 第2引数にハンドラー
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/authenticate", authenticate)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/todos", index) //ログインしているユーザーしか/todosにアクセスできない
	http.HandleFunc("/todos/new", todoNew)
	http.HandleFunc("/todos/save", todoSave)

	return http.ListenAndServe(":"+config.Config.Port, nil)
	/*
		config.iniで設定したポート番号の値を読み込む
		第2引数は通常、デフォルトのマルチプレクサを使う
		デフォルトのマルチプレクサでは登録されていないURLにアクセスしたら
		404 page not found を返す
	*/

}
