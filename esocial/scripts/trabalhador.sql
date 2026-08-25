-- 1. Tabela de Identificação do Trabalhador (Dados Pessoais e Civis)
CREATE TABLE trabalhador (
    id_trabalhador SERIAL PRIMARY KEY,
    cpf VARCHAR(11) UNIQUE NOT NULL,
    nis_pis VARCHAR(11) NOT NULL,
    nome VARCHAR(255) NOT NULL,
    sexo CHAR(1) NOT NULL, -- 'M', 'F'
    raca_cor INT NOT NULL, -- Tabela 2 do eSocial
    estado_civil INT, -- Tabela 1 do eSocial
    data_nascimento DATE NOT NULL,
    pais_nascimento VARCHAR(3) NOT NULL,
    nacionalidade VARCHAR(3) NOT NULL
);

-- 2. Tabela de Endereço do Trabalhador
CREATE TABLE trabalhador_endereco (
    id_endereco SERIAL PRIMARY KEY,
    trabalhador_id INT REFERENCES trabalhador(id_trabalhador) ON DELETE CASCADE,
    tipo_logradouro VARCHAR(20),
    logradouro VARCHAR(255) NOT NULL,
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cep VARCHAR(8) NOT NULL,
    codigo_municipio_ibge INT NOT NULL,
    uf CHAR(2) NOT NULL
);

-- 3. Tabela de Vínculo Contratual (Cabeçalho do S-2200 / S-2300)
CREATE TABLE vinculo_contratual (
    id_vinculo SERIAL PRIMARY KEY,
    trabalhador_id INT REFERENCES trabalhador(id_trabalhador) ON DELETE CASCADE,
    empregador_id INT NOT NULL, -- Referência à tabela de empregador
    matricula VARCHAR(30) NOT NULL,
    tipo_regime_trabalhista INT NOT NULL, -- 1: CLT, 2: Estatutário
    tipo_regime_previdenciario INT NOT NULL, -- 1: RGPS, 2: RPPS
    data_admissao DATE NOT NULL,
    CONSTRAINT uk_vinculo_empregador UNIQUE (empregador_id, matricula)
);

-- 4. Tabela de Contrato de Trabalho (Condições Remuneratórias e Jornada)
CREATE TABLE contrato_trabalho (
    id_contrato SERIAL PRIMARY KEY,
    vinculo_id INT REFERENCES vinculo_contratual(id_vinculo) ON DELETE CASCADE,
    cargo_descricao VARCHAR(100) NOT NULL,
    funcao_codigo VARCHAR(30),
    salario_base NUMERIC(15,2) NOT NULL,
    unidade_salario INT NOT NULL, -- 1: Por hora, 2: Por dia, 3: Por semana, 4: Por quinzena, 5: Por mês
    tipo_contr_trabalho INT NOT NULL, -- 1: Prazo indeterminado, 2: Prazo determinado
    data_inicio_vigencia DATE NOT NULL,
    data_fim_vigencia DATE
);

-- 5. Tabela de Dependentes (S-2200 / S-2205)
CREATE TABLE trabalhador_dependente (
    id_dependente SERIAL PRIMARY KEY,
    trabalhador_id INT REFERENCES trabalhador(id_trabalhador) ON DELETE CASCADE,
    cpf VARCHAR(11),
    nome VARCHAR(255) NOT NULL,
    data_nascimento DATE NOT NULL,
    tipo_dependente INT NOT NULL, -- Tabela 07 do eSocial (Cônjuge, Filho, etc.)
    dep_irrf BOOLEAN NOT NULL DEFAULT FALSE, -- Dependente para fins de IRRF
    dep_sf BOOLEAN NOT NULL DEFAULT FALSE     -- Dependente para fins de Salário Família
);