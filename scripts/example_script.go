package main

import "fmt"

// Prereq acts as a filter - it allows scripts to abort early if certain system
// requirements aren't met (for example, if the wrong operating system is being
// used), it is optional
func Prereq() bool {
    return true
}

// Name is a human-readable identifier of the script (it is NOT a unique ID, and
// should not be used as so)
const Name string = "Example test"

// Description is a short description of what the script tests for
const Description string = "Example test"

// Categories is a list of tags that can help a user piece together what a 
// given script checks for
var Categories = [...]string{"example_tag"}

// Check is where the actual checking logic occurs, if it returns true then a
// fault has been found and will be flagged to the user
func Check() bool {
    return true
}

// FixDescription returns a description to the user on how to manually fix the
// fault. It is a function in order to allow instructions to be tailored to
// the user's system. It should also include extra information on what aspect
// of the fault compromises security and insight into the severity of the fault
func FixDescription() string {
    return "Do xyz to fix it, alternatively do this, watch out for abc."
}

// FixAuto is a function that attempts to fix the fault automatically, it is
// optional
func FixAuto() {
    fmt.Printf("\n\n\n\n\n\n")
}
