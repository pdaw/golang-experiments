package samples

import (
	"fmt"
	"time"
)

func Run3() {
	var channelForCats = make(chan string)
	var channelForDogs = make(chan string)

	go func() {
		// let's make cat sounds!
		for i := 0; i < 5; i++ {
			channelForCats <- fmt.Sprintf("miau %v", i)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		// let's make dog sound
		for i := 0; i < 5; i++ {
			channelForDogs <- fmt.Sprintf("hau %v", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		// with select statement we can wait on many channels - non-blocking, async experience
		select {
		case sound1 := <-channelForCats:
			fmt.Println(sound1)
		case sound2 := <-channelForDogs:
			fmt.Println(sound2)
		case <-time.After(3 * time.Second):
			fmt.Println("Well, it looks like no need value will be fetched from channels")
			return
		}
	}
}
