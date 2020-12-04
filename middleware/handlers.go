package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-task/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return db
}

// CreatePatrol create a user in the postgres db
func CreatePatrol(w http.ResponseWriter, r *http.Request) {

	var patrol models.Patrol

	err := json.NewDecoder(r.Body).Decode(&patrol)

	if err != nil {

		panic("Unable to decode the request body.")

	}

	insertID := insertPatrol(patrol)

	res := response{
		ID:      insertID,
		Message: "Patrol created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// GetPatrol will return a single graph by its id
func GetPatrol(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	patrol, err := getPatrol(int64(id))

	if err != nil {
		log.Fatalf("Unable to get graph. %v", err)
	}

	json.NewEncoder(w).Encode(patrol)
}

// GetAllPatrol will return all the patrols
func GetAllPatrol(w http.ResponseWriter, r *http.Request) {

	patrols, err := getAllPatrol()

	if err != nil {
		log.Fatalf("Unable to get all graph. %v", err)
	}

	json.NewEncoder(w).Encode(patrols)
}

// UpdatePatrol update patrol's detail in the postgres db
func UpdatePatrol(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var patrol models.Patrol

	err = json.NewDecoder(r.Body).Decode(&patrol)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updatePatrol(int64(id), patrol)

	msg := fmt.Sprintf("Patrol updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeletePatrol delete graph's detail in the postgres db
func DeletePatrol(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deletePatrol(int64(id))

	msg := fmt.Sprintf("Patrol deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// insert one patrol in the DB
func insertPatrol(patrol models.Patrol) int64 {

	db := createConnection()

	defer db.Close()

	// create the insert sql query
	// returning patrolid will return the id of the inserted patrol
	sqlStatement := `INSERT INTO patrols (graphID, mapVerID, name, points) VALUES ($1, $2, $3, $4) RETURNING patrolid`

	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, patrol.GraphID, patrol.MapVerID, patrol.Name, patrol.Points).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println("Inserted a single record as id:", id)

	return id
}

// get one graph from the DB by its id
func getPatrol(id int64) (models.Patrol, error) {

	db := createConnection()

	defer db.Close()

	var patrol models.Patrol

	sqlStatement := `SELECT * FROM patrols WHERE graphID=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&patrol.ID, &patrol.GraphID, &patrol.MapVerID, &patrol.Name, &patrol.Points)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return patrol, nil
	case nil:
		return patrol, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return patrol, err
}

// get all patrol from the DB
func getAllPatrol() ([]models.Patrol, error) {

	db := createConnection()

	defer db.Close()

	var patrols []models.Patrol

	sqlStatement := `SELECT * FROM patrols`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var patrol models.Patrol

		err = rows.Scan(&patrol.ID, &patrol.GraphID, &patrol.MapVerID, &patrol.Name, &patrol.Points)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		patrols = append(patrols, patrol)

	}

	return patrols, err
}

// update patrol in the DB
func updatePatrol(id int64, patrol models.Patrol) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE patrols SET graphID=$2, mapVerID=$3, name=$4, points=$5 WHERE patrolid=$1`

	res, err := db.Exec(sqlStatement, id, patrol.GraphID, patrol.MapVerID, patrol.Name, patrol.Points)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete graph in the DB
func deletePatrol(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM patrols WHERE patrolid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Println("Total rows/record affected ", rowsAffected)

	return rowsAffected
}
