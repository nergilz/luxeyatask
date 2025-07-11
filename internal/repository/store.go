package repository

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	logger *slog.Logger
	conn   *pgx.Conn
}

func New(ctx context.Context, url string, logger *slog.Logger) (*Storage, error) {
	connection, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}

	err = connection.Ping(ctx)
	if err != nil {
		return nil, err
	}

	logger.Info("DB Ping: OK!")

	return &Storage{
		conn:   connection,
		logger: logger,
	}, nil
}

// create tournament
// update tournament
// delete tournament
// run turnament

func (s Storage) GetAllTournaments() ([]Tournament, error) {
	res := make([]Tournament, 100)

	return res, nil
}

func (s Storage) RegistrationUserTournament(ctx context.Context, userId uint64) error {

	return nil
}

func (s Storage) GetTournamentByUserID(ctx context.Context, userId uint64) (Tournament, error) {

	return Tournament{}, nil
}

func (s Storage) GetTournamentByID(ctx context.Context, userId uint64) (Tournament, error) {

	return Tournament{}, nil
}

func (s Storage) GetTournamentByName(ctx context.Context, name string) (Tournament, error) {

	return Tournament{}, nil
}

func (s Storage) GetMembersByTournamentID(ctx context.Context, tourID uint64) ([]TournamentMembers, error) {
	members := make([]TournamentMembers, 100)

	return members, nil
}

func (s Storage) StopTournament(ctx context.Context, tourID uint64) error {
	// изменение статуса запущенного утрнира

	return nil
}

func (s Storage) SetBalanceByUserID(ctx context.Context, bonus uint64, userID uint64) error {
	// начислить бонус
	return nil
}
