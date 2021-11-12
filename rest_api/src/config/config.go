package config

import "os"

var Config map[string]string

func init() {
	Config = map[string]string{
		"PORT":              os.Getenv("PORT"),
		"MESSAGING_ADDRESS": os.Getenv("MESSAGING_ADDRESS"),
	}
}
