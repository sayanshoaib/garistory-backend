-- Analyze a slow query
EXPLAIN ANALYZE SELECT * FROM vehicles WHERE make = 'BMW';

-- Create an index on 'make' column
CREATE INDEX idx_vehicles_make ON  vehicles (make);
