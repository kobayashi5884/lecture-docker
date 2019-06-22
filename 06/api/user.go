package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID       uint
	Name     string
	Password string `json:"-"`
	// `json:"-"`は、タグといわれるもので、jsonパッケージに対して、
	// エンコードの際に、このフィールドを省略するように指示しています。
}

// リクエストヘッダーのJWTからユーザーIDを取得する関数です。
func readUserID(r *http.Request) (uint, error) {
	user, ok := r.Context().Value("user").(*jwt.Token)
	if !ok {
		return 0, errors.New("cannot read user ID")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("cannot read user ID")
	}
	userID, ok := claims["sub"].(float64)
	if !ok {
		return 0, errors.New("cannot read user ID")
	}
	if userID < 1 {
		return 0, errors.New("cannot read user ID")
	}
	return uint(userID), nil
}

func postUser(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"name":     "",
		"password": "",
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	m["name"] = strings.TrimSpace(m["name"])
	if m["name"] == "" {
		http.Error(w, "Name is empty", http.StatusInternalServerError)
		return
	}
	if !db.Where("name = ?", m["name"]).Take(&User{}).RecordNotFound() {
		http.Error(w, "この名前は既に登録されています", http.StatusInternalServerError)
		return
	}

	// パスワードの暗号化
	hash, err := bcrypt.GenerateFromPassword([]byte(m["password"]), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user := User{
		Name:     m["name"],
		Password: string(hash),
	}
	if err := db.Create(&user).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// JWTを作成し、レスポンスボディに書き込みます。
	// JWTのClaimsについては、下記を参照してください。
	// https://www.iana.org/assignments/jwt/jwt.xhtml#claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["iat"] = time.Now()

	signedToken, err := token.SignedString([]byte("YourSecretKey"))
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{"token": signedToken})
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{
		"name":     "",
		"password": "",
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user := User{}
	if err := db.Where("name = ?", m["name"]).Take(&user).Error; err != nil {
		log.Println(err)
		http.Error(w, "名前またはパスワードが間違っています", http.StatusInternalServerError)
		return
	}

	// パスワードのチェック
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(m["password"])); err != nil {
		log.Println(err)
		http.Error(w, "名前またはパスワードが間違っています", http.StatusInternalServerError)
		return
	}

	// JWTを作成し、レスポンスボディに書き込みます。
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["iat"] = time.Now()

	signedToken, err := token.SignedString([]byte("YourSecretKey"))
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]string{"token": signedToken})
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := readUserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user := User{}
	if err := db.First(&user, userID).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"ID":   user.ID,
		"Name": user.Name,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(response)
}
