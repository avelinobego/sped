-- 1. Tabela de Tipos de Eventos de SST (Domínio do eSocial)
CREATE TABLE tipo_evento_sst (
    id_tipo_evento SERIAL PRIMARY KEY,
    codigo_evento VARCHAR(10) UNIQUE NOT NULL, -- Ex: 'S-2210', 'S-2220', 'S-2240'
    descricao VARCHAR(100) NOT NULL
);

-- 2. Tabela Principal de SST (Cabeçalho do Evento)
CREATE TABLE sst_evento (
    id_sst SERIAL PRIMARY KEY,
    empregador_id INT NOT NULL, -- Referência à tabela de empregador
    trabalhador_cpf VARCHAR(11) NOT NULL,
    tipo_evento_id INT REFERENCES tipo_evento_sst(id_tipo_evento),
    data_evento DATE NOT NULL,
    nr_recibo VARCHAR(44) -- Recibo do evento retornado pelo eSocial
);

-- 3. Tabela de Condições Ambientais do Trabalho (S-2240 - Fatores de Risco)
CREATE TABLE sst_fator_risco (
    id_fator_risco SERIAL PRIMARY KEY,
    sst_id INT REFERENCES sst_evento(id_sst) ON DELETE CASCADE,
    codigo_tabela_24 VARCHAR(10) NOT NULL, -- Código do agente nocivo (Tabela 24 do eSocial)
    descricao_risco TEXT,
    dsc_eficaz VARCHAR(1) -- 'S' ou 'N' para EPI eficaz
);

-- 4. Tabela de Monitoramento da Saúde - ASOs (S-2220)
CREATE TABLE sst_aso (
    id_aso SERIAL PRIMARY KEY,
    sst_id INT REFERENCES sst_evento(id_sst) ON DELETE CASCADE,
    tp_exame_ocupa INT NOT NULL, -- 0: Admissional, 1: Periódico, 2: Retorno, etc.
    data_aso DATE NOT NULL,
    resultado_aso INT NOT NULL -- 1: Apto, 2: Inapto
);

-- 5. Tabela de Acidente de Trabalho (S-2210)
CREATE TABLE sst_acidente_trabalho (
    id_acidente SERIAL PRIMARY KEY,
    sst_id INT REFERENCES sst_evento(id_sst) ON DELETE CASCADE,
    data_acidente TIMESTAMP NOT NULL,
    tipo_acidente INT NOT NULL, -- 1: Típico, 2: Trajeto, 3: Doença profissional
    houve_obito BOOLEAN NOT NULL,
    cat_originaria_id VARCHAR(44) -- Caso seja CAT de reabertura
);