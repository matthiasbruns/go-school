## Exercise

- Create a program that 
- Spawns 10 goroutines 
- Each goroutine should create a random amount of Recipes and send it to a channel 
  - You can use `rand.Intn(n)` for random int generation
- Add a loop that collects the values in the channel and prints them in the order they were received with fmt.Print
