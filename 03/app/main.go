package main

import (
	"log"
	"net/http"      // https://golang.org/pkg/net/http/
	"text/template" // https://golang.org/pkg/text/template/
)

func main() {
	// ルーティング（URLのパスが/の場合には、関数greetが実行されます。)
	http.HandleFunc("/", greet)
	// 8080番ポートでWEBアプリケーションを起動します。
	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	// index.htmlファイルをgoで扱えるデータ形式（*Template型）に変換します。
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// index.htmlのTemplateに出力すべきデータを定義します。
	// まずは、Data型を定義します。
	type Data struct {
		Name string
	}
	// そして、Data型の変数dataを定義します。
	data := Data{Name: "Docker"}

	// index.htmlのTemplateにdataを出力し、最終的なhtmlをレスポンスとして返します。
	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

