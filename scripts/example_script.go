package main

import "fmt"

func Prereq() bool {
    return true;
}

func Name() string {
    return "Example test"
}

func Description() string {
    return "Example test"
}

func Categories() []string {
    return []string{}
}

func FixDescription() string {
    return "Do xyz to fix it, alternatively do this, watch out for abc."
}

func FixAuto() {
    fmt.Printf("\n\n\n\n\n\n")
}
