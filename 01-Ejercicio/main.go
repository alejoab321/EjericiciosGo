package main

import (
	"fmt"
	"github.com/alejoab321/01-Ejercicio/pkg/customer"
	"github.com/alejoab321/01-Ejercicio/pkg/invoice"
	"github.com/alejoab321/01-Ejercicio/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Popayán",
		customer.New("Alejandro","Cl 123 #23-04","12345"),
		invoiceitem.NewItems(
			invoiceitem.New(1,"Curso de Go",12.36),
			invoiceitem.New(2,"Curso de GoPoo",12.36),
			invoiceitem.New(3,"Curso de Go BD",13.26),
		),
	)

	i.SetTotal()
	fmt.Printf("%v", i)
}