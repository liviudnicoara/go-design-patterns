package main

import "fmt"

// The Adapter Design Pattern is a structural design pattern that allows two incompatible interfaces to work together.
// This incompatibility could be because of differences in method names, argument types, or other differences
// in the interfaces.
// The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to
// format and interface recognizable by the second object.

type SlowEmailProvider struct {
}

func (p *SlowEmailProvider) SendEmail(from string, to string) {
	fmt.Printf("Email sent from %s to %s\n", from, to)
}

type EmailProvider interface {
	Send(Email)
}

type Email struct {
	From string
	To   string
}

type FastEmailProvider struct {
}

func (p *FastEmailProvider) Send(email Email) {
	fmt.Println("Email sent: ", email)
}

type SlowEmailAdaptar struct {
	slowEmailProvider *SlowEmailProvider
}

func (p *SlowEmailAdaptar) Send(email Email) {
	p.slowEmailProvider.SendEmail(email.From, email.To)
}

func sendSubscriptionEmail(email Email, provider EmailProvider) {
	provider.Send(email)
}

func main() {
	email := Email{
		From: "sender",
		To:   "receiver",
	}

	fast := &FastEmailProvider{}
	sendSubscriptionEmail(email, fast)

	adapter := &SlowEmailAdaptar{slowEmailProvider: &SlowEmailProvider{}}
	sendSubscriptionEmail(email, adapter)
}
