func main() {
  f1()
  f2(10, 1.5)
  fmt.Println(f3(10))
  fmt.Println(f4(10, 6))
}

// No args
func f1() {
  fmt.Println("F1 got called!")
}

// With args...
func f2(i int, f float32) {
  fmt.Printf("%d -- %1.1f\n", i, f)
}

// With single return
func f3(i int) string {
  return fmt.Sprintf("Temperature today is %d°F", i)
}

// Multi args, Multi returns (NOTE: Repeating args of same value)
func f4(i, j int) (temp, down string) {
  return fmt.Sprintf("Temperature is %d°F.", i), fmt.Sprintf("Sun down at %dpm.", j)
}
