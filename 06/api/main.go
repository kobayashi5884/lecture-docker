package main

import (
	"log"
	"net/http"
	"time"

	"github.com/auth0/go-jwt-middleware" // 認証に必要なJWTを利用するためのパッケージです。
	"github.com/dgrijalva/jwt-go"        // 認証に必要なJWTを利用するためのパッケージです。
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/cors"
	"github.com/urfave/negroni" // ミドルウェアを利用するためのパッケージです。
)

var db *gorm.DB

func main() {
	d, err := gorm.Open("mysql", "root:root@tcp(db)/my_db?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	db = d
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Comment{})

	rRoot := mux.NewRouter()                                 // デフォルトのルート
	rAuth := mux.NewRouter().PathPrefix("/auth").Subrouter() // 認証が必要なサブルート

	rRoot.HandleFunc("/user", postUser).Methods("POST") // 新規登録用のルート
	rRoot.HandleFunc("/user", putUser).Methods("PUT")   // ログイン用のルート

	rAuth.HandleFunc("/user", getUser).Methods("GET")

	rAuth.HandleFunc("/room", postRoom).Methods("POST")
	rAuth.HandleFunc("/room", getManyRooms).Methods("GET")
	rAuth.HandleFunc("/room/{id:[0-9]+}", getRoom).Methods("GET")
	rAuth.HandleFunc("/room/{id:[0-9]+}", deleteRoom).Methods("DELETE")

	rAuth.HandleFunc("/room/{id:[0-9]+}/comment", postComment).Methods("POST")
	rAuth.HandleFunc("/comment/{id:[0-9]+}", deleteComment).Methods("DELETE")

	// サブルーティングとミドルウェアの設定です。
	rRoot.PathPrefix("/auth").Handler(middlewareAuth(rAuth))
	handler := middlewareRoot(rRoot)

	s := &http.Server{
		Addr:           ":5000",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

// 全てのルートに適用されるミドルウェアです。
func middlewareRoot(rRoot *mux.Router) *negroni.Negroni {
	mwCors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	n := negroni.Classic() // アクセスログの出力等、negroniが用意している基本的なミドルウェアです。
	n.Use(mwCors)
	n.UseHandler(rRoot)
	return n
}

// 認証が必要なルート/auth~で適用されるミドルウェアです。
func middlewareAuth(rAuth *mux.Router) *negroni.Negroni {
	mwJwt := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("YourSecretKey"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return negroni.New(
		negroni.HandlerFunc(mwJwt.HandlerWithNext),
		negroni.Wrap(rAuth),
	)
}
