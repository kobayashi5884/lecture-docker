package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Room struct {
	ID       uint
	Title    string
	User     *User     // RoomがUserに帰属するというアソシエーションを定義しています。
	UserID   uint      // RoomがUserに帰属するというアソシエーションを定義しています。
	Comments []Comment // 複数のCommentがRoomに帰属するというアソシエーションを定義しています。
}

func postRoom(w http.ResponseWriter, r *http.Request) {
	userID, err := readUserID(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	m := map[string]string{"title": ""}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&m); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if strings.TrimSpace(m["title"]) == "" {
		m["title"] = "無題"
	}

	room := Room{
		Title:  m["title"],
		UserID: userID,
	}
	if err := db.Create(&room).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(&room)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func getManyRooms(w http.ResponseWriter, r *http.Request) {
	rooms := []Room{}
	if err := db.Find(&rooms).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(&rooms)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	// Preloadにより、RoomのCommentsにデータが読み込まれ、CommentsのUserにもデータが読み込まれます。
	room := Room{}
	switch err := db.
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		Preload("Comments.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, name")
		}).
		First(&room, id).
		Error; err {
	case nil:
	default:
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(&room)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

func deleteRoom(w http.ResponseWriter, r *http.Request) {
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

	if err := db.Where("id = ? AND user_id = ?", id, userID).Delete(Room{}).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := db.Where("room_id = ?", id).Delete(Comment{}).Error; err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
