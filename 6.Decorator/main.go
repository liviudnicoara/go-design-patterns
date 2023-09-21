package main

import "fmt"

// The Decorator design pattern allows us to add new functionalities to existing objects without altering their structures.
// This pattern involves a set of decorator classes that are used to wrap concrete components.
// Decorator classes mirror the type of the components they decorate but add or override behavior.
// The key feature of this pattern is the ability to "wrap" or "decorate" an object multiple times,
// with different decorators, adding cumulative behavior to the original object.
// The pattern is particularly useful for adding responsibilities to objects dynamically and optionally.

type Sender interface {
	Send(string)
}

type EmailSender struct {
}

func (s EmailSender) Send(msg string) {
	fmt.Println("Sent email: ", msg)
}

type CompressedEmailSender struct {
	sender Sender
}

func (c CompressedEmailSender) Send(msg string) {
	msg = "Compressed " + msg
	c.sender.Send(msg)
}

type EncryptedEmailSender struct {
	sender Sender
}

func (e EncryptedEmailSender) Send(msg string) {
	msg = "Encrypted " + msg
	e.sender.Send(msg)
}

func main() {
	emailSender := EmailSender{}
	compressedEmailSender := CompressedEmailSender{emailSender}
	encryptedEmailSender := EncryptedEmailSender{emailSender}
	encryptedAndCompressedEmailSender := EncryptedEmailSender{compressedEmailSender}

	msg := "MESSAGE"

	emailSender.Send(msg)
	compressedEmailSender.Send(msg)
	encryptedEmailSender.Send(msg)
	encryptedAndCompressedEmailSender.Send(msg)
}
