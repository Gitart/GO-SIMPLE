type IntContainer []int

func (i IntContainer) Iterator(cancel <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		for _, val := range i {
			select {
			case ch <- val:
			case <-cancel:
				close(ch)
				return
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	c := IntContainer([]int{1, 2, 3, 4, 5})
	cancel := make(chan struct{})
	for x := range c.Iterator(cancel) {
		println(x)
		break
	}
	close(cancel)
}
