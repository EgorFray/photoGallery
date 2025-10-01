package testutils

import (
	"context"
	"database/sql"
	postRepo "gallery/backend/internal/repository/posts"
	userRepo "gallery/backend/internal/repository/user"
	"log"
	"testing"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type TestRepo struct {
	DB *sql.DB
	PostRepo postRepo.PostRepositoryInterface
	UserRepo userRepo.UserRepositoryInterface
}

func SetupTestRepo(t *testing.T) *TestRepo {
	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
			"postgres:16-alpine",
			postgres.WithDatabase("testDb"),
			postgres.WithUsername("testUser"),
			postgres.WithPassword("testPassword"),
			postgres.BasicWaitStrategies(),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return nil
	}
	connStr, err := postgresContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("failed to get connection string: %v", err)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	_, err = db.Exec(`
	CREATE TABLE posts (
		id SERIAL PRIMARY KEY ,
		image TEXT NOT NULL,
		description TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		user_id INTEGER
	)
	`)

	if err != nil {
		log.Fatalf("failed to create table posts: %v", err)
	}

	t.Cleanup(func() {
		db.Close()
		if err := postgresContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})
	
		return &TestRepo {
			DB: db,
			PostRepo: postRepo.NewPostRepository(db),
			UserRepo: userRepo.NewUserRepository(db),
		}
}