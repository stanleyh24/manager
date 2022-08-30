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
	createClient = `INSERT INTO clients (name, last_name, address, phone, payment_date, service_id, router_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, last_name, address, phone,payment_date, service_id, router_id`

	getClient = `SELECT id, name, last_name, address, phone,payment_date, service_id, router_id FROM clients WHERE id = $1 LIMIT 1`

	//getClientByName = `SELECT id, name, last_name, address, phone,payment_date, service_id, router_id FROM clients WHERE name = $1 LIMIT 1`

	listClients = `SELECT id, name, last_name, address, phone,payment_date, service_id, router_id FROM clients ORDER BY id`

	deleteClient = `DELETE FROM clients WHERE id = $1`

	updateClient = `UPDATE clients set name=$2, last_name=$3, address=$4, phone=$5, payment_date=$6, service_id=$7, router_id=$8 WHERE id = $1 RETURNING id, name, last_name, address, phone,payment_date, service_id, router_id`
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	rows, err := db.Query(context.Background(), listClients)

	if err != nil {
		log.Println(err)
	}
	var clientlist []models.Client
	for rows.Next() {
		var client models.Client

		err := rows.Scan(&client.Id, &client.Name, &client.Last_name, &client.Address, &client.Phone, &client.Payment_date, &client.Service_id, &client.Router_id)

		if err != nil {
			log.Println("Error while scan rows")
			return

		}
		clientlist = append(clientlist, client)
	}

	rows.Close()

	res, _ := json.Marshal(clientlist)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetClientById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	clientId := vars["serviceId"]
	ID, err := strconv.ParseInt(clientId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing serviceId")
	}

	var client models.Client

	err = db.QueryRow(context.Background(), getClient, ID).Scan(&client.Id, &client.Name, &client.Last_name, &client.Address, &client.Phone, &client.Payment_date, &client.Service_id, &client.Router_id)

	if err != nil {
		msg := fmt.Sprintf("not found service with ID %s", clientId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}

	res, _ := json.Marshal(client)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var ctx = context.Background()
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var c1 models.CreateClientParams
	var client models.Client

	_ = json.NewDecoder(r.Body).Decode(&c1)

	/* err := db.QueryRow(ctx, getClientByName, c1.Name).Scan(&client.Id, &client.Name, &client.Last_name, &client.Address, &client.Phone, &client.Payment_date, &client.Service_id, &client.Router_id)

	if err == nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Service with same name already exists"))
		return

	} */

	err := db.QueryRow(ctx, createClient, &c1.Name, &c1.Last_name, &c1.Address, &c1.Phone, &c1.Payment_date, &c1.Service_id, &c1.Router_id).Scan(&client.Id, &client.Name, &client.Last_name, &client.Address, &client.Phone, &client.Payment_date, &client.Service_id, &client.Router_id)

	if err != nil {
		fmt.Println("error while Create the client")
	}

	res, _ := json.Marshal(client)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {

	var ctx = context.Background()
	w.Header().Set("Content-Type", "application/json")

	db := database.ConnectDB()
	defer db.Close()

	var c1, client models.Client
	_ = json.NewDecoder(r.Body).Decode(&c1)

	err := db.QueryRow(ctx, updateClient, &c1.Id, &c1.Name, &c1.Last_name, &c1.Address, &c1.Phone, &c1.Payment_date, &c1.Service_id, &c1.Router_id).Scan(&client.Id, &client.Name, &client.Last_name, &client.Address, &client.Phone, &client.Payment_date, &client.Service_id, &client.Router_id)

	if err != nil {
		log.Println(err)
	}

	res, _ := json.Marshal(client)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ctx = context.Background()

	db := database.ConnectDB()
	defer db.Close()

	vars := mux.Vars(r)
	clientId := vars["clientId"]
	ID, err := strconv.ParseInt(clientId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing clientId")
	}

	b, _ := db.Exec(ctx, deleteClient, ID)

	if b.RowsAffected() == 0 {
		msg := fmt.Sprintf("not found client with ID %s", clientId)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(msg))
		return
	}

	w.WriteHeader(http.StatusOK)
}
