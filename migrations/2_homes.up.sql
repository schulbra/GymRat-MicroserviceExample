CREATE TABLE IF NOT EXISTS homes {
    id SERIAL PRIMARY KEY,
    price BIGINT NOT NULL
    agent_id BIGINT REFERENCES agents(id)
    }
