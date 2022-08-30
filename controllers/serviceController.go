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
	createService = `INSERT INTO services (name,price,rate) VALUES ($1, $2, $3) RETURNING id, name, price, rate`

	getService = `SELECT id,name,price,rate FROM services WHERE id = $1 LIMIT 1`

	getServiceByName = `SELECT id, name,price,rate FROM services WHERE name = $1 LIMIT 1`

	listService = `SELECT id, name,price,rate FROM services ORDER BY id`

	deleteService = `DELETE FROM services WHERE id = $1`

	updateService = `UPDATE services set name = $2,price= $3,rate= $4 WHERE id = $1 RETURNING id, ip, name,price,rate`
)

func GetServices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query(context.Background(), listService)

	if err != nil {
		log.Println(err)
	}
	var servicelist []models.Service
	for rows.Next() {
		var s models.Service

		err := rows.Scan(&s.Id, &s.Name, &s.Price, &s.Rate)

		if err != nil {
			log.Println("Error while scan rows")
			return

		}
		servicelist = append(servicelist, s)
	}

	rows.Close()

	res, _ := json.Marshal(servicelist)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetServiceById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	serviceId := vars["serviceId"]
	ID, err := strconv.ParseInt(serviceId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing serviceId")
	}

	var s models.Service

	err = db.QueryRow(context.Background(), getService, ID).Scan(&s.Id, &s.Name, &s.Price, &s.Rate)

	if err != nil {
		msg := fmt.Sprintf("not found service with ID %s", serviceId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}

	res, _ := json.Marshal(s)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateService(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var s1 models.CreateServiceParams
	var service models.Service

	_ = json.NewDecoder(r.Body).Decode(&s1)

	err := db.QueryRow(ctx, getServiceByName, s1.Name).Scan(&service.Id, &service.Name, &service.Price, &service.Rate)

	if err == nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Service with same name already exists"))
		return

	}

	err = db.QueryRow(ctx, createService, s1.Name, s1.Price, s1.Rate).Scan(&service.Id, &service.Name, &service.Price, &service.Rate)

	if err != nil {
		fmt.Println("error while Create the service")
	}

	res, _ := json.Marshal(service)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateService(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var s1, service models.Service
	_ = json.NewDecoder(r.Body).Decode(&s1)

	err := db.QueryRow(ctx, updateService, s1.Id, s1.Name, s1.Price, s1.Rate).Scan(&service.Id, &service.Name, &service.Price, &service.Rate)

	if err != nil {
		log.Println(err)
	}

	res, _ := json.Marshal(service)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx = context.Background()

	db := database.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	serviceId := vars["serviceId"]
	ID, err := strconv.ParseInt(serviceId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing serviceId")
	}

	b, _ := db.Exec(ctx, deleteService, ID)

	if b.RowsAffected() == 0 {
		msg := fmt.Sprintf("not found service with ID %s", serviceId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusOK)
}
