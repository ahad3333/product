
CREATE TABLE "user" (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    login VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    branch_id UUID REFERENCES branch(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE client (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
