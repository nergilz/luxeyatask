package service

import (
	"log/slog"

	"github.com/nergilz/luxeyatask/internal/repository"
)

type IStorage interface {
	GetAllTournaments() ([]repository.Tournament, error)
	RegistrationUserTournament(userId uint64) error
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

// // запускает одиночный турнир
// func (s Sevice) StartSingleTournament(ctx context.Context, tour repository.Tournament) {
// 	stopAt := time.After()

// 	go func() {

// 	}()
// }

// // запускает автоматический турнир
// func (s Sevice) AutoRepeatTournament(ctx context.Context, tour repository.Tournament) {
// 	stopAt := time.After()

// 	go func() {

// 	}()
// }

// // запускает все не законченные турниры при запуске сервиса
// // нужна проверка участия пользователя в турнире
// func (s Sevice) StartAllTournaments(ctx context.Context, tour repository.Tournament) {
// 	stopAt := time.After()

// 	go func() {

// 	}()
// }
