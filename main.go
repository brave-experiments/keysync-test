package main

import (
	"log"

	"github.com/brave/nitriding"
)

// TODO: Set to 'false' on the other enclave.
const originEnclave = true

type keyMaterial struct {
	Key string
}

func main() {
	enclave := nitriding.NewEnclave(
		&nitriding.Config{
			SOCKSProxy: "socks5://127.0.0.1:1080",
			FQDN:       "TODO: EC2-Host1",
			Port:       8443,
			UseACME:    true,
		},
	)
	go func() {
		if err := enclave.Start(); err != nil {
			log.Fatalf("Enclave terminated: %v", err)
		}
	}()

	k := &keyMaterial{}

	// We're the origin enclave.  Configure the key material.
	if originEnclave {
		k.Key = "foobar"
		enclave.SetKeyMaterial(k)
		log.Println("Set key material.  Waiting for other enclave to request it.")
	} else {
		if err := nitriding.RequestKeys("TODO: EC2-Host2", k); err != nil {
			log.Fatalf("Failed to request keys from remote enclave: %v", err)
		}
		log.Printf("Successfully retrieved keys from remote enclave: %s", k.Key)
	}

	// Wait.
	c := make(chan bool)
	<-c
}
