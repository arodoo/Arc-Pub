-- Servers table
CREATE TABLE servers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    region VARCHAR(50) NOT NULL,
    host VARCHAR(255),
    port INT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Add server_id to users
ALTER TABLE users ADD COLUMN server_id UUID REFERENCES servers(id);

-- Seed Mexico server
INSERT INTO servers (name, region) VALUES ('Mexico', 'latam');
