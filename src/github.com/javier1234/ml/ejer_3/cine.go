package main

import (
	"os"
	"strings"
	"fmt"
)

/**
Precio base de los tickets (valor
 */
const Base_price  = float64(1)



type ticket interface {
	price() float64
}


type Normal struct {
}
type Jubilado struct {
}
type Invitado struct {
}


func (t *Normal) price() float64 {
	fmt.Println("precio normal")
	return Base_price
}
func (t *Jubilado) price() float64 {
	fmt.Println("precio jubilado")
	return Base_price / 2
}
func (t *Invitado) price() float64 {
	fmt.Println("precio invitado")
	return 0
}


/**
Modelar la funcionalidad para un sistema de cines que calcule los ingresos netos de una
función en base a los asistentes y al precio base de la entrada. Existen tres tipos de
asistentes y tienen las siguientes características:
Normales: Pagan el 100%
Jubilados: Tienen un 50% de descuento
Invitados: No pagan nada
 */
func main() {
	var ticket_normal = Normal{}
	var ticket_jubilado  = Jubilado{}
	var ticket_invitado  = Invitado{}

	people := strings.Split(os.Args[1], "")
	tickets := make([]ticket, len(people))

	//vendiendo
	for i, person := range people {
		fmt.Println("person:", person)
		// vendiendo
		switch person {
		case "n":
			tickets[i] = &ticket_normal
		case "j":
			tickets[i] = &ticket_jubilado
		default:
			tickets[i] = &ticket_invitado
		}
	}

	// juntando la plata
	money := float64(0)
	for _,t := range tickets {
		money += t.price()
	}
	fmt.Print("recaudamos = ", money)
}