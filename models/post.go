package models

// json package can only serialze exported fields of the struct
// all fields must start with a capital letter or they will not be serialized/marshaled
type Post struct {
	Id      int
	Title   string
	Content string
}

// Returns all posts
func AllPosts() ([]*Post, error) {

	rows, err := db.Query("SELECT * FROM post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pts []*Post

	for rows.Next() {
		pt := new(Post)
		err := rows.Scan(&pt.Id, &pt.Title, &pt.Content)
		if err != nil {
			return nil, err
		}

		pts = append(pts, pt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return pts, err
}
