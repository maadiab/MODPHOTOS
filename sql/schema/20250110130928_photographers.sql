-- +goose Up
-- +goose StatementBegin
CREATE TABLE photographers (
    id SERIAL PRIMARY KEY,
    phgname TEXT NOT NULL,
    phgusername TEXT NOT NULL,
    phgmobile TEXT NOT NULL,
    phgpassword TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE photographers;
-- +goose StatementEnd
