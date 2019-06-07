package main

import (
	"log"
	"net/http"

	"github.com/rs/cors" // https://github.com/rs/cors
	// 上記は、CORSの設定をするためのパッケージです。
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm" // https://github.com/jinzhu/gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// Gormは、Goでよく使われているORMです。
)

// 下記のようにグローバル変数を定義することができます。
var db *gorm.DB

func main() {
	// データベースに接続します。
	d, err := gorm.Open("mysql", "root:root@tcp(db)/my_db?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	db = d
	// deferで関数が終了・中断した場合に実行すべきコマンドを先に指定することができます。
	defer db.Close()

	// todo.goで定義しているTodoモデルのマイグレーション（テーブル作成）です。
	db.AutoMigrate(&Todo{})

	// REST APIに従って、ルーティングします。
	r := mux.NewRouter()
	r.HandleFunc("/todo", postTodo).Methods("POST")
	r.HandleFunc("/todo", getManyTodos).Methods("GET")
	r.HandleFunc("/todo/{id:[0-9]+}", putTodo).Methods("PUT")
	r.HandleFunc("/todo/{id:[0-9]+}", deleteTodo).Methods("DELETE")

	// CORSの設定をします。
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	log.Fatal(http.ListenAndServe(":5000", c.Handler(r)))
}
