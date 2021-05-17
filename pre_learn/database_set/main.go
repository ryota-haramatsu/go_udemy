package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
			name STRING,
			age  INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// インサート
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Nancy", 19)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// アップデート
	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 19, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// マルチセレクトの場合 Query()
	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd)
	// defer rows.Close()

	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	if err := rows.Scan(&p.Name, &p.Age); err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p) // Person型のスライスにappend
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// シングルセレクト QueryRow()
	// cmd = "SELECT * FROM person WHERE age = ?"
	// row := DbConnection.QueryRow(cmd, 19)
	// var p Person
	// if err := row.Scan(&p.Name, &p.Age); err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("NO Row")
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// デリート
	// cmd = `DELETE FROM person WHERE name = ?`
	// _, err = DbConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// tableName := "person"
	// cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
}
