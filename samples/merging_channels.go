package samples

import (
	"fmt"
	"time"
)

var iterationsNumber int = 10

// cat sound will always be received before dog sound
func Run2SyncWay() {
	catChannel := speakCatLanguage()
	dogChannel := speakDogLanguage()

	for i := 0; i < iterationsNumber; i++ {
		fmt.Println(<-catChannel)
		fmt.Println(<-dogChannel)
	}
}

func Run2AsyncWayWithChannelsMerging() {
	channel := make(chan string)

	go func() {
		for {
			channel <- <-speakCatLanguage()
		}
	}()

	go func() {
		for {
			channel <- <-speakDogLanguage()
		}
	}()

	for i := 0; i < 2*iterationsNumber; i++ {
		fmt.Println(<-channel)
	}
}

// creates a channel that receives cat sounds
func speakCatLanguage() <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; i < iterationsNumber; i++ {
			channel <- fmt.Sprintf("Miau %v", i)
			time.Sleep(time.Second)
		}
	}()

	return channel
}

func speakDogLanguage() <-chan string {
	channel := make(chan string)

	go func() {
		for i := 0; i < iterationsNumber; i++ {
			channel <- fmt.Sprintf("Hau %v", i)
			time.Sleep(time.Second)
		}
	}()

	return channel
}
