//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer szenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, tweetCh chan<- Tweet, wg *sync.WaitGroup) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(tweetCh)
			wg.Done()
			return
		}

		tweetCh <- *tweet
	}
}

func consumer(tweetCh <-chan Tweet, wg *sync.WaitGroup) {
	for t := range tweetCh {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	wg.Done()
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	wg := &sync.WaitGroup{}
	tweetCh := make(chan Tweet)
	wg.Add(2)
	// Producer
	go producer(stream, tweetCh, wg)
	// Consumer
	go consumer(tweetCh, wg)
	wg.Wait()

	fmt.Printf("Process took %s\n", time.Since(start))
}
