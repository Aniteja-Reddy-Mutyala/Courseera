package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int

const (
	club Suit = iota
	diamond
	heart
	spade
)

type Card struct {
	s   Suit
	pip int
}

func shuffle(d *[52]Card) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range d {
		j := rand.Intn(52)
		d[i], d[j] = d[j], d[i]
	}
}
func isFlush(h []Card) bool {
	var ccount, dcount, hcount, scount int
	for _, v := range h {
		switch v.s {
		case club:
			ccount++
		case diamond:
			dcount++
		case heart:
			hcount++
		case spade:
			scount++

		}

	}
	if ccount>=5 || dcount>=5 || hcount>=5 ||scount>=5{
		return true;
	} else{
		return false;
	}

}
func (c Card) printCard(){
	var suitName,pipName string;
	switch c.s{
	case club:
		suitName="Clubs";
	case diamond:
		suitName="Diamonds";
	case heart:
		suitName="Hearts";
	case spade:
		suitName="Spades";
	}
	switch c.pip {
	case 1: pipName="Ace of ";
	case 2: pipName="Two of ";
	case 3: pipName="Three of ";
	case 4:pipName="Four of ";
	case 5:pipName="Five of ";
	case 6:pipName="Six of ";
	case 7:pipName="Seven of ";
	case 8:pipName="Eight of ";
	case 9:pipName="Nine of ";
	case 10:pipName="Ten of ";
	case 11:pipName="Jack of ";
	case 12:pipName="Queen of ";
	case 13:pipName="King of ";
		
	}
	fmt.Printf("The card is %s%s\t",pipName,suitName);
}
func main(){
	fmt.Println("Using Card Struct");
	var deck [52]Card;
	var fcnt int;
	var totcnt int;
	fmt.Println("No of trials");
	fmt.Scanf("%d",&totcnt);
	for i:=0; i< 52;i++{
		deck[i].pip=i%13 +1;
	}
	for i:=0;i<13;i++{
		deck[i].s=club;
	}
	for i:=13;i<26;i++{
		deck[i].s=diamond
	}
	for i:=26;i<39;i++{
		deck[i].s=heart;
	}
	for i:=39;i<52;i++{
		deck[i].s=spade;
	}
	hand:=make([]Card,7);
	for i:=0;i<totcnt;i++{
		shuffle(&deck);
		hand[0]=deck[0];
		hand[1]=deck[1];
		hand[2]=deck[2];
		hand[3]=deck[3];
		hand[4]=deck[4];
		hand[5]=deck[5];
		hand[6]=deck[6];
		if isFlush(hand){
			fcnt++;
			if fcnt%100 ==0{
				hand[0].printCard();
				hand[1].printCard();
				hand[2].printCard();
				hand[3].printCard();
				hand[4].printCard();
				hand[5].printCard();
				hand[6].printCard();
			}
		}
	

	}
   fmt.Println("\n Flushes per " ,fcnt, "/ ",totcnt);
}
