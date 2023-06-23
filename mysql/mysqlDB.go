package mysql

import (
	"database/sql"
	"fmt"
	"log"

	entities "github.com/mohammadMghi/sqrs/entities"
	_ "github.com/go-sql-driver/mysql"
)


type MysqlDB struct{
	db *sql.DB 
}

var db *sql.DB

func NewMysql()*MysqlDB{
	var err error
	db , err = ConnectToDB()

	if err != nil{
		log.Fatal(err)
	}
	return &MysqlDB{
	  db: db,
	}
}



func  ConnectToDB()(*sql.DB , error){

	if db != nil{

		return db , nil
	}
	

	connection := "user:password@tcp(hostname:port)/database_name"


	db , err := sql.Open("mysql" ,  connection)
	
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()

	if err != nil{
		log.Fatal(err.Error())
	}


	return db,nil
}


func (mysql *MysqlDB)saveProduct(product entities.Product)error{
    // Prepare the SQL statement
    stmt, err := mysql.db.Prepare("INSERT INTO products (title, price) VALUES (?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the prepared statement with the provided values
    _, err = stmt.Exec(product.Title, product.Price)
    if err != nil {
        return err
    }

    return nil
	

}