CREATE TABLE IF NOT EXISTS produtos (
    id            BIGSERIAL PRIMARY KEY,
    nome          VARCHAR(255) NOT NULL,
    descricao     VARCHAR(80) NOT NULL,
    preco         decimal(10,2) NOT NULL,
    quantidade    INT NOT NULL
);