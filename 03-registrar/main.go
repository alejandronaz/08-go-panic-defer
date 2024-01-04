package main

import (
	"errors"
	"fmt"
)

type Customer struct {
	Name        string
	ID          string
	PhoneNumber string
	Home        string
}

type CustomersList []Customer

func (c *CustomersList) AddCustomer(customer Customer) (err error) {

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("error: client already exists")
		}
	}()

	// check if customer exists
	for _, c := range *c {
		if c.ID == customer.ID {
			panic("Error: client already exists")
		}
	}

	errVal := checkCustomer(customer)
	if errVal != nil {
		return errors.New("error: client has zero values")
	}

	*c = append(*c, customer)
	return nil
}

func checkCustomer(customer Customer) (err error) {

	if customer.ID == "" || customer.Home == "" || customer.Name == "" || customer.PhoneNumber == "" {
		return errors.New("error: client already exists")
	}

	return nil
}

func main() {

	defer func() {
		fmt.Println("End of execution")
		fmt.Println("Several errors were detected at runtime")
	}()

	list := CustomersList{} // []Customer{}

	c1 := Customer{
		Name:        "Juan",
		ID:          "123",
		PhoneNumber: "123456789",
		Home:        "Calle 123",
	}

	c2 := Customer{
		Name:        "Pedro",
		ID:          "123456",
		PhoneNumber: "123456789",
		Home:        "Calle 123",
	}

	err := list.AddCustomer(c1)
	if err != nil {
		fmt.Println(err)
	}

	err = list.AddCustomer(c2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(list)

	// add existing customer
	err = list.AddCustomer(c1)
	if err != nil {
		fmt.Println(err)
	}

	err = list.AddCustomer(Customer{})
	if err != nil {
		fmt.Println(err)
	}
}
