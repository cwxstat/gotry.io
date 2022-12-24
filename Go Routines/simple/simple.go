package simple

import "time"

func SumOfSum(count int) <-chan int {
	c := make(chan int)
	sum := 0
	go func() {
		for i := 0; i < count; i++ {
			sum += i
			c <- sum
			time.Sleep(1 * time.Second)
		}
		close(c)
	}()
	return c
}

func DoPassInFunc(count int, f func(int) <-chan int) <-chan int {
	c := make(chan int, count)
	done := make(chan bool)

	for i := 0; i < count; i++ {
		go func(c chan int, i int) {
			for j := range f(i) {
				c <- j
			}
			done <- true
		}(c, i)

	}
	go func() {
		for i := 0; i < count; i++ {
			<-done
		}
		close(c)
	}()
	return c
}
