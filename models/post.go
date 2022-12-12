// Post model
package models

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func GetPost() Post {
	var post Post
	return post
}

func GetPosts() []Post {
	var posts []Post
	return posts
}
