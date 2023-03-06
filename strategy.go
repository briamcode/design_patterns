/* Strategy
es un patrón de diseño de comportamiento que te permite definir una familia de algoritmos, colocar cada uno de ellos en una clase separada y hacer sus
objetos intercambiables.

*/

// Paquete main proporciona una aplicación para proteger contraseñas utilizando diferentes algoritmos hash.
package main

import "fmt"

// PasswordProtector es una estructura que almacena el nombre de usuario, el nombre de la contraseña y el algoritmo hash utilizado.
type PasswordProtector struct {
	user          string
	passwordName  string
	hashAlgorithm HashAlgorithm
}

// HashAlgorithm es una interfaz que define el método Hash.
type HashAlgorithm interface {
	Hash(p *PasswordProtector)
}

// NewPasswordProtector crea un nuevo objeto PasswordProtector y lo devuelve.
func NewPasswordProtector(user string, passwordName string, hash HashAlgorithm) *PasswordProtector {
	return &PasswordProtector{
		user:          user,
		passwordName:  passwordName,
		hashAlgorithm: hash,
	}
}

// SetHashAlgorithm establece el algoritmo hash utilizado por el objeto PasswordProtector.
func (p *PasswordProtector) SetHashAlgorithm(hash HashAlgorithm) {
	p.hashAlgorithm = hash
}

// Hash llama al método Hash en el objeto algoritmo hash utilizado por el objeto PasswordProtector.
func (p *PasswordProtector) Hash() {
	p.hashAlgorithm.Hash(p)
}

// SHA es una estructura que implementa el método HashAlgorithm para el algoritmo SHA.
type SHA struct {
}

// Hash implementa el método HashAlgorithm para el algoritmo SHA.
func (SHA) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using SHA for %s\n", p.passwordName)
}

// MD5 es una estructura que implementa el método HashAlgorithm para el algoritmo MD5.
type MD5 struct {
}

// Hash implementa el método HashAlgorithm para el algoritmo MD5.
func (MD5) Hash(p *PasswordProtector) {
	fmt.Printf("Hashing using MD5 for %s\n", p.passwordName)
}

// main es la función principal que crea un objeto PasswordProtector y utiliza diferentes algoritmos hash para proteger una contraseña.
func main() {
	sha := &SHA{}
	md5 := &MD5{}

	passwordProtector := NewPasswordProtector("nestor", "gmail password", sha)
	passwordProtector.Hash()

	passwordProtector.SetHashAlgorithm(md5)
	passwordProtector.Hash()
}

/*
¡Por supuesto! El patrón de diseño Strategy es un patrón de comportamiento que
permite a un objeto cambiar su comportamiento en tiempo de ejecución mediante
la selección de un algoritmo específico en una familia de algoritmos.

En el patrón de diseño Strategy, hay tres componentes principales: el contexto,
la estrategia y las diferentes estrategias concretas. El contexto es un objeto
que se comporta de manera diferente según el algoritmo que se utiliza.
La estrategia es una interfaz que define un conjunto de métodos que el contexto
puede utilizar. Por último, las diferentes estrategias concretas son objetos que
implementan la interfaz de estrategia y proporcionan diferentes algoritmos.

En resumen, el patrón de diseño Strategy permite a los objetos cambiar
su comportamiento dinámicamente al delegar el trabajo a diferentes algoritmos.
Esto hace que el código sea más flexible y fácil de mantener, ya que los
algoritmos pueden cambiarse sin afectar al objeto principal.

Veamos un ejemplo concreto para entender mejor cómo funciona el patrón de
diseño Strategy. Imagina que tienes una tienda en línea que vende diferentes
productos y quieres permitir a tus clientes elegir diferentes formas de pago.
Para ello, puedes utilizar el patrón de diseño Strategy:

    El contexto es la tienda en línea.
    La estrategia es una interfaz que define el método pay().
    Las diferentes estrategias concretas son objetos que implementan
	la interfaz pay(), como PayPal, tarjeta de crédito, transferencia bancaria,
	etc.

De esta manera, cuando un cliente realiza una compra, la tienda en línea
puede cambiar dinámicamente el método de pago que se utiliza según la preferencia
del cliente.

Espero que esto te haya ayudado a comprender mejor el patrón de diseño
Strategy. Si tienes alguna pregunta adicional, no dudes en preguntar.


*/
