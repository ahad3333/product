
CREATE TABLE branch (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE category (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    photo text,
    type_id UUID DEFAULT NULL,
    branch_id UUID NOT NULL REFERENCES branch(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE product (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT,
    photo text,
    category_id UUID NOT NULL REFERENCES category(id),
    branch_id UUID NOT NULL REFERENCES branch(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
