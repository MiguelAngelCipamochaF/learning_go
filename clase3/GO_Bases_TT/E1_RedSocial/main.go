package main

import "fmt"

type Usuario struct {
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Edad       string
	Correo     string
	Contrasena string
}

func main() {

	u1 := &Usuario{}
	u1.SetNombre("Miguel")
	u1.SetEdad("19")
	u1.Setcorreo("correo1@gmail.com")
	u1.SetContrasena("contrasena")

	fmt.Printf("Usuarios: \n")
	fmt.Printf("Nombre: %v\n", u1.Nombre)
	fmt.Printf("Apellido: %v\n", u1.Apellido)
	fmt.Printf("Edad: %v\n", u1.Edad)
	fmt.Printf("Correo: %v\n", u1.Correo)
}

func (u *Usuario) SetNombre(nombre string) {
	u.Nombre = nombre
}

func (u *Usuario) SetEdad(edad string) {
	u.Edad = edad
}

func (u *Usuario) Setcorreo(correo string) {
	u.Correo = correo
}

func (u *Usuario) SetContrasena(contrasena string) {
	u.Contrasena = contrasena
}
