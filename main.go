package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	delay = 1 * time.Second
)

func main() {
	ch := make(chan struct{})
	close(ch)
	chWork := ch
	chBackup := ch
	chControl := make(chan struct{})

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go worker(wg, &chWork, chControl)
	go control(&chWork, &chBackup, chControl)

	wg.Wait()
	fmt.Println("done")
}

func worker(wg *sync.WaitGroup, chWork *chan struct{}, chControl chan struct{}) {
	defer wg.Done()

	for {
		select {
		case <-*chWork:
			fmt.Println("worker work")
			time.Sleep(delay)
		case _, ok := <-chControl:
			if ok {
				continue
			}
			return
		}
	}
}

func control(chWork *chan struct{}, chBackup *chan struct{}, chControl chan struct{}) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		switch strings.Trim(text, " \n") {
		case "exit":
			*chWork = nil
			close(chControl)
			return
		case "play":
			*chWork = *chBackup
			chControl <- struct{}{}
		case "pause":
			*chWork = nil
			chControl <- struct{}{}
		default:
			fmt.Println(text)
		}
	}
}
