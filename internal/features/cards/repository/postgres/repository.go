package cards_postgres_repository

import core_postgres_pool "github.com/dmitdub/go-flashcards/internal/core/repository/postgres/pool"

type CardsRepository struct {
	pool core_postgres_pool.Pool
}

func NewCardsRepository(
	pool core_postgres_pool.Pool,
) *CardsRepository {
	return &CardsRepository{
		pool: pool,
	}
}
