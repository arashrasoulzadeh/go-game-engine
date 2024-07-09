package main

import (
	"fmt"
	"go-game/models"
	"math/rand"
	"testing"
)

func TestConstructor(t *testing.T) {
	lb := models.CreateLeaderboard()

	if lb == nil {
		t.Error("Constructor fails!")
		t.Fail()
	}
	lb.Init()

	for _ = range 100 {
		lb.Add(models.LeaderboardItem{
			Index: rand.Intn(100),
		})
	}
	fmt.Println(lb.Sorted())

}
