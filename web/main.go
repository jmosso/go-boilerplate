package main

import (
  "net/http"
  "log"
  "github.com/jmosso/golang/crudbypackages/pkg/handlers"
)

func main() {

  // URL: Home page
  http.HandleFunc("/", handlers.Home)

  // URL: Create new employee (func create + func insert)
  // '/create' calls func. Create that exec. create.tmpl (POST form with Name and Email)
  http.HandleFunc("/create", handlers.Create)
  // '/insert' calls func. Insert which inserts the received data from create.tmpl into the DB
  http.HandleFunc("/insert", handlers.Insert)

  // URL: Delete an employee using Id
  // '/delete' calls func. Delete which deletes an employee based on a GET request delete?id=Id
  http.HandleFunc("/delete", handlers.Delete)

  // URL: Edit employee data using its Id
  // '/edit' calls func. Edit which retieves employee data based on a GET request edit?id=Id
  // On click it calls '/update' function with a POST (id, name, and email)
  http.HandleFunc("/edit", handlers.Edit)
  // 'update' takes data from update->POST and executes an SQL UPDATE command
  http.HandleFunc("/update", handlers.Update)

  // HTTP server initialization (loop)
  log.Println("Server running...")
  http.ListenAndServe(":8080", nil)

}
