package main

import (
	//formato entrada y salida
	"database/sql"
	"log"           //datos por consola
	"net/http"      //mostrar web
	"text/template" //

	_ "github.com/go-sql-driver/mysql" //paquete de conexión a myql (_ es para los dirver) ---terceros
	//"database/sql"
)

var templates = template.Must(template.ParseGlob("templates/*"))

type Products struct {
	Id    int
	Name  string
	Price float64
	Img   string
}

type User struct {
	Id    int
	Name  string
	Email string
	Pass  string
}

func Log(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola desde index")
	templates.ExecuteTemplate(w, "lgn.tm", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	connon := connect()

	if r.Method == "POST" {
		name := r.FormValue("Nameus")
		pass := r.FormValue("Passus")

		sele, err := connon.Query("SELECT * FROM empleados")

		if err != nil {
			panic(err.Error())
		}

		user := User{}
		if err != nil {
			panic(err.Error())
		}
		for sele.Next() {
			var id int
			var nombre, correo, Contrasena string
			err = sele.Scan(&id, &nombre, &correo, &Contrasena)
			if err != nil {
				panic(err.Error())
			}
			user.Name = nombre
			user.Pass = Contrasena
			if user.Pass == pass && user.Name == name {
				http.Redirect(w, r, "/index", 301)
				Selectpt(w, r)
			} else {
				http.Redirect(w, r, "/", 301)
			}
		}
		//user.Id = id
		//user.Email = correo

	}

}

func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola desde index")
	templates.ExecuteTemplate(w, "idx.tm", nil)

	//prueba bd
	/*
		connon := connect()
		inst, err := connon.Prepare("INSERT INTO productos(nombre, precio, imagen) VALUES('jabón',5.5,'https://img.freepik.com/vector-gratis/jabon-barra-flotante-jabon-liquido-dibujos-animados-vector-icono-ilustracion-concepto-icono-objeto-sanitario_138676-4675.jpg?size=338&ext=jpg&ga=GA1.1.1826414947.1699401600&semt=sph')")

		if err != nil {
			panic(err.Error())
		}
		inst.Exec()
	*/
	//
}

func Create(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola desde index")
	templates.ExecuteTemplate(w, "cre.tm", nil)
}

func connect() (conn *sql.DB) {
	Driver := "mysql"
	User := "root"
	Pass := ""
	Bd := "productosml"
	conn, err := sql.Open(Driver, User+":"+Pass+"@tcp(127.0.0.1)/"+Bd)
	if err != nil {
		panic(err.Error())
	}
	return conn
}

func Insertpt(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("Namept")
		price := r.FormValue("Pricept")
		img := r.FormValue("Imgpt")

		connon := connect()
		inst, err := connon.Prepare("INSERT INTO productos(nombre, precio, imagen) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		inst.Exec(name, price, img)

		http.Redirect(w, r, "/", 301)
	}
}

func Selectpt(w http.ResponseWriter, r *http.Request) {
	connon := connect()
	sele, err := connon.Query("SELECT * FROM productos")

	if err != nil {
		panic(err.Error())
	}
	product := Products{}
	products := []Products{}

	for sele.Next() {
		var Id int
		var Name, Img string
		var Price float64
		err = sele.Scan(&Id, &Name, &Price, &Img)
		if err != nil {
			panic(err.Error())
		}
		product.Id = Id
		product.Name = Name
		product.Price = Price
		product.Img = Img

		products = append(products, product)
	}
	templates.ExecuteTemplate(w, "idx.tm", products)
}

func main() {
	http.HandleFunc("/", Log)
	http.HandleFunc("/index", Selectpt)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/insertpt", Insertpt)
	http.HandleFunc("/login", Login)

	log.Println("servidor corriendo.")
	http.ListenAndServe(":8080", nil)
}
