package service

import (
	"database/sql"
	//"kit/log"
	//"os"

	dt "cargo-tracking/datastruct"

	_ "github.com/go-sql-driver/mysql"
)

func dbcon() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "cargo"
	dbIp := "127.0.0.1:3306"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIp+")/"+dbName)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func UpdateStatus(a dt.Delivery) []dt.Delivery {
	db := dbcon()
	updDb, err := db.Query(`UPDATE t_trx_delivery AS x INNER JOIN t_mtr_route_spec AS w on x.fk_id_route_spec = w.id_route_spec SET routing_status="load", transport_status="docking" where x.last_known_location = w.destination`)

	if err != nil {
		panic(err.Error())
	}

	//defer db.Close()

	status := dt.Delivery{}
	res := []dt.Delivery{}

	for updDb.Next() {
		var routing_status, transport_status, last_known_location string
		var fk_id_route_spec, fk_id_itenary, idt_delivery int

		err = updDb.Scan(&idt_delivery, &routing_status, &transport_status, &last_known_location, &fk_id_itenary, &fk_id_route_spec)
		if err != nil {
			panic(err.Error())
		}
		status.DELIVERY_ID = idt_delivery
		status.ROUTING_STATUS = routing_status
		status.TRANSPORT_STATUS = transport_status
		status.LAST_KNOWN_LOCATION = last_known_location
		status.ITENARY_ID = fk_id_itenary
		status.ROUTE_ID = fk_id_route_spec
		res = append(res, status)
	}
	return res
}
