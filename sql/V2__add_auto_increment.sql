CREATE SEQUENCE IF NOT EXISTS produtos_id_seq;

ALTER TABLE produtos
    ALTER COLUMN id SET DEFAULT nextval('produtos_id_seq');

SELECT setval(
    'produtos_id_seq',
    COALESCE((SELECT MAX(id) FROM produtos), 1),
    true
);

ALTER SEQUENCE produtos_id_seq
    OWNED BY produtos.id;
