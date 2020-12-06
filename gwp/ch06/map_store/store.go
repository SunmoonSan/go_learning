package main

import "fmt"

// Post -
type Post struct {
	Id      int
	Content string
	Author  string
}

// PostById -
var PostById map[int]*Post

// PostsByAuthor -
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello world1", Author: "San1"}
	post2 := Post{Id: 2, Content: "Hello world2", Author: "San1"}
	post3 := Post{Id: 3, Content: "Hello world3", Author: "San3"}
	post4 := Post{Id: 4, Content: "Hello world4", Author: "San4"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["San3"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthor["San1"] {
		fmt.Println(post)
	}

}
