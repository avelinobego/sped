-- 1. Tabela de Metadados das Tabelas do eSocial (Catálogo de Tabelas)
CREATE TABLE tabela_esocial (
    id_tabela SERIAL PRIMARY KEY,
    codigo_tabela VARCHAR(10) UNIQUE NOT NULL, -- Ex: 'TAB-01', 'TAB-03', 'TAB-24'
    nome_tabela VARCHAR(255) NOT NULL,
    descricao TEXT
);

-- 2. Tabela de Versões de Vigência das Tabelas
CREATE TABLE vigencia_tabela (
    id_vigencia SERIAL PRIMARY KEY,
    tabela_id INT REFERENCES tabela_esocial(id_tabela) ON DELETE CASCADE,
    versao_leiaute VARCHAR(20) NOT NULL,
    data_inicio_vigencia DATE NOT NULL,
    data_fim_vigencia DATE
);

-- 3. Tabela de Registros/Itens Genéricos das Tabelas do eSocial
CREATE TABLE item_tabela_esocial (
    id_item SERIAL PRIMARY KEY,
    vigencia_id INT REFERENCES vigencia_tabela(id_vigencia) ON DELETE CASCADE,
    codigo_item VARCHAR(50) NOT NULL, -- O código interno do item na tabela (ex: código da natureza da rubrica)
    descricao_item TEXT NOT NULL,
    informacao_complementar JSONB -- Flexibilidade para atributos específicos de cada tabela do eSocial
);