-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id UUID PRIMARY KEY DEFAULT gen_random_UUID(),
    email TEXT UNIQUE,
    phone TEXT UNIQUE,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE IF NOT EXISTS auth_identities(
    id UUID PRIMARY KEY DEFAULT gen_random_UUID(),
    user_id UUID NOT NULL,
    identity TEXT NOT NULL,
    username TEXT UNIQUE,
    password TEXT,
    email TEXT,
    phone TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),

    CONSTRAINT unique_user_identity UNIQUE (user_id, identity),
    CONSTRAINT fk_auth_identities_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,

    CONSTRAINT chk_identity_type CHECK (identity in ('password','google')),
    CONSTRAINT chk_password_required CHECK (
        (identity = 'password' AND password IS NOT NULL)
        OR
        (identity <> 'password' AND password IS NULL)
    )
);

CREATE TABLE IF NOT EXISTS auth_sessions(
    id UUID PRIMARY KEY DEFAULT gen_random_UUID(),
    user_id UUID NOT NULL,
    identity_id UUID NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    
    CONSTRAINT fk_auth_sessions_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_auth_sessions_auth_identities FOREIGN KEY (identity_id) REFERENCES auth_identities(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS auth_sessions;
DROP TABLE IF EXISTS auth_identities;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
