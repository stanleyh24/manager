package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/stanleyh24/manager/database"
	"github.com/stanleyh24/manager/models"
)

const (
	createRouter = `INSERT INTO router (ip,name,username,password) VALUES ($1, $2, $3, $4) RETURNING id, ip, name, username, password`

	getRouter = `SELECT id, ip, name, username, password FROM router WHERE id = $1 LIMIT 1`

	getRouterByIp = `SELECT id, ip, name, username, password FROM router WHERE ip = $1 LIMIT 1`

	listRouter = `SELECT id, ip, name, username, password FROM router ORDER BY id`

	deleteRouter = `DELETE FROM router WHERE id = $1`

	updateRouter = `UPDATE router set ip = $2,name= $3,username= $4,password= $5 WHERE id = $1 RETURNING id, ip, name, username, password`
)

func GetRouters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query(context.Background(), listRouter)

	if err != nil {
		log.Println(err)
	}
	var routerlist []models.Router
	for rows.Next() {
		var r models.Router

		err := rows.Scan(&r.ID, &r.Ip, &r.Name, &r.Username, &r.Password)

		if err != nil {
			log.Println("Error while scan rows")
			return

		}
		routerlist = append(routerlist, r)
	}

	rows.Close()

	res, _ := json.Marshal(routerlist)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetRouterById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	routerId := vars["routerId"]
	ID, err := strconv.ParseInt(routerId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing routerId")
	}
	var router models.Router

	err = db.QueryRow(context.Background(), getRouter, ID).Scan(&router.ID, &router.Ip, &router.Name, &router.Username, &router.Password)

	if err != nil {
		msg := fmt.Sprintf("not found router with ID %s", routerId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}

	res, _ := json.Marshal(router)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateRouter(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var r1 models.CreateRouterParams
	var router models.Router

	_ = json.NewDecoder(r.Body).Decode(&r1)

	err := db.QueryRow(ctx, getRouterByIp, r1.Ip).Scan(&router.ID, &router.Ip, &router.Name, &router.Username, &router.Password)

	if err == nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ip address already exists"))
		return

	}

	err = db.QueryRow(ctx, createRouter, r1.Ip, r1.Name, r1.Username, r1.Password).Scan(&router.ID, &router.Ip, &router.Name, &router.Username, &router.Password)

	if err != nil {
		fmt.Println("error while Create a router")
	}

	res, _ := json.Marshal(router)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateRouter(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var r1, router models.Router
	_ = json.NewDecoder(r.Body).Decode(&r1)

	err := db.QueryRow(ctx, updateRouter, r1.ID, r1.Ip, r1.Name, r1.Username, r1.Password).Scan(&router.ID, &router.Ip, &router.Name, &router.Username, &router.Password)

	if err != nil {
		log.Println(err)
	}

	res, _ := json.Marshal(router)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteRouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx = context.Background()

	db := database.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	routerId := vars["routerId"]
	ID, err := strconv.ParseInt(routerId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing routerId")
	}

	b, _ := db.Exec(ctx, deleteRouter, ID)

	if b.RowsAffected() == 0 {
		msg := fmt.Sprintf("not found router with ID %s", routerId)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusOK)
}
