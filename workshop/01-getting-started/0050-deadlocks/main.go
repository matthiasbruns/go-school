package main

type Recipe struct {
	Name    string
	Minutes int
}

// simulate a deadlock
func main() {
	c := make(chan Recipe)

	for {
		select {
		case r := <-c:
			println(r.Name)
		}
	}

	go func() {
		for i := 0; i < 10; i++ {
			c <- Recipe{
				Name:    "Spaghetti Carbonara",
				Minutes: 20,
			}
		}
	}()
}
