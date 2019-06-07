package main

import (
	"log"
	"net/http"      // https://golang.org/pkg/net/http/

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
	// レスポンスボディに文字列を書き込みます。
	w.Write([]byte("Hello World"))
}

