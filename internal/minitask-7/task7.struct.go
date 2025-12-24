package minitask7

import "fmt"

type Person struct {
	Name string
	Address string
	Phone string
}

func NewPerson(name, address, phone string) *Person  {
	return &Person{
		Name: name,
		Address: address,
		Phone: phone,
	}
}

func (p *Person) Print() string  {
	return fmt.Sprintf("Name: %s\nAddress: %s\nPhone: %s\n", p.Name, p.Address, p.Phone)
}

func (p *Person) Greet() string  {
	return fmt.Sprintf("Hello, %s!", p.Name)
}

func (p *Person) UpdatePersonName(name string)  {
	p.Name = name
}