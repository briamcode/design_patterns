/*
Es un patrón de diseño estructural, es un intermedio que permite adaptar el
comportamiento de un struct a una interfaz, cuando existe una incompatibilidad
sin la necesidad de tener que refactorizar código.

*/

package main

import "fmt"

// Payment es una interfaz que define el método Pay.
type Payment interface {
	Pay()
}

// CashPayment es una implementación de Payment para pagos en efectivo.
type CashPayment struct {
}

// Pay implementa el método Pay de la interfaz Payment para pagos en efectivo.
func (CashPayment) Pay() {
	fmt.Println("Payment Using Cash")
}

// ProcessPayment toma una implementación de Payment y llama al método Pay en ella.
func ProcessPayment(p Payment) {
	p.Pay()
}

// BankPayment es una implementación de Payment para pagos bancarios.
type BankPayment struct {
}

// BankPaymentAdapter es una adaptación para BankPayment que toma un número de cuenta bancaria.
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

// Pay implementa el método Pay de la interfaz Payment para pagos bancarios.
func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

// Pay toma un número de cuenta bancaria e imprime un mensaje de pago usando la cuenta especificada.
func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using Bankaccount %d\n", bankAccount)
}

func main() {
	// Crea una instancia de CashPayment y la pasa a ProcessPayment.
	cash := &CashPayment{}
	ProcessPayment(cash)

	// Crea una instancia de BankPaymentAdapter y la pasa a ProcessPayment.
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
