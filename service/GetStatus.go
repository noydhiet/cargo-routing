package service

// import (
// 	"database/sql"
// 	//"kit/log"
// 	//"os"

// 	dt "cargo-tracking/datastruct"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func dbcon() (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "cargo"
// 	dbIp := "127.0.0.1"

// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbIp+")/"+dbName)

// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// func GetStatus(a dt.Delivery) []dt.Delivery {
// 	db := dbcon()
// 	selDb, err := db.Query("SELECT * from t_trx_delivery where fk_id_itenary=? or fk_id_route_spec=? or idt_delivery=?", a.ITENARY_ID, a.ROUTE_ID, a.DELIVERY_ID)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	status := dt.Delivery{}
// 	res := []dt.Delivery{}

// 	for selDb.Next() {
// 		var routing_status, transport_status, last_known_location string
// 		var route_id, itenary_id, delivery_id int

// 		err = selDb.Scan(&delivery_id, &routing_status, &transport_status, &last_known_location, &itenary_id, &route_id)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		status.DELIVERY_ID = delivery_id
// 		status.ROUTING_STATUS = routing_status
// 		status.TRANSPORT_STATUS = transport_status
// 		status.LAST_KNOWN_LOCATION = last_known_location
// 		status.ITENARY_ID = itenary_id
// 		status.ROUTE_ID = route_id
// 		//println("ROUTING_STATUS %s, TRANSPORT_STATUS %s, LAST_KNOWN_LOCATION%s", routing_status, transport_status, last_known_location)
// 		res = append(res, status)
// 	}
// 	return res
// }
