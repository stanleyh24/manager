package repositories

import (
	"context"
	"log"

	"github.com/stanleyh24/manager/database"
	"github.com/stanleyh24/manager/models"
)

/* type router interface {
	Getall() models.Routers
	Getbyid(id int) (*models.Router, error)
	Create(r models.Createrouter) (*models.Router, error)
	Update(id int, r models.Router) (*models.Router, error)
	Delete(id int) error
} */
const (
	createRouter = `INSERT INTO router (ip,name,username,password) VALUES ($1, $2, $3, $4) RETURNING id, ip, name, username, password`

	getRouter = `SELECT id, ip, name, username, password FROM router WHERE id = $1 LIMIT 1`

	listRouter   = `SELECT id, ip, name, username, password FROM router ORDER BY id`
	deleteRouter = `DELETE FROM router WHERE id = $1`

	updateRouter = `UPDATE router set ip = $2,name= $3,username= $4,password= $5 WHERE id = $1 RETURNING id, ip, name, username, password`
)

func Getall() ([]models.Router, error) {
	db := database.ConnectDB()
	defer db.Close(context.Background())

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
			return nil, err

		}
		routerlist = append(routerlist, r)
	}

	return routerlist, nil
}

func Getbyid(id int) (*models.Router, error) {
	return nil, nil
}

func Create(r models.CreateRouterParams) (*models.Router, error) {
	return nil, nil
}

func Update(id int, r models.Router) (*models.Router, error) {
	return nil, nil
}

func Delete(id int) error {
	return nil
}
