package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Todoモデルを定義します。
type Todo struct {
	ID   uint
	Text string
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからJSONを読み込み、マップにデコードします。
	m := map[string]string{"text": ""}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// レコードを作成します。作成結果は、todoに反映されます。
	todo := Todo{Text: m["text"]}
	if err := db.Create(&todo).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// todoをJSONにエンコードします。
	resp, err := json.Marshal(&todo)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// エンコードしたJSONをレスポンスに書き込みます。
	w.Write(resp)
}

func getManyTodos(w http.ResponseWriter, r *http.Request) {
	// 全てのレコードを取得します。取得結果は、todosに反映されます。
	todos := []Todo{}
	if err := db.Find(&todos).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// todosをJSONにエンコードします。
	resp, err := json.Marshal(&todos)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// エンコードしたJSONをレスポンスに書き込みます。
	w.Write(resp)
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	// URLからidを抽出し、文字列から数値に変換します。
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// リクエストボディからJSONを読み込み、マップにデコードします。
	m := map[string]string{"text": ""}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// idでレコードを取得します。取得結果は、todoに反映されます。
	todo := Todo{}
	if err := db.First(&todo, id).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// todoを上書きし、レコードを更新します。
	todo.Text = m["text"]
	if err := db.Save(&todo).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// todosをJSONにエンコードします。
	resp, err := json.Marshal(&todo)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// エンコードしたJSONをレスポンスに書き込みます。
	w.Write(resp)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	// URLからidを抽出し、文字列から数値に変換します。
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// idでレコードを削除します。
	if err := db.Where("id = ?", id).Delete(Todo{}).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
