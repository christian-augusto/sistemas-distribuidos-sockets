package models

type Client struct {
	Name    string
	Cpf     string
	Address string
}

func NewClient(name string, cpf string, address string) *Client {
	return &Client{
		Name:    name,
		Cpf:     cpf,
		Address: address,
	}
}
