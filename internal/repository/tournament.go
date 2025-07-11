package repository

import "time"

type Tournament struct {
	ID                  uint64
	Name                string    // оргигинальное имя турнира
	CreateAT            time.Time // время создания
	DurationAT          uint16    // продолжительность турнира в количестве часов
	TimeEnd             time.Time // окончание турнира
	Price               uint64    // стоимость участия, вычитается из баланса
	MaxParticipant      int       // максимальное кольчество участников
	MinParticipant      int       // минимальное кольчество участников
	MaxQuantityAttempts int       // максимальное количество попыток оптравить рекорд
	IsAutomatic         bool
	IsRunning           bool
	// TournamentParticipants []uint64 // участники турнира // ?
}

type User struct {
	ID       uint64
	UserName string
	Balance  uint64 // баланс
}

type TournamentMembers struct {
	ID               uint64
	TournamentID     uint64
	UserID           uint64
	Score            int // рекорд, отправленный юзером в турнир
	QuantityAttempts int // количество попыток оптравить рекорд
}

type Winners struct {
	ID            uint64
	TournamentID  uint64
	FirstPlaceID  uint64
	SecondPlaceID uint64
	ThirdPlaceID  uint64
}
