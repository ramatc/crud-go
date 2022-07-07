package main

import(
	"fmt"
	"os"
	"bufio" // Para poder realizar la entrada por teclado
	"strings" // Para borrar el salto de linea
	"strconv" // Para convertir de string a entero
	"os/exec" // Para poder limpiar la consola
)

var reader *bufio.Reader

type User struct {
	id int
	username string
	email string
	age int
}

var id int
var users map[int]User

func createUser(){
	clearConsole()

	fmt.Print("Ingresa un valor para username: ")
	username := readLine()

	fmt.Print("Ingresa un valor para email: ")
	email := readLine()

	fmt.Print("Ingresa un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil{
		panic("No es posible convertir de un string a un entero")
	}

	id++
	user := User {id, username, email, age}
	users[id] = user

	fmt.Println(">>> Usuario creado exitosamente!\n")
}

func readUser(){
	clearConsole()

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

	fmt.Println("\n")
}

func updateUser(){
	clearConsole()

	fmt.Print("Ingresa el id del usuario a editar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil{
		panic("No es posible convertir de un string a un entero")
	}

	if _, ok := users[id]; ok{
		fmt.Print("Ingresa un valor para username: ")
		username := readLine()

		fmt.Print("Ingresa un valor para email: ")
		email := readLine()

		fmt.Print("Ingresa un valor para edad: ")
		age, err := strconv.Atoi(readLine())

		if err != nil{
			panic("No es posible convertir de un string a un entero")
		}
	
		user := User {id, username, email, age}
		users[id] = user
	}

	fmt.Println(">>> Usuario actualizado exitosamente!\n")
}

func deleteUser(){
	clearConsole()

	fmt.Print("Ingresa el id del usuario a eliminar: ")
	id, err := strconv.Atoi(readLine())

	if err != nil{
		panic("No es posible convertir de un string a un entero")
	}

	if _, ok := users[id]; ok{
		delete(users, id)
	}

	fmt.Println(">>> Usuario eliminado exitosamente!\n")
}

func clearConsole(){
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func readLine() string{
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor!")
	}else{
		return strings.TrimSuffix(option, "\n")
	}
}

func main(){

	var option string

	users = make(map[int]User)

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A) Crear")
		fmt.Println("B) Listar")
		fmt.Println("C) Actualizar")
		fmt.Println("D) Eliminar")
	
		fmt.Print("Ingresa una opcion valida: ")

		option = readLine()
	
		if option == "Quit" || option == "Q"{
			break
		}

		switch option {
			case "A", "Crear":
				createUser()
			case "B", "Listar":
				readUser()
			case "C", "Actualizar":
				updateUser()
			case "D", "Eliminar":
				deleteUser()
			default:
				fmt.Println("Opción no valida.")
		}
	}

	fmt.Println("Adios!")
}