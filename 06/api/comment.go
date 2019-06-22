package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Comment struct {
	ID        uint
	CreatedAt *time.Time
	Body      string
	User      *User // CommentがUserに帰属するというアソシエーションを定義しています。
	UserID    uint  // CommentがUserに帰属するというアソシエーションを定義しています。
	RoomID    uint  // 複数のCommentがRoomに帰属するというアソシエーションを定義しています。
}

func postComment(w http.ResponseWriter, r *http.Request) {
	userID, err := readUserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	m := map[string]string{"body": ""}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	comment := Comment{
		Body:   m["body"],
		UserID: userID,
		RoomID: uint(id),
	}
	if err := db.Create(&comment).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	user := User{}
	if err := db.Select("name").First(&user, userID).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	comment.User = &user

	resp, err := json.Marshal(&comment)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	userID, err := readUserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := db.Where("id = ? AND user_id = ?", id, userID).Delete(Comment{}).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
