package main

import (
	"fmt"
	"github.com/LeakIX/go-emaildecoder"
	"io"
	"os"
	"path/filepath"
)

func main() {
	// Creates an attachments directory
	if err := os.MkdirAll("attachments", 0750); err != nil && err != os.ErrExist {
		panic(err)
	}
	// Creates a new decoder from stdin, with an attachment handler
	decoder := emaildecoder.NewDecoder(os.Stdin, handleAttachment)
	// Parse the email
	email, err := decoder.Decode()
	if err != nil {
		panic(err)
	}
	// Display results
	fmt.Printf("From: %s\nSubject: %s\n", email.Headers.Get("From"), email.Headers.Get("Subject"))
	fmt.Println(string(email.PlainText))
}

// Handle attachments
func handleAttachment(attachment emaildecoder.Attachment) {
	fmt.Printf("Saving %s...\n", attachment.Filename)
	if file, err := os.OpenFile(filepath.Join("attachments", attachment.Filename), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0640); err != nil {
		panic(err)
	} else if _, err = io.Copy(file, attachment.Reader); err != nil {
		panic(err)
	} else if err = file.Close(); err != nil {
		panic(err)
	}
}
