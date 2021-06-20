package main

import model "github.com/mozarik/majoov2/models"

func main() {
	db, err := model.InitDatabase()
	model.Drop(db)
	if err != nil {
		panic(err)
	}

	model.Migrate(db)

	user := &model.User{
		Username: "username",
		Password: "password",
		Role:     "merchant",
	}
	err = user.Create(db, user)
	if err != nil {
		panic(err)
	}
}
