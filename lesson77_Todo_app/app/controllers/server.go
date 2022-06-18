package controllers

import (
	"fmt"
	"net/http"

	"regexp"
	"strconv"
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

//URLの正規表現のパターンをコンパイル
var validPath = regexp.MustCompile("^/todos/(edit|update|delete)/([0-9]+$)")

func parseURL(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc { //HandlerFuncはHandler関数を返す関数
	return func(w http.ResponseWriter, r *http.Request) {
		// /todos/edit/1
		//このURLからIDを受け取って処理

		//validPathとURLがマッチした部分をスライスで取得したい
		q := validPath.FindStringSubmatch(r.URL.Path)

		fmt.Println(q)

		if q == nil {
			http.NotFound(w, r)
			return
		}
		//最後のパスをint型として受け取りたい
		qi, err := strconv.Atoi(q[2])
		if err != nil {
			http.NotFound(w, r)
		}

		fn(w, r, qi) //引数で渡したレスポンスライター、リクエスト、IDを渡す
	}
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
	http.HandleFunc("/todos/edit/", parseURL(todoEdit))
	http.HandleFunc("/todos/update/", parseURL(todoUpdate))
	http.HandleFunc("/todos/delete/", parseURL(todoDelete))

	return http.ListenAndServe(":"+config.Config.Port, nil)
	/*
		config.iniで設定したポート番号の値を読み込む
		第2引数は通常、デフォルトのマルチプレクサを使う
		デフォルトのマルチプレクサでは登録されていないURLにアクセスしたら
		404 page not found を返す
	*/

}
