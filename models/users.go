package models

type User struct {
	id int
	email string
	username string
	passwordHash string
}

func getUser() (*User, error) {
	user, err := db.Query("SELECT * FROM users WHERE email='r.g.com'")
	if err != nil {
		return nil, err
	}
	defer user.Close()

	//bks := make([]*Book, 0)
	//for rows.Next() {
	//	bk := new(Book)
	//	err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	//	if err != nil {
	//		return nil, err
	//	}
	//	bks = append(bks, bk)
	//}
	//if err = rows.Err(); err != nil {
	//	return nil, err
	//}
	return nil, nil
}
