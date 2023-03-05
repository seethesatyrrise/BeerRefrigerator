create table beer
(
    id bigserial not null
        constraint beer_pk
            primary key,
    title text,
    abv text,
    expires_at timestamp
);

INSERT INTO beer (title, abv, expires_at)
VALUES ('test', '10', '2023-03-10')