package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"myGoWebserver/models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var dbconn *sqlx.DB

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var posts = models.GetPosts()

	sqlStmt := `SELECT * FROM posts`
	rows, err := dbconn.Queryx(sqlStmt)

	if err == nil {
		var tempPost = models.GetPost()

		for rows.Next() {
			err = rows.StructScan(&tempPost)
			posts = append(posts, tempPost)
		}

		switch err {
		case sql.ErrNoRows:
			{
				log.Println("No rows returned.")
				http.Error(w, err.Error(), 204)
			}
		case nil:
			json.NewEncoder(w).Encode(&posts)
		default:
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var searchpost = models.GetPost()

	sqlStmt := `SELECT * FROM posts WHERE id=$1`
	row := dbconn.QueryRowx(sqlStmt, id)
	switch err := row.StructScan(&searchpost); err {
	case sql.ErrNoRows:
		{
			log.Println("No rows returned.")
			http.Error(w, err.Error(), 204)
		}
	case nil:
		json.NewEncoder(w).Encode(&searchpost)
	default:
		http.Error(w, err.Error(), 400)
		return
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post = models.GetPost()
	var id int

	_ = json.NewDecoder(r.Body).Decode(&post)

	sqlStmt := `INSERT INTO posts(title, body) VALUES($1,$2) RETURNING id`
	err := dbconn.QueryRow(sqlStmt, post.Title, post.Body).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	post.ID = id
	log.Println("New record ID is:", id)
	json.NewEncoder(w).Encode(&post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var post = models.GetPost()
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID, _ = strconv.Atoi(params["id"])

	id := 0
	sqlStmt := `UPDATE posts SET title=$1, body=$2 WHERE id=$3 RETURNING id`
	err := dbconn.QueryRow(sqlStmt, post.Title, post.Body, params["id"]).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Println("Updated record ID is:", id)
	json.NewEncoder(w).Encode(&post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := 0
	sqlStmt := `DELETE FROM posts WHERE id=$1 RETURNING id`
	err := dbconn.QueryRow(sqlStmt, params["id"]).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	log.Println("Deleted record ID is:", id)
	json.NewEncoder(w).Encode(id)
}

func SetDB(db *sqlx.DB) {
	dbconn = db
}
