/*
Observer
es un patrón de diseño de comportamiento que te permite definir
un mecanismo de suscripción para notificar a varios objetos sobre
cualquier evento que le suceda al objeto que están observando.


*/

package main

import "fmt"

// Topic es una interfaz que define un tema observable.
type Topic interface {
	register(observer Observer)
	broadcast()
}

// Observer es una interfaz que define un observador de temas.
type Observer interface {
	getId() string
	updateValue(string)
}

// Item es una estructura que define un elemento observado.
type Item struct {
	observers []Observer
	name      string
	available bool
}

// NewItem es una función que devuelve una instancia de Item.
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// updateAvailable es una función que se encarga de actualizar el estado de disponibilidad del Item
// y notificar a los observadores registrados.
func (i *Item) updateAvailable() {
	fmt.Printf("El artículo %s está disponible\n", i.name)
	i.available = true
	i.broadcast()
}

// register es una función que se encarga de agregar un nuevo observador a la lista de observadores
// del Item.
func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

// broadcast es una función que se encarga de notificar a todos los observadores registrados
// sobre un cambio en el estado del Item.
func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// EmailClient es una estructura que define un cliente de correo electrónico.
type EmailClient struct {
	id string
}

// updateValue es una función que se encarga de actualizar el valor de un observador.
func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Enviando correo electrónico - %s está disponible para el cliente %s\n", value, eC.id)
}

// getId es una función que devuelve el ID de un cliente de correo electrónico.
func (eC *EmailClient) getId() string {
	return eC.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}

	secondObserver := &EmailClient{
		id: "34dc",
	}

	// Registra los observadores para el elemento y notifícalos de que está disponible.
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.updateAvailable()
}
