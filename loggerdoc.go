package main

import (
	"log"
	"os"
)

func main() {
	// Standard logger
	log.Printf("You rock!")

	// Custom logger
	logger := log.New(os.Stdout, ">> MyStuff-", log.Ldate|log.Ltime|log.Llongfile|log.LUTC)

	// Use it!
	logger.Println("YO!")
	logger.Printf("%s -- %2.2f", "Blee", 10.5678)

	// Calls os.Exit(1)
	logger.Fatalf("And we're %s!", "dead")

	// Calls os.Panic()
	logger.Panicln("Ruroh... Oh Dear!")
}
