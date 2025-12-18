-- Add faction type enum
CREATE TYPE faction_type AS ENUM ('red', 'blue', 'green');

-- Add faction column to users
ALTER TABLE users ADD COLUMN faction faction_type;

-- Ships table
CREATE TABLE ships (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    ship_type VARCHAR(50) NOT NULL DEFAULT 'betha_1',
    slot INT NOT NULL CHECK (slot BETWEEN 1 AND 5),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE(user_id, slot)
);

CREATE INDEX idx_ships_user ON ships(user_id);
