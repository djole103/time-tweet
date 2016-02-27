package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	userName = flag.String("user-name", "winkface", "Username login")
	password = flag.String("password", "idk", "Password")
)

func SetFromEnv(flagName string, envName string) {
	if env := os.Getenv(envName); len(env) > 0 {
		flag := flag.Lookup(flagName)
		flag.Value.Set(env)
	}
}
