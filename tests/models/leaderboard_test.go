package main

import (
	"fmt"
	"github.com/arashrasoulzadeh/go-game-engine/models"
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
			Score: rand.Intn(100),
		})
	}
	fmt.Println(lb.Sorted())

}
