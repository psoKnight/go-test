package main

func NilChannel() {
	var ch = make(chan int)

	defer func() {
		close(ch)
		ch = nil
	}()

}

func main() {
	NilChannel()
}
