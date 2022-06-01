//START_1 OMIT
ch <- v   // Send v to channel ch
v := <-ch // Receive from ch, and assign value to v
//END_1 OMIT

//START_2 OMIT
ch := make(chan int)
//END_2 OMIT

//START_3 OMIT
ch := make(chan int, 100)
//END_3 OMIT

//START_4 OMIT
v, ok := <-ch
//END_4 OMIT

//START_5 OMIT
for i := range c
//END_5 OMIT