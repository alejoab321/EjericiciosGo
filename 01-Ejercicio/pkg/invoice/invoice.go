package invoice;
import (
	"github.com/alejoab321/01-Ejercicio/pkg/customer"
	"github.com/alejoab321/01-Ejercicio/pkg/invoiceitem"
)

type Invoice struct {
	country string
	city string
	total float64
	client customer.Customer
	items invoiceitem.Items 
}

func New(country,city string,client customer.Customer, items invoiceitem.Items) Invoice{
	return Invoice{
		country: country,
		city: city,
		client: client,
		items: items,
	}
}

func (i *Invoice) SetTotal(){
	i.total = i.items.Total()
}
