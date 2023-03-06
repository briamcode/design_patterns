/*
Singleton
Es un patrón creacional, nos permite manejar y restringir una sola instancia
de una clase. El caso de uso mas común es para crear conexiones de bases de
datos y así evitar la creación de varias conexiones a la base de datos.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Database representa una conexión de base de datos.
type Database struct {
}

// CreateSingleConnection crea una única conexión de base de datos.
func (Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton for Database")
	time.Sleep(2 * time.Second)
}

var db *Database
var lock sync.Mutex

// getDatabaseInstance devuelve una instancia de base de datos única.
func getDatabaseInstance() *Database {
	lock.Lock()         // Adquiere el candado.
	defer lock.Unlock() // Libera el candado al salir de la función.
	if db == nil {      // Si la instancia de base de datos no existe.
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection() // Crea una nueva conexión de base de datos.

	} else {
		fmt.Println("DB Already Created")
	}
	return db // Devuelve la instancia de base de datos.
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			getDatabaseInstance() // Obtiene la instancia de base de datos.
		}()
	}
	wg.Wait()
}
