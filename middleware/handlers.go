package middleware

import (
	"io/ioutil"
	"bytes"
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

	// pq lib
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

// AppRunning test if running
func AppRunning(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/running" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "App is running!")
}

// CreatePatrol create a patrol in the postgres db
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

// CreateTask create a task in the postgres db
func CreateTask(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		panic("Unable to decode the request body.")

	}

	insertID := insertTask(task)

	notifyTrafficController(insertID)

	res := response{
		ID:      insertID,
		Message: "Task created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

// CreateGraph create a graph in the postgres db
func CreateGraph(w http.ResponseWriter, r *http.Request) {

	var graph models.Graph

	err := json.NewDecoder(r.Body).Decode(&graph)

	if err != nil {

		panic("Unable to decode the request body.")

	}

	insertID := insertGraph(graph)

	res := response{
		ID:      insertID,
		Message: "Graph created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// CreateRobot create a robot in the postgres db
func CreateRobot(w http.ResponseWriter, r *http.Request) {

	var robot models.Robot

	err := json.NewDecoder(r.Body).Decode(&robot)

	if err != nil {

		panic("Unable to decode the request body.")

	}

	insertID := insertRobot(robot)

	res := response{
		ID:      insertID,
		Message: "Robot created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// CreateCollection create a collection in the postgres db
func CreateCollection(w http.ResponseWriter, r *http.Request) {

	var collection models.Collection

	err := json.NewDecoder(r.Body).Decode(&collection)

	if err != nil {

		panic("Unable to decode the request body.")

	}

	insertID := insertCollection(collection)

	res := response{
		ID:      insertID,
		Message: "Collection created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

// GetGraph will return a single graph by its id
func GetGraph(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	graph, err := getGraph(int64(id))

	if err != nil {
		log.Fatalf("Unable to get graph. %v", err)
	}

	json.NewEncoder(w).Encode(graph)
}

// GetAllGraphDetailed will return all the graphs in detail
func GetAllGraphDetailed(w http.ResponseWriter, r *http.Request) {

	graphs, err := getAllGraphDetailed()

	if err != nil {
		log.Fatalf("Unable to get all graph. %v", err)
	}

	json.NewEncoder(w).Encode(graphs)
}

// GetAllGraphNonDetailed will return all the graphs not in detail
func GetAllGraphNonDetailed(w http.ResponseWriter, r *http.Request) {

	graphs, err := getAllGraphNonDetailed()

	if err != nil {
		log.Fatalf("Unable to get all graph. %v", err)
	}

	json.NewEncoder(w).Encode(graphs)
}

// GetPatrol will return a single patrol by its id
func GetPatrol(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	patrol, err := getPatrol(int64(id))

	if err != nil {
		log.Fatalf("Unable to get patrol. %v", err)
	}

	json.NewEncoder(w).Encode(patrol)
}

// GetTask will return a single task by its id
func GetTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	task, err := getTask(int64(id))

	if err != nil {
		log.Fatalf("Unable to get patrol. %v", err)
	}

	json.NewEncoder(w).Encode(task)
}


// GetAllPatrol will return all the patrols
func GetAllPatrol(w http.ResponseWriter, r *http.Request) {

	patrols, err := getAllPatrol()

	if err != nil {
		log.Fatalf("Unable to get all patrol. %v", err)
	}

	json.NewEncoder(w).Encode(patrols)
}

// GetAllRobots will return all robots
func GetAllRobots(w http.ResponseWriter, r *http.Request) {

	robots, err := getAllRobots()

	if err != nil {
		log.Fatalf("Unable to get all robots. %v", err)
	}

	json.NewEncoder(w).Encode(robots)

}

// GetAllTasks will return all goto tasks
func GetAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := getAllTasks()

	if err != nil {
		log.Fatalf("Unable to get all goto tasks. %v", err)
	}

	json.NewEncoder(w).Encode(tasks)

}

// GetAllPatrolTasks will return all patrol tasks
func GetAllPatrolTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := getAllPatrolTasks()

	if err != nil {
		log.Fatalf("Unable to get all patrol tasks. %v", err)
	}

	json.NewEncoder(w).Encode(tasks)

}

// GetAllGotoTasks will return all goto tasks
func GetAllGotoTasks(w http.ResponseWriter, r *http.Request) {

	tasks, err := getAllGotoTasks()

	if err != nil {
		log.Fatalf("Unable to get all goto tasks. %v", err)
	}

	json.NewEncoder(w).Encode(tasks)

}

// GetAllCollection will return all collections
func GetAllCollection(w http.ResponseWriter, r *http.Request) {

	tasks, err := getAllCollection()

	if err != nil {
		log.Fatalf("Unable to get all collections. %v", err)
	}

	json.NewEncoder(w).Encode(tasks)

}

// UpdateGraph update graph's detail in the postgres db
func UpdateGraph(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var graph models.Graph

	err = json.NewDecoder(r.Body).Decode(&graph)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateGraph(int64(id), graph)

	msg := fmt.Sprintf("Graph updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
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

// DeleteGraph delete graph's detail in the postgres db
func DeleteGraph(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteGraph(int64(id))

	msg := fmt.Sprintf("Graph updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeletePatrol delete patrol's detail in the postgres db
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

// DeleteRobot delete robot's detail in the postgres db
func DeleteRobot(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteRobot(int64(id))

	msg := fmt.Sprintf("Robot deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteTask delete task's detail in the postgres db
func DeleteTask(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteTask(int64(id))

	msg := fmt.Sprintf("Task deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteCollection delete collection's detail in the postgres db
func DeleteCollection(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteCollection(int64(id))

	msg := fmt.Sprintf("Collection deleted successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// insert one graph in the DB
func insertGraph(graph models.Graph) int64 {

	db := createConnection()

	defer db.Close()

	// create the insert sql query
	// returning graphid will return the id of the inserted graph
	sqlStatement := `INSERT INTO graphs ("mapVerID", "collectionID", scale, name, location, level, lanes, vertices) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING graphid`

	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, graph.MapVerID, graph.CollectionID, graph.Scale, graph.Name, graph.Location, graph.Level, graph.Lanes, graph.Vertices).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println("Inserted a single record as id:", id)

	return id
}

// insert one patrol in the DB
func insertPatrol(patrol models.Patrol) int64 {

	db := createConnection()

	defer db.Close()

	// create the insert sql query
	// returning patrolid will return the id of the inserted patrol
	sqlStatement := `INSERT INTO patrols ("graphID", "mapVerID", name, points) VALUES ($1, $2, $3, $4) RETURNING patrolid`

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

// insert one collection in the DB
func insertCollection(collection models.Collection) int64 {

	db := createConnection()

	defer db.Close()

	// create the insert sql query
	// returning graphid will return the id of the inserted graph
	sqlStatement := `INSERT INTO collections (name) VALUES ($1) RETURNING collectionid`

	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, collection.Name).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println("Inserted a single record as id:", id)

	return id
}

// insert one robot in the DB
func insertRobot(robot models.Robot) int64 {

	db := createConnection()

	defer db.Close()

	// create the insert sql query
	// returning robotid will return the id of the inserted patrol
	sqlStatement := `INSERT INTO robots ("robotID", name) VALUES ($1, $2) RETURNING robotid`

	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, robot.RobotID, robot.Name).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println("Inserted a single record as id:", id)

	return id
}

// insert one task in the DB
func insertTask(task models.Task) int64 {

	db := createConnection()

	defer db.Close()

	// create the insert sql query
	// returning id will return the id of the inserted task
	sqlStatement := `INSERT INTO tasks (type, "taskDetails") VALUES ($1, $2) RETURNING taskid`

	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, task.Type, task.TaskDetails).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Println("Inserted a single record as id:", id)

	return id
}

// get one graph from the DB by its id
func getGraph(id int64) (models.Graph, error) {

	db := createConnection()

	defer db.Close()

	var graph models.Graph

	sqlStatement := `SELECT * FROM graphs WHERE graphid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&graph.ID, &graph.MapVerID, &graph.CollectionID, &graph.Scale, &graph.Name, &graph.Location, &graph.Level, &graph.Lanes, &graph.Vertices)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return graph, nil
	case nil:
		return graph, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return graph, err
}

// get one patrolfrom the DB by its id
func getPatrol(id int64) (models.Patrol, error) {

	db := createConnection()

	defer db.Close()

	var patrol models.Patrol

	sqlStatement := `SELECT * FROM patrols WHERE "graphID"=$1`

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

// get one graph from the DB by its id
func getTask(id int64) (models.Task, error) {

	db := createConnection()

	defer db.Close()

	var task models.Task

	sqlStatement := `SELECT * FROM tasks WHERE taskid=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&task.ID, &task.Type, &task.TaskDetails)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return task, nil
	case nil:
		return task, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return task, err
}

// get all graph in detail from the DB
func getAllGraphDetailed() ([]models.Graph, error) {

	db := createConnection()

	defer db.Close()

	var graphs []models.Graph

	sqlStatement := `SELECT * FROM graphs`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var graph models.Graph

		err = rows.Scan(&graph.ID, &graph.MapVerID, &graph.CollectionID, &graph.Scale, &graph.Name, &graph.Location, &graph.Level, &graph.Lanes, &graph.Vertices)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		graphs = append(graphs, graph)

	}

	return graphs, err
}

// get all graph not in detail from the DB
func getAllGraphNonDetailed() ([]models.GraphNonDetailed, error) {

	db := createConnection()

	defer db.Close()

	var graphs []models.GraphNonDetailed

	sqlStatement := `SELECT graphid, "collectionID", name, location FROM graphs`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var graph models.GraphNonDetailed

		err = rows.Scan(&graph.ID, &graph.CollectionID, &graph.Name, &graph.Location)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		graphs = append(graphs, graph)

	}

	return graphs, err
}

// get all patrol from the DB
func getAllPatrol() ([]models.Patrol, error) {

	db := createConnection()

	defer db.Close()

	var patrols []models.Patrol

	sqlStatement := `SELECT * FROM patrols`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query there. %v", err)
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

// get all robots from the DB
func getAllRobots() ([]models.Robot, error) {

	db := createConnection()

	defer db.Close()

	var robots []models.Robot

	sqlStatement := `SELECT * FROM robots`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var robot models.Robot

		err = rows.Scan(&robot.ID, &robot.RobotID, &robot.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		robots = append(robots, robot)

	}

	return robots, err
}

// get all patrol task from the DB
func getAllTasks() ([]models.Task, error) {

	db := createConnection()

	defer db.Close()

	var tasks []models.Task

	sqlStatement := `SELECT * FROM tasks`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var task models.Task

		err = rows.Scan(&task.ID, &task.Type, &task.TaskDetails)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		tasks = append(tasks, task)

	}

	return tasks, err
}

// get all patrol task from the DB
func getAllPatrolTasks() ([]models.Task, error) {

	db := createConnection()

	defer db.Close()

	var tasks []models.Task

	sqlStatement := `SELECT * FROM tasks where type = 0`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var task models.Task

		err = rows.Scan(&task.ID, &task.Type, &task.TaskDetails)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		tasks = append(tasks, task)

	}

	return tasks, err
}

// get all goto tasks from the DB
func getAllGotoTasks() ([]models.Task, error) {

	db := createConnection()

	defer db.Close()

	var tasks []models.Task

	sqlStatement := `SELECT * FROM tasks where type = 1`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var task models.Task

		err = rows.Scan(&task.ID, &task.Type, &task.TaskDetails)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		tasks = append(tasks, task)

	}

	return tasks, err
}

// get all collection from the DB
func getAllCollection() ([]models.Collection, error) {

	db := createConnection()

	defer db.Close()

	var collections []models.Collection

	sqlStatement := `SELECT * FROM collections`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {

		var collection models.Collection

		err = rows.Scan(&collection.ID, &collection.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		collections = append(collections, collection)

	}

	return collections, err
}

// update graph in the DB
func updateGraph(id int64, graph models.Graph) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE graphs SET "mapVerID"=$2, "collectionID"=$3, scale=$4, name=$5, location=$6, level=$7, lanes=$8, vertices=$9 WHERE graphid=$1`

	res, err := db.Exec(sqlStatement, id, graph.MapVerID, graph.CollectionID, graph.Scale, graph.Name, graph.Location, graph.Level, graph.Lanes, graph.Vertices)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v. ", rowsAffected)

	return rowsAffected
}

// update patrol in the DB
func updatePatrol(id int64, patrol models.Patrol) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE patrols SET "graphID"=$2, "mapVerID"=$3, name=$4, points=$5 WHERE patrolid=$1`

	res, err := db.Exec(sqlStatement, id, patrol.GraphID, patrol.MapVerID, patrol.Name, patrol.Points)

	if err != nil {
		log.Fatalf("Unable to execute the query here. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v. ", rowsAffected)

	return rowsAffected
}

// delete graph in the DB
func deleteGraph(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM graphs WHERE graphid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Println("Total rows/record affected", rowsAffected)

	return rowsAffected
}

// delete patrol in the DB
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

// delete robot in the DB
func deleteRobot(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM robots WHERE robotid=$1`

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

// delete task in the DB
func deleteTask(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM tasks WHERE taskid=$1`

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

// delete collection in the DB
func deleteCollection(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM collections WHERE collectionid=$1`

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

// notifyTrafficController notifies traffic controller when new task is added
func notifyTrafficController(insertID int64) string {

	str := strconv.FormatInt(insertID, 10)
	values := map[string]string{"taskID":str}
	requestBody, err := json.Marshal(values)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("http://18.140.162.221:3003/enqueue", "application/json",
		bytes.NewBuffer(requestBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}