package main

import (
	"fmt"
	"math/rand"
	"time"
)

var candidates = []string{"Ваня", "Рустам", "Куаныш", "Нурислам"}

func main() {

	zhambylCh := make(chan map[string]int)
	almatyCh := make(chan map[string]int)
	akmolaCh := make(chan map[string]int)

	go getZhambylOblVotes(zhambylCh)
	go getAlmatyOblVotes(almatyCh)
	go getAkmolaOblVotes(akmolaCh)

	results := calcVotes(zhambylCh, almatyCh, akmolaCh)

	for name, numOfVotes := range results {
		fmt.Printf("Кандидат [%s] набрал [%d] голосов\n", name, numOfVotes)
	}
}

func calcVotes(zhambylCh <-chan map[string]int, almatyCh <-chan map[string]int, akmolaCh <-chan map[string]int) map[string]int {
	ZhambylVotes := <-zhambylCh
	AlmatyVotes := <-almatyCh
	AkmolaVotes := <-akmolaCh

	results := make(map[string]int)

	for _, name := range candidates {
		results[name] = ZhambylVotes[name] + AlmatyVotes[name] + AkmolaVotes[name]
	}
	return results
}

func getZhambylOblVotes(ch chan<- map[string]int) {
	fmt.Println("Идет подсчет голосов в Жамбылской области...")

	votes := make(map[string]int)
	maxNumOfVotes := 100

	for _, name := range candidates {
		votes[name] = rand.Intn(maxNumOfVotes)
	}

	time.Sleep(1 * time.Second)

	ch <- votes

	fmt.Println("Подсчет голосов в Жамбылской области завершен")
}

func getAlmatyOblVotes(ch chan<- map[string]int) {
	fmt.Println("Идет подсчет голосов в Алматинской области...")

	votes := make(map[string]int)
	maxNumOfVotes := 200

	for _, name := range candidates {
		votes[name] = rand.Intn(maxNumOfVotes)
	}

	time.Sleep(2 * time.Second)

	ch <- votes

	fmt.Println("Подсчет голосов в Алматинской области завершен")
}

func getAkmolaOblVotes(ch chan<- map[string]int) {
	fmt.Println("Идет подсчет голосов в Акмолинской области...")

	votes := make(map[string]int)
	maxNumOfVotes := 300

	for _, name := range candidates {
		votes[name] = rand.Intn(maxNumOfVotes)
	}

	time.Sleep(3 * time.Second)

	ch <- votes

	fmt.Println("Подсчет голосов в Акмолинской области завершен")
}
