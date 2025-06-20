package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("HOST")
	port     = os.Getenv("PORT")
	user     = os.Getenv("USER")
	password = os.Getenv("PASS")
	dbname   = os.Getenv("DB_NAME")
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "farhan143143"
// 	dbname   = "goapi"
// )

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error while loading .env")
	}

	host := os.Getenv("HOST")
	portstr := os.Getenv("PORT")
	user := os.Getenv("user")
	password := os.Getenv("PASS")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portstr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connecting as:", user, "to DB:", dbname, "on", host, "with port", port)
	fmt.Println("Password is:", password)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	// SqlStat := `
	// INSERT INTO USERS (id,age,first_Name,last_Name,email)
	// VALUES ($1,$2,$3,$4,$5) RETURNING ID
	// `
	// id := 0 // This will get Overwritten when the actual ID is called
	// // var id int
	// // _, err = db.Exec(SqlStat, 4, 12, "haggu", "sahhu", "saggu.dev@gmail.com")
	// err = db.QueryRow(SqlStat, 8, 7822, "Laal", "Laddu", "Laal.dev@gmail.com").Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("New record ID is:", id)
	// SqlDelete := `
	// UPDATE USERS
	//  SET FIRST_NAME=$1 , LAST_NAME=$2
	// WHERE ID=$3
	// `
	// value, err := db.Exec(SqlDelete, "Monkey", "Luffy", 2)
	// fmt.Println("Value  : ", &value)
	// if err != nil {
	// 	panic(err)
	// }
	SQLQuery(db, 3)
	SQLMultipleQuery(db, 5)

}

// TODO : Create Different Functions for different operations
func SQLQuery(db *sql.DB, UserId int) {
	//NOTE : This is for Querying a Single Record but sometimes you would want to return entire record
	type User struct {
		id         int
		age        int
		first_Name string
		last_Name  string
		email      string
	}
	//FIXME : Query the entire Record
	SqlQuery := `
	SELECT ID , EMAIL FROM USERS WHERE ID=$1
	`
	var email string
	var id int
	row := db.QueryRow(SqlQuery, 3)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No Rows Were Found")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}
}
func SQLMultipleQuery(db *sql.DB, UserId int) {
	rows, err := db.Query(`SELECT ID,FIRST_NAME,LAST_NAME FROM USERS LIMIT $1`, UserId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	// Loops through the records and stores the values with the help of scan
	for rows.Next() {
		var id int
		var first_Name string
		var last_Name string
		err = rows.Scan(&id, &first_Name, &last_Name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, first_Name, last_Name)
	}
	err = rows.Err()

	if err != nil {
		panic(err)
	}
}
