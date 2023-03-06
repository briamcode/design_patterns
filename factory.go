/*
Factory
Es un patrón creacional, que nos permite crear una “fabrica” de objetos a partir
de una clase base y a su vez va implementar comportamientos polimórficos que
permite modificar el comportamiento de las clases heredadas
*/

package main

import "fmt"

// Definimos la interfaz IProduct que describe el comportamiento
// de cualquier producto que se venderá en nuestra tienda de computadoras
type IProduct interface {
	setStock(stock int)  // Establece la cantidad de stock de un producto
	getStock() int       // Devuelve la cantidad de stock actual de un producto
	setName(name string) // Establece el nombre de un producto
	getName() string     // Devuelve el nombre de un producto
}

// Definimos la estructura Computer para representar a una computadora
type Computer struct {
	name  string // El nombre de la computadora
	stock int    // La cantidad de stock de la computadora
}

// Implementamos los métodos de la interfaz IProduct para la estructura Computer
func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

// Definimos la estructura Laptop que es una computadora portátil
type Laptop struct {
	Computer // Usamos la composición para incluir los campos de Computer en Laptop
}

// La función newLaptop devuelve un nuevo objeto Laptop con valores predeterminados
func newLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// Definimos la estructura Desktop que es una computadora de escritorio
type Desktop struct {
	Computer // Usamos la composición para incluir los campos de Computer en Desktop
}

// La función newDesktop devuelve un nuevo objeto Desktop con valores predeterminados
func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

// La función GetComputerFactory es una función de fábrica que devuelve un objeto que implementa la interfaz IProduct
func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLaptop(), nil // Devuelve una nueva Laptop
	}

	if computerType == "desktop" {
		return newDesktop(), nil // Devuelve una nueva Desktop
	}

	return nil, fmt.Errorf("invalid computer type") // Devuelve un error si el tipo de computadora es inválido
}

// La función printNameAndStock imprime el nombre y la cantidad de stock de un objeto IProduct
func printNameAndStock(p IProduct) {
	fmt.Printf("Product name : %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("laptop")   // Obtenemos una nueva Laptop
	desktop, _ := GetComputerFactory("desktop") // Obtenemos una nueva Desktop

	printNameAndStock(laptop)  // Imprimimos el nombre y la cantidad de stock de la Laptop
	printNameAndStock(desktop) // Imprimimos el nombre y la cantidad de stock de la Desktop
}
