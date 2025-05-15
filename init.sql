CREATE TABLE IF NOT EXISTS vehicle_locations (
    id SERIAL PRIMARY KEY,
    vehicle_id VARCHAR(255) NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    logitude DOUBLE PRECISION NOT NULL,
    timestamp BIGINT NOT NULL
);
