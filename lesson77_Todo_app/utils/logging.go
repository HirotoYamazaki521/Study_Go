package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) //読み書き|ファイルの作成|ファイルの追記
	if err != nil {
		log.Fatalln(err)
	}

	multiLogFile := io.MultiWriter(os.Stdout, logfile)   //ログの書き込み先を標準出力とlogfileに指定
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) //フォーマットを指定
	log.SetOutput(multiLogFile)                          //ログの出力先をマルチログファイルとして指定
}
