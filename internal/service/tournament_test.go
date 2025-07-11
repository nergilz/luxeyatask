package service_test

import (
	"testing"

	"github.com/nergilz/luxeyatask/internal/repository"
	"github.com/nergilz/luxeyatask/internal/service"
)

func TestCalculate(t *testing.T) {
	members := []repository.TournamentMembers{
		{ID: 123, TournamentID: 121234, UserID: 1, Score: 250},
		{ID: 124, TournamentID: 221234, UserID: 2, Score: 200},
		{ID: 129, TournamentID: 721234, UserID: 7, Score: 1050},
		{ID: 125, TournamentID: 321234, UserID: 3, Score: 350},
		{ID: 126, TournamentID: 421234, UserID: 4, Score: 850},
		{ID: 120, TournamentID: 821234, UserID: 8, Score: 950},
		{ID: 127, TournamentID: 521234, UserID: 5, Score: 450},
		{ID: 128, TournamentID: 621234, UserID: 6, Score: 550},
	}

	expect := map[string]repository.TournamentMembers{
		"first":  {ID: 129, TournamentID: 721234, UserID: 7, Score: 1050},
		"second": {ID: 120, TournamentID: 821234, UserID: 8, Score: 950},
		"third":  {ID: 126, TournamentID: 421234, UserID: 4, Score: 850},
	}

	actual := service.Calculation(members)

	if expect["first"].ID != actual["first"].ID {
		t.Errorf("error: expect: %d, actual: %d", expect["first"].ID, actual["first"].ID)
	}

	if expect["first"].Score != actual["first"].Score {
		t.Errorf("error: expect: %d, actual: %d", expect["first"].Score, actual["first"].Score)
	}

	if expect["second"].Score != actual["second"].Score {
		t.Errorf("error: expect: %d, actual: %d", expect["first"].Score, actual["first"].Score)
	}

	if expect["third"].Score != actual["third"].Score {
		t.Errorf("error: expect: %d, actual: %d", expect["first"].Score, actual["first"].Score)
	}
}
