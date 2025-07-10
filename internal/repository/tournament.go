package repository

import "time"

type Tournament struct {
	ID                  uint64
	Name                string
	CreateAT            time.Time // время создания
	DurationAT          uint16    // продолжительность турнира в количестве часов
	TimeEnd             time.Time // окончание турнира
	Price               uint64    // стоимость участия
	MaxParticipant      uint32    // максимальное кольчество участников
	MinParticipant      uint32    // минимальное кольчество участников
	MaxQuantityAttempts uint16    // количество попыток
	IsAutomatic         bool
	IsRunning           bool
	// TournamentParticipants []uint64 // участники турнира // ?
}

type User struct {
	ID       uint64
	UserName string
	Balance  uint64 // баланс
	Score    uint64 // рекорд
}

type TournamentMemebers struct {
	ID           uint64
	TournamentID uint64
	UserID       uint64
}

type Winners struct {
	ID            uint64
	TournamentID  uint64
	FirstPlaceID  uint64
	SecondPlaceID uint64
	ThirdPlaceID  uint64
}
