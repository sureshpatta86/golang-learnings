package main

import (
	"fmt"
)

type Cricket struct {
	Score        int
	TotalWickets int
}

type Football struct {
	Score       string
	YellowCards int
}

type Sports interface {
	ScoreInfo()
}

func (c Cricket) ScoreInfo() {
	fmt.Println("Cricket Score is ", c.Score)
}

func (f Football) ScoreInfo() {
	fmt.Println("Football Score is ", f.Score)
}

func (c Cricket) Wickets() {
	fmt.Println("Cricket No of Wickets is ", c.TotalWickets)
}

func (f Football) NoOfYellowCards() {
	fmt.Println("Football No of Yellow Cards is ", f.YellowCards)
}

func main() {
	//Type Assertion
	sportsUpdates := []Sports{
		Cricket{Score: 100, TotalWickets: 2},
		Football{Score: "1-1", YellowCards: 2},
	}

	for _, value := range sportsUpdates {
		value.ScoreInfo()
		//value.Wickets() //Error: value.Wickets undefined (type Sports has no field or method Wickets)

		//Type Assertion
		//value.(Cricket).Wickets() //Panic: interface conversion: Sports is Football, not Cricket

		//Type Assertion with ok for Cricket
		if cricket, ok := value.(Cricket); ok {
			cricket.Wickets()
		}

		//Type Assertion with ok for Football
		if football, ok := value.(Football); ok {
			football.NoOfYellowCards()
		}

	}

}
