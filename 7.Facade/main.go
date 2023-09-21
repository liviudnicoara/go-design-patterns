package main

import "fmt"

// Facade defines a higher-level interface that makes the subsystem easier to use. Essentially, it simplifies a
// complex system by providing a single entry point to coordinate various functionalities.
// The facade pattern can often be found in complex systems and frameworks to provide a simpler,
// more developer-friendly API. One common area in software engineering where the façade pattern is used is in
// the design of software libraries and frameworks.

type Connection struct{}

func (c *Connection) Open() {
	fmt.Println("Opening connection...")
}

func (c *Connection) GetToken() {
	fmt.Println("Getting connection token...")
}

type Session struct{}

func (s *Session) Init() {
	fmt.Println("Initializing session...")
}

// EmailProvider provides a simplified interface to interact with the email server
type EmailProvider struct {
	connection *Connection
	session    *Session
}

func (p *EmailProvider) Start() {
	if p.connection == nil {
		p.connection = &Connection{}
	}

	p.connection.Open()
	p.connection.GetToken()

	if p.session == nil {
		p.session = &Session{}
	}

	p.session.Init()
	fmt.Println("Email provider started!")
}

func main() {
	provider := EmailProvider{}
	provider.Start()
}
