package decks_postgres_repository

import core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"

type DecksRepository struct {
	pool core_postgres_pool.Pool
}

func NewDecksRepository(
	pool core_postgres_pool.Pool,
) *DecksRepository {
	return &DecksRepository{
		pool: pool,
	}
}
