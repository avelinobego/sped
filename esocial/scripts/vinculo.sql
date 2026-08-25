-- 1. Tabela de Tipos de Vínculo e Regimes (Tabelas de Domínio do eSocial)
CREATE TABLE tipo_vinculo (
    id_tipo_vinculo SERIAL PRIMARY KEY,
    codigo INT UNIQUE NOT NULL, -- 1: Empregado CLT, 2: Trabalhador Temporário, 3: Diretor não empregado, etc.
    descricao VARCHAR(150) NOT NULL
);

-- 2. Tabela Principal de Vínculo Empregatício
CREATE TABLE vinculo_empregaticio (
    id_vinculo SERIAL PRIMARY KEY,
    empregador_id INT NOT NULL, -- Referência à tabela de empregador
    trabalhador_id INT NOT NULL, -- Referência à tabela de trabalhador
    matricula VARCHAR(30) NOT NULL,
    tipo_vinculo_id INT REFERENCES tipo_vinculo(id_tipo_vinculo),
    data_admissao DATE NOT NULL,
    indicador_admissao INT NOT NULL, -- 1: Normal, 2: Transferência de outro empregador
    CONSTRAINT uk_empregador_matricula UNIQUE (empregador_id, matricula)
);

-- 3. Tabela de Condições Contratuais e Categoria (S-2200 / S-2206)
CREATE TABLE vinculo_contrato (
    id_contrato SERIAL PRIMARY KEY,
    vinculo_id INT REFERENCES vinculo_empregaticio(id_vinculo) ON DELETE CASCADE,
    categoria_trabalhador INT NOT NULL, -- Tabela 1 do eSocial (Ex: 101 - Empregado geral)
    codigo_cargo VARCHAR(30),
    descricao_cargo VARCHAR(150) NOT NULL,
    codigo_funcao VARCHAR(30),
    remuneracao_base NUMERIC(15,2) NOT NULL,
    unidade_salario INT NOT NULL, -- 1: Hora, 2: Dia, 3: Semana, 4: Quinzena, 5: Mês
    data_inicio_vigencia DATE NOT NULL,
    data_fim_vigencia DATE
);

-- 4. Tabela de Jornada de Trabalho Associada ao Vínculo
CREATE TABLE vinculo_jornada (
    id_jornada SERIAL PRIMARY KEY,
    vinculo_id INT REFERENCES vinculo_empregaticio(id_vinculo) ON DELETE CASCADE,
    tipo_jornada INT NOT NULL, -- 1: Submetido a horário, 2: Exoneração de horário (Art. 62), etc.
    descricao_jornada TEXT,
    data_inicio_vigencia DATE NOT NULL
);

-- 5. Tabela de Desligamento / Rescisão (S-2299)
CREATE TABLE vinculo_desligamento (
    id_desligamento SERIAL PRIMARY KEY,
    vinculo_id INT REFERENCES vinculo_empregaticio(id_vinculo) ON DELETE CASCADE,
    data_desligamento DATE NOT NULL,
    motivo_desligamento INT NOT NULL, -- Tabela 19 do eSocial
    indicador_obito BOOLEAN NOT NULL DEFAULT FALSE,
    observacao TEXT
);