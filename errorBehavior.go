package main

import "fmt"

type (
	Retryable interface {
		Retry() bool
	}
	ErrAuth struct {
		reason string
	}
)

func (r ErrAuth) Error() string { return fmt.Sprintf(r.reason) }
func (r ErrAuth) Retry() bool   { return true }

func isRetryable(e error) (b bool) {
	if err, ok := e.(Retryable); ok {
		b = err.Retry()
	}
	return
}

func main() {
	for i, j := 1, 0; isRetryable(fred(i)); j++ {
		if j == 5 {
			fmt.Println("Giving up! Done retrying...")
			break
		}
		i++
	}
	fmt.Println("Good bye!")
}

func Bee(i int) error {
	if i < 5 {
		fmt.Printf("Bee (%d) auth Failed!\n", i)
		return ErrAuth{reason: "Invalid auth!! Your password is `secret, you dope!"}
	}
	fmt.Println("Bee Succeeded!")
	return nil
}
