package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//conexion a database
//username:password@tcp(localhost:3306)/database
const url = "root:123@tcp(localhost:3306)/webGo.db"

var db *sql.DB

func conex() {
	//requiere nombre del driver y de la bd de datos
	//esta funcion devuelve la conexion y un error
	conexion, err := sql.Open("mysql", url)

	if err != nil {

		panic(err)

	} else {
		fmt.Println("conexion exitosa")
		db = conexion

	}

}
func closeDB() {
	db.Close()

}
