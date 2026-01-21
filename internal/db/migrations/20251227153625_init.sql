-- +goose Up
CREATE TYPE account_status AS ENUM ('active', 'inactive');

CREATE TYPE auth_provider AS ENUM ('username', 'google', 'phone', 'email');

CREATE TABLE
    users (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        email text UNIQUE,
        phone text UNIQUE,
        name text NOT NULL,
        date_of_birth date,
        profile_pic text,
        email_verified boolean DEFAULT false,
        phone_verified boolean DEFAULT false,
        last_login timestamptz,
        account_status account_status DEFAULT 'active',
        deactivated_at timestamptz,
        deactivation_reason text,
        created_at timestamptz DEFAULT now (),
        updated_at timestamptz DEFAULT now (),
        CHECK (
            email IS NOT NULL
            OR phone IS NOT NULL
        )
    );

CREATE TABLE
    auth_identities (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        provider auth_provider NOT NULL,
        identifier text NOT NULL,
        verified boolean DEFAULT false,
        created_at timestamptz DEFAULT now (),
        updated_at timestamptz DEFAULT now (),
        UNIQUE (provider, identifier)
    );

CREATE TABLE
    sessions (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        user_id uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
        school_id uuid NOT NULL,
        provider auth_provider NOT NULL,
        refresh_token_hash text NOT NULL,
        expires_at timestamptz NOT NULL,
        is_active boolean DEFAULT true,
        ip_address text,
        user_agent text,
        device_type text,
        device_name text,
        created_at timestamptz DEFAULT now ()
    );

CREATE INDEX idx_sessions_user_active ON sessions (user_id, is_active);

CREATE INDEX idx_sessions_school_user ON sessions (school_id, user_id);

CREATE INDEX idx_auth_identity_user ON auth_identities (user_id);

-- +goose Down
DROP INDEX IF EXISTS idx_auth_identity_user;

DROP INDEX IF EXISTS idx_sessions_school_user;

DROP INDEX IF EXISTS idx_sessions_user_active;

DROP TABLE IF EXISTS sessions;

DROP TABLE IF EXISTS auth_identities;

DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS device_type;

DROP TYPE IF EXISTS auth_provider;

DROP TYPE IF EXISTS account_status;