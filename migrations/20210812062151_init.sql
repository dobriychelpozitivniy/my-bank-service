-- +goose Up
CREATE TABLE bank (id INT NOT NULL, cash FLOAT NOT NULL, currency VARCHAR(10) NOT NULL);
INSERT OR IGNORE INTO bank (id, cash, currency) VALUES ('21', '0', 'SBP');

-- +goose Down
DROP TABLE bank;