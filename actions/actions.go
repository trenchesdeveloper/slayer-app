package actions

import (
	"math/rand"
	"time"
)

// generate rand source
var randSource = rand.NewSource(time.Now().UnixNano())

// generate random number
var randGenerator = rand.New(randSource)

func AttackMoster (isSpecialAttack bool){

}

func HealPlayer(){

}

func AttackPlayer (){

}

func generateRandBetween (min int, max int) int {
	return randGenerator.Intn(max - min) + min
}