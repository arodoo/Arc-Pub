// Arc-Pub - Metaverso 2D MMO Social
// Copyright (c) 2024. MIT License.

package postgres

import (
	"context"
	"log"

	"github.com/arc-pub/server/internal/application/auth"
	"github.com/arc-pub/server/internal/domain/user"
	"github.com/google/uuid"
)

const (
	adminEmail    = "admin@dev.local"
	adminPassword = "admin123"
)

// Seeder creates initial data for development.
type Seeder struct {
	users  auth.UserRepository
	hasher auth.PasswordHasher
}

// NewSeeder creates a Seeder with dependencies.
func NewSeeder(
	users auth.UserRepository,
	hasher auth.PasswordHasher,
) *Seeder {
	return &Seeder{users: users, hasher: hasher}
}

// SeedAdmin creates admin user if not exists.
func (s *Seeder) SeedAdmin(ctx context.Context) error {
	exists, err := s.users.ExistsByEmail(ctx, adminEmail)
	if err != nil {
		return err
	}
	if exists {
		log.Println("Admin user already exists, skipping seed")
		return nil
	}

	hash, err := s.hasher.Hash(adminPassword)
	if err != nil {
		return err
	}

	admin := user.NewUser(
		uuid.New(),
		adminEmail,
		hash,
		user.RoleAdmin,
	)

	if err := s.users.Create(ctx, admin); err != nil {
		return err
	}

	log.Printf("Created admin user: %s", adminEmail)
	return nil
}
