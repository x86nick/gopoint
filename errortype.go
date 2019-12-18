package main

import (
	"fmt"
	"log"
)

func main() { // on calling side, check with error type is 
	switch e := DoIt(10).(type) {
	case UserError:
		fmt.Println("UserError", e)
	case NetworkError:
		log.Println("NetworkError", e)
	default:
		fmt.Println("All cool!")
	}
}

// ----------------------------------------------------------------------------
// Let's assume this lives in an imported package
// Package Blah

type (
	UserError struct {
		user, reason string
	}
	NetworkError struct {
		code   int
		reason string
	}
)

func (e UserError) Error() string { // this is contract to create new error //
	return fmt.Sprintf("[%s] - %s", e.user, e.reason)
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("(%d): %s", e.code, e.reason)
}

func DoIt(i int) error { // exposed function 
	switch i {
	case 10:
		return UserError{user: "Blah", reason: "Access denied!"} // expose error
	case 20:
		return NetworkError{code: 1234, reason: "Invalid protocol"}
	}
	return nil
}
