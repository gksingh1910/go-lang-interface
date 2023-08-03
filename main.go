package main

import (
	"fmt"
	"time"
)

type sports_token interface {
	createPlayerToken() string
	mintPlayerToken() (string, error)
	launchinOffering() string
	launchinSecondryMarket() string
}

type trade interface {
	buyPlayerToken() bool
	sellPalyerToken() bool
}

type player struct {
	token_id    int32
	name        string
	age         int32
	description string
	status      []string
}

func (p *player) createPlayerToken() string {
	p.status = append(p.status, "CREATE")
	return fmt.Sprintf("created with id as %v", (p.name + fmt.Sprintf("%v", p.age)))
}

func (p *player) mintPlayerToken() (string, error) {
	if p.status[0] != "CREATED" {
		return "", fmt.Errorf("token yet not created")
	}
	p.status = append(p.status, "MINTED")
	return fmt.Sprintf("Player with token id as %v has been minted", p.token_id), nil
}

func (p *player) launchinOffering() string {
	p.status = append(p.status, "INOFFERING")
	return fmt.Sprintf("player with token id as %v has been put under offering", p.token_id)
}

func (p *player) launchinSecondryMarket() string {
	p.status = append(p.status, "INMARKET")
	return fmt.Sprintf("player with token id as %v has been put in secondry market", p.token_id)
}

func completeLifecycle(sp sports_token) {
	i := sp.createPlayerToken()
	j, e := sp.mintPlayerToken()
	k := sp.launchinOffering()
	l := sp.launchinSecondryMarket()
	fmt.Printf("\n %v \n %v \n %v \n %v \n %v", i, j, k, l, e)
	// fmt.Printf("player struct is %v", *p)
}

func gettingCalled(id int) {
	fmt.Printf("\n getting called with id %v", id)
}

func main() {

	// Call1
	p := player{name: "dhoni", age: 32, description: "100s of 100"}
	// x := p.createPlayerToken()
	// fmt.Println("output of player token creation", x)
	// fmt.Println("status of player token creation", p.status[0])

	// Call2
	completeLifecycle(&p)
	fmt.Printf("\n player struct is %v", p)

	callsPerMilliSecond := int64(1000 / 800)
	ch := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case <-ch:
				return
			default:
				go gettingCalled(i)
				time.Sleep(time.Duration(callsPerMilliSecond) * time.Millisecond)
			}
			i++
		}
	}()

	time.Sleep(2 * time.Second)
	close(ch)

}
