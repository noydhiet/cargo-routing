package service

import (
	"database/sql"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	dt "cargo-routing/datastruct"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "cargo"
	dbIP := "127.0.0.1"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func updateStatus(a dt.UpdateStatusDeliveryRequest, dt.UpdateStatusDeliveryResponse)  {
	
	// var message string
	
	db := dbConn()

	updForm, err := db.Prepare("UPDATE t_trx_delivery  WHERE route_id = ? AND itenary_id = ?")

	if err != nil {
		panic(err.Error())
	}

	updForm.Exec(dt.UpdateStatusDeliveryRequest)
	message :="data berhasil di update"
	return message
	defer db.Close()

}
