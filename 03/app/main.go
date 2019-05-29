package main

import (
	"log"
	"net/http"      // https://golang.org/pkg/net/http/
	"text/template" // https://golang.org/pkg/text/template/

	"github.com/gorilla/mux" //https://github.com/gorilla/mux
	// ルーティングのサードパーティーパッケージです。
	// ルーティングは、net/httpという標準パッケージでも可能ですが、
	// gorilla/muxは、色々な機能が強化されています。
)

func main() {
	// ルーティング（URLのパスが/の場合には、関数greetが実行されます。)
	r := mux.NewRouter()
	r.HandleFunc("/", greet)

	// 8080番ポートでWEBアプリケーションを起動します。
	log.Fatal(http.ListenAndServe(":8080", r))
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
	data := Data{Name: "Gopher"}

	// index.htmlのTemplateにdataを渡して、生成されたhtmlをレスポンスとして返します。
	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}

