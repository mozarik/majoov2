package main

import model "github.com/mozarik/majoov2/models"

func main() {
	db, err := model.InitDatabase()
	model.Drop(db)
	if err != nil {
		panic(err)
	}

	model.Migrate(db)

	user := model.User{}
	err = user.Create("zein", "123456", "merchant", db)
	if err != nil {
		panic(err)
	}
}
