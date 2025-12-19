-- Database schema for Arc-Pub
-- This file is the source of truth for sqlc code generation

CREATE TYPE faction_type AS ENUM ('red', 'blue', 'green');

CREATE TABLE servers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    region VARCHAR(50) NOT NULL,
    host VARCHAR(255),
    port INT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    hashed_password TEXT NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    faction faction_type,
    server_id UUID REFERENCES servers(id),
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);

CREATE TABLE ships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    ship_type VARCHAR(50) NOT NULL DEFAULT 'betha_1',
    slot INT NOT NULL CHECK (slot BETWEEN 1 AND 5),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(user_id, slot)
);

CREATE INDEX idx_ships_user ON ships(user_id);
