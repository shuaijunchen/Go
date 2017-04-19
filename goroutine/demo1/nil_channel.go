// nil_channel.go
package main

func main() {
	ch := make(chan bool)
	ch <- true
}
