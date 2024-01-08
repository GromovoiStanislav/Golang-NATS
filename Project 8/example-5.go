package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Close()

	// "*" matches any token, at any level of the subject.
	nc.Subscribe("foo.*.baz", func(m *nats.Msg) {
		fmt.Printf("Msg received on [%s] : %s\n", m.Subject, string(m.Data));
	})

	nc.Subscribe("foo.bar.*", func(m *nats.Msg) {
		fmt.Printf("Msg received on [%s] : %s\n", m.Subject, string(m.Data));
	})

	// ">" matches any length of the tail of a subject, and can only be the last token
	// E.g. 'foo.>' will match 'foo.bar', 'foo.bar.baz', 'foo.foo.bar.bax.22'
	nc.Subscribe("foo.>", func(m *nats.Msg) {
		fmt.Printf("Msg received on [%s] : %s\n", m.Subject, string(m.Data));
	})

	// Matches all of the above
	nc.Publish("foo.bar.baz", []byte("Hello World"))

	select {}
}

// Output:
// Msg received on [foo.bar.baz] : Hello World
// Msg received on [foo.bar.baz] : Hello World
// Msg received on [foo.bar.baz] : Hello World