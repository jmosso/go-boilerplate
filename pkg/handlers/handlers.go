package handlers

import (
  "net/http"
  //"fmt"
  "html/template"
  "github.com/jmosso/golang/crudbypackages/pkg/dbc"
)

// Templates directory (.tmpl)
var templates= template.Must(template.ParseGlob("web/templates/*"))
var err error

// Struct to read from DB, replicates empoyees table's columns
type Employee struct {
    Id    int
    Name string
    Email string
}

// Home page based on 'home.tmpl'
func Home(w http.ResponseWriter, r *http.Request){

  // create a db connection and defer close it
  db := dbc.ConnectDB()
  defer db.Close()

  // Executes a * Query and return all employees, then load them into rows
  // 'rows' is an array of structs
  rows, err := db.Query("SELECT * FROM employees ORDER BY id ASC")
  CheckError(err)
  defer rows.Close()

  // Estructura para guardar los datos que vienen del query SQL
  employee := Employee{}
  groupEmployee := []Employee{}

  // Read the array of structs, line by line and add it to 'groupEmployee'
  for rows.Next() {
    var id int
    var name, email string
    err = rows.Scan(&id, &name, &email)
    CheckError(err)
    employee.Id=id
    employee.Name=name
    employee.Email=email
    groupEmployee=append(groupEmployee, employee)
  }
  // Print the whole array of structs 'groupEmployee'
  // fmt.Println(groupEmployee)

  // Calls the home template and send data to it to be presented in a table
  templates.ExecuteTemplate(w, "home", groupEmployee)
}


// Function to create a new employee, triggers 'create' template that calls 'Instert' function
func Create(w http.ResponseWriter, r *http.Request){
  // muestra texto estatico en browser
  // fmt.Fprintf(w, "Hola Gregorio")

  // con la variable nil se enviarian valores al template
  templates.ExecuteTemplate(w, "create", nil)
}

// Create an employee in the postgres db
func Insert(w http.ResponseWriter, r *http.Request) {

  if r.Method == "POST" {
    name := r.FormValue("name")           // Values from Form "id" field
    email := r.FormValue("email")

    // create a db connection and defer close it
    db := dbc.ConnectDB()
    defer db.Close()

    // create the insert sql query
    insertDynStmt := `insert into "employees"("name", "email") values($1, $2)`
    _, err = db.Exec(insertDynStmt, name, email)
    CheckError(err)

    //Registro transaccion en terminal (log)
    // fmt.Printf("Insertando usuario %s con email %s:\n", name, email)
    // log.Println("Insertando usuario: \"" + name + "\", con email: \"" + email + "\"")

    // Devolvemos el control a la pagina de inicio / con msg 301
    http.Redirect(w,r,"/",301)
  }

}

// Delete an employee in the postgres db by 'id'
func Delete(w http.ResponseWriter, r *http.Request) {

  idEmployee := r.URL.Query().Get("id")
  //fmt.Println(idEmployee)

  // create a db connection and defer close it
  db := dbc.ConnectDB()
  defer db.Close()

  // create the delete sql query
  delEmploy, err := db.Prepare(`delete from employees where id=$1`)
  CheckError(err)
  delEmploy.Exec(idEmployee)
  CheckError(err)

  //Registro transaccion en terminal (log)
  // fmt.Printf("Insertando usuario %s con email %s:\n", name, email)
  // log.Println("Deleted usuario: \"" + idEmployee + "\"")

  // Devolvemos el control a la pagina de inicio / con msg 301
  http.Redirect(w,r,"/",301)
}

// Edit a employee from/in table with 'id'
func Edit(w http.ResponseWriter, r *http.Request) {

  idEmployee := r.URL.Query().Get("id")
  //fmt.Println(idEmployee)

  // create a db connection and defer close it
  db := dbc.ConnectDB()
  defer db.Close()

  // create the sql select query based on 'idEmployee'
  row, err := db.Query("SELECT * FROM employees where id = $1", idEmployee)
  CheckError(err)
  defer row.Close()

  // Estructura para guardar los datos que vienen del query SQL
  employee := Employee{}

  // llemos los valores, linea a linea
  for row.Next() {
      var id int
      var name, email string
      err = row.Scan(&id, &name, &email)
      CheckError(err)
      employee.Id=id
      employee.Name=name
      employee.Email=email
      }
      // fmt.Println(employee)
      templates.ExecuteTemplate(w, "edit", employee)

}

// Update Employee from ./edit
func Update(w http.ResponseWriter, r *http.Request) {

  //Values from Form fields
  if r.Method == "POST" {
    id := r.FormValue("id")
    name := r.FormValue("name")
    email := r.FormValue("email")

    // create a db connection and defer close it
    db := dbc.ConnectDB()
    defer db.Close()

    // Create the update sql query based on 'id'
    updateStmt := `update employees set name=$1, email=$2 where id=$3`
    _, err := db.Exec(updateStmt, name, email, id)
    CheckError(err)

    //Registro transaccion en terminal (log)
    //fmt.Printf("Updating employee %s con email %s %s:\n", name, email, id)

    // Devolvemos el control a la pagina de inicio / con msg 301
    http.Redirect(w,r,"/",301)
  }

}

// General purpose error checking function
func CheckError(err error) {
  if err != nil {
    panic(err.Error())
  }
}
