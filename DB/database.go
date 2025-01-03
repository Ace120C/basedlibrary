package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type BooksDb struct {
  db *sql.DB
}

type Book struct {
  ID int `json:"book_id"`
  Title string `json:"name"`
  Link string `json:"LK"`
}

func main()  {
  //opens the database, requires a SQL driver
  conn, err := sql.Open("sqlite3", "books.db")
  //error handling
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  //verfiying the connection with ping()
  err = conn.Ping()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Connected Successfully.")

  rows, err := conn.Query("SELECT * FROM lightnovels")
  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  var data []Book

  for rows.Next(){
    var book Book
    if err := rows.Scan(&book.ID, &book.Title, &book.Link); err != nil {
      log.Fatal(err)
    }
    data = append(data, book)
  }

  if err = rows.Err(); err != nil {
    log.Fatal(err)
  }

  jsonData, err := json.MarshalIndent(data, "", " ")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(jsonData))

  dbFileJSON := "dbOUT.json" 
  content := string(jsonData)

  fileJSON, err := os.Create(dbFileJSON)
  if err != nil {
    log.Fatal(err)
  }
  defer fileJSON.Close()

  _, err = fileJSON.WriteString(content)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("Database has been exported successfully!")

  //input SQL data
  // var createTableQuery string = `
  // CREATE TABLE IF NOT EXISTS lightnovels (
  //   book_id INTERGER PRIMARY KEY,
  //   name TEXT,
  //   LK TEXT
  // );
  // `
  var createValues string = `
   INSERT OR REPLACE INTO lightnovels VALUES (
     1, 'Ascendance of a Bookworm - Vol1 P1','https://drive.google.com/file/d/1q0rsMEev7TnITti0j2SLn5tR5Z-xXmFO/view?usp=sharing' 
   );
   `

   // _, err = conn.Exec(createTableQuery)
   // if err != nil {
   //   log.Fatal(err)
   // }
   // fmt.Println("Table created (or already exists).")

   _, err = conn.Exec(createValues)
   if err != nil {
     log.Fatal(err)
   }
   fmt.Println("Values inserted into table (or already exists.)")

}
