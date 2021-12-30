CREATE TABLE chargers
(
    id           INTEGER PRIMARY KEY,

    name         TEXT                                               NOT NULL UNIQUE,
    lat          DECIMAL(12)                                        NOT NULL,
    lon          DECIMAL(12)                                        NOT NULL,

    date_created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE reservations
(
    user_id    TEXT                                               NOT NULL,
    charger_id INTEGER REFERENCES chargers (id) ON DELETE CASCADE NOT NULL,
    time_from  TIMESTAMP WITH TIME ZONE                           NOT NULL,
    time_until TIMESTAMP WITH TIME ZONE                           NOT NULL
);