-- name: CreateAccount :exec
INSERT INTO account (
    display_name
) VALUES (
    @display_name
);