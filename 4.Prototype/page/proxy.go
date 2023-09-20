package page

import (
	"os"
)

type Page interface {
	Clone() Page
}

type PaymentPage struct {
	Provider    string
	Token       string
	Certificate os.File
}

func (p *PaymentPage) Clone() Page {
	return &PaymentPage{
		Provider:    p.Provider,
		Token:       p.Token,
		Certificate: p.Certificate,
	}
}
