package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/smebellis/lenslocked/models"
)

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connected ...")

	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("bob5@bob.com", "bob123")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
	// 	id SERIAL PRIMARY KEY,
	// 	name TEXT,
	// 	email TEXT NOT NULL
	// );

	// CREATE TABLE IF NOT EXISTS orders (
	// 	id SERIAL PRIMARY KEY,
	// 	user_id INT NOT NULL,
	// 	amount INT,
	// 	description TEXT
	// );`)

	// name := "Jon Calhoun"
	// email := "jon@calhoun.io"

	// row := db.QueryRow(`
	// 	INSERT INTO users(name, email)
	// 	VALUES($1, $2) RETURNING id;`, name, email)
	// var id int

	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("User Created. id = ", id)

	// userID := 1

	// for i := 1; i < 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("Fake order #%d", i)
	// 	_, err := db.Exec(`
	// 		INSERT INTO orders(user_id, amount, description)
	// 		VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("Created fake orders.")
	// }

	// type Order struct {
	// 	ID          int
	// 	UserId      int
	// 	Amount      int
	// 	Description string
	// }

	// var orders []Order

	// rows, err := db.Query(`
	// 	SELECT id, amount, description
	// 	FROM orders
	// 	WHERE user_id=$1`, userID)
	// if err != nil {
	// 	panic(err)
	// }

	// defer rows.Close()

	// for rows.Next() {
	// 	var order Order
	// 	order.UserId = userID
	// 	err := rows.Scan(&order.ID, &order.Amount, &order.Description)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	orders = append(orders, order)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Orders:", orders)

	// id := 200
	// row := db.QueryRow(`
	// 	SELECT name, email
	// 	FROM users
	// 	WHERE id=$1;`, id)
	// var name, email string
	// err = row.Scan(&name, &email)
	// if err == sql.ErrNoRows {
	// 	fmt.Println("Error, no rows!")
	// }
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("User information: name=%s, email=%s\n", name, email)

}
