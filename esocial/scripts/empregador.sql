1. Tabela de Tipos de Inscrição (Domínio do eSocial - S-1000)
CREATE TABLE tipo_inscricao (
    id_tipo_inscricao SERIAL PRIMARY KEY,
    codigo INT UNIQUE NOT NULL, -- 1: CNPJ, 2: CPF, 3: CAEPF, 4: CNO
    descricao VARCHAR(50) NOT NULL
);

-- 2. Tabela de Empregadores
CREATE TABLE empregador (
    id_empregador SERIAL PRIMARY KEY,
    tipo_inscricao_id INT REFERENCES tipo_inscricao(id_tipo_inscricao),
    nr_inscricao VARCHAR(14) UNIQUE NOT NULL, -- CNPJ, CPF, etc.
    razao_social VARCHAR(255) NOT NULL,
    nome_fantasia VARCHAR(255),
    classificacao_tributaria INT NOT NULL,
    indicativo_pj INT, -- 1: Empresa, 2: Cooperativa, 3: Órgão Público
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE
);

-- 3. Tabela de Endereços (Normalizada para evitar repetição)
CREATE TABLE endereco (
    id_endereco SERIAL PRIMARY KEY,
    empregador_id INT REFERENCES empregador(id_empregador),
    tipo_logradouro VARCHAR(20),
    logradouro VARCHAR(255) NOT NULL,
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cep VARCHAR(8) NOT NULL,
    codigo_municipio_ibge INT NOT NULL,
    uf CHAR(2) NOT NULL
);

-- 4. Tabela de Contatos
CREATE TABLE contato (
    id_contato SERIAL PRIMARY KEY,
    empregador_id INT REFERENCES empregador(id_empregador),
    nome_responsavel VARCHAR(255),
    cpf_responsavel VARCHAR(11),
    telefone_fixo VARCHAR(15),
    telefone_celular VARCHAR(15),
    email VARCHAR(255)
);