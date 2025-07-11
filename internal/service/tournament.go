package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/nergilz/luxeyatask/internal/repository"
)

type IStorage interface {
	GetAllTournaments() ([]repository.Tournament, error)
	RegistrationUserTournament(ctx context.Context, userId uint64) error
	GetTournamentByName(ctx context.Context, name string) (repository.Tournament, error)
	GetMembersByTournamentID(ctx context.Context, tourID uint64) ([]repository.TournamentMembers, error)
	StopTournament(ctx context.Context, tourID uint64) error
	SetBalanceByUserID(ctx context.Context, bonus uint64, userID uint64) error
}

type Sevice struct {
	logger *slog.Logger
	Store  IStorage
}

func New(log *slog.Logger, store IStorage) Sevice {
	return Sevice{
		logger: log,
		Store:  store,
	}
}

// запускает одиночный турнир
// не предусмотренно удаление турнира после работы
func (s Sevice) StartSingleTournament(ctx context.Context, name string) {
	tour, err := s.Store.GetTournamentByName(ctx, name)
	if err != nil {
		s.logger.Error("autorepeat", slog.String("error", err.Error()))
	}

	timerStopAt := time.NewTimer(time.Hour * time.Duration(tour.DurationAT))

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-timerStopAt.C:
			err := s.Store.StopTournament(ctx, tour.ID)
			if err != nil {
				return
			}

			members, err := s.Store.GetMembersByTournamentID(ctx, tour.ID)
			if err != nil {
				return
			}

			err = s.calculationUsers(ctx, tour, members)
			if err != nil {
				return
			}

			return
		}
	}()
}

// запускает автоматический турнир
// todo - вопрос когда его остановить
func (s Sevice) AutoRepeatTournament(ctx context.Context, name string) {
	tour, err := s.Store.GetTournamentByName(ctx, name)
	if err != nil {
		s.logger.Error("autorepeat", slog.String("error", err.Error()))
	}

	tickerStopAt := time.NewTicker(time.Hour * time.Duration(tour.DurationAT))

	go func() {
		select {
		case <-ctx.Done():
			return
		case <-tickerStopAt.C:
			// calculation()

			return
		}
	}()
}

// запускает все не законченные турниры при запуске сервиса
// нужна проверка участия пользователя в турнире
// предусматреть механизм остановки турнира
func (s Sevice) StartAllTournaments(ctx context.Context, tour repository.Tournament) {

}

// калькулирует призы для участников турнира
func (s Sevice) calculationUsers(ctx context.Context, tour repository.Tournament, users []repository.TournamentMembers) error {
	if len(users) < tour.MinParticipant {
		return fmt.Errorf("min participant")
	}

	placesInTour := Calculation(users)

	firstUser := placesInTour["first"]
	secondUser := placesInTour["second"]
	thirdUser := placesInTour["third"]

	err := s.Store.SetBalanceByUserID(ctx, 500, firstUser.ID)
	if err != nil {
		return err
	}

	err = s.Store.SetBalanceByUserID(ctx, 300, secondUser.ID)
	if err != nil {
		return err
	}

	err = s.Store.SetBalanceByUserID(ctx, 100, thirdUser.ID)
	if err != nil {
		return err
	}

	return nil
}

// калькулирует места для участников турнира
func Calculation(members []repository.TournamentMembers) map[string]repository.TournamentMembers {
	winners := make(map[string]repository.TournamentMembers, 0)
	scores := make(map[int]repository.TournamentMembers, 0)

	// если рекорды одинаковые оставляем первый
	for _, user := range members {
		if _, ok := scores[user.Score]; !ok {
			scores[user.Score] = user
		}
	}

	first := 0
	second := 0
	third := 0

	for _, user := range scores {
		if user.Score > first {
			third = second
			second = first
			first = user.Score
		} else if user.Score > second {
			third = second
			second = user.Score
		} else if user.Score > third {
			third = user.Score
		}
	}

	winners["first"] = scores[first]
	winners["second"] = scores[second]
	winners["third"] = scores[third]

	return winners
}
