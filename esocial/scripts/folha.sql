-- 1. Tabela de Períodos de Apuração (Domínio de Competência)
CREATE TABLE periodo_apuracao (
    id_periodo SERIAL PRIMARY KEY,
    mes_ano CHAR(7) UNIQUE NOT NULL, -- Formato 'AAAA-MM' ou 'AAAA-MM (Mensal/13º)'
    indicador_apuracao INT NOT NULL, -- 1: Mensal, 2: 13º salário
    status_fechamento VARCHAR(20) NOT NULL DEFAULT 'ABERTO'
);

-- 2. Tabela de Tipos de Rubrica (Tabela de Domínio - S-1010)
CREATE TABLE tipo_rubrica (
    id_tipo_rubrica SERIAL PRIMARY KEY,
    codigo_rubrica VARCHAR(30) NOT NULL,
    empregador_id INT NOT NULL, -- Referência à tabela de empregador
    desc_rubrica VARCHAR(255) NOT NULL,
    natureza_rubrica INT NOT NULL, -- Código da natureza da rubrica (Tabela 3 do eSocial)
    tipo_apuracao_rubrica INT NOT NULL, -- 1: Vencimento, 2: Desconto, 3: Informativa
    -- unique para a combinação de rubrica por empregador
    CONSTRAINT uk_rubrica_empregador UNIQUE (codigo_rubrica, empregador_id)
);

-- 3. Tabela Principal de Remuneração (Cabeçalho S-1200 por Trabalhador/Período)
CREATE TABLE remuneracao_trabalhador (
    id_remuneracao SERIAL PRIMARY KEY,
    empregador_id INT NOT NULL,
    trabalhador_cpf VARCHAR(11) NOT NULL,
    id_periodo INT REFERENCES periodo_apuracao(id_periodo),
    nr_recibo VARCHAR(44)
);

-- 4. Tabela de Itens da Remuneração / Lançamentos de Rubricas (Itens do S-1200)
CREATE TABLE item_remuneracao (
    id_item SERIAL PRIMARY KEY,
    remuneracao_id INT REFERENCES remuneracao_trabalhador(id_remuneracao) ON DELETE CASCADE,
    tipo_rubrica_id INT REFERENCES tipo_rubrica(id_tipo_rubrica),
    valor_rubrica NUMERIC(15,2) NOT NULL,
    quantidade_rubrica NUMERIC(10,4),
    fator_rubrica NUMERIC(5,2)
);

-- 5. Tabela de Bases de Cálculo e Tributos (INSS, FGTS, IRRF consolidados)
CREATE TABLE base_calculo_tributo (
    id_base SERIAL PRIMARY KEY,
    remuneracao_id INT REFERENCES remuneracao_trabalhador(id_remuneracao) ON DELETE CASCADE,
    tipo_tributo VARCHAR(20) NOT NULL, -- 'INSS', 'FGTS', 'IRRF'
    valor_base NUMERIC(15,2) NOT NULL,
    valor_tributo NUMERIC(15,2) NOT NULL
);