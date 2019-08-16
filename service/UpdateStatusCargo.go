package service

import (
	"database/sql"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	dt "cargo-tracking/datastruct"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "PASSWORD"
	dbName := "db_go"
	dbIP := "192.168.20.7"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIP+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func UpdateStatus(a dt.UpdateStatusDeliveryRequest, dt.UpdateStatusDeliveryResponse)  {
	
	var message string
	
	db := dbConn()

	updForm, err := db.Prepare("UPDATE t_trx_delivery  WHERE route_id = ? AND itenary_id = ?")

	if err != nil {
		panic(err.Error())
	}

	updForm.Exec(dt.UpdateStatusDeliveryRequest)
	message := "data berhasil di update"
	return message
	defer db.Close()

}
