CREATE TABLE periodos_folha (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    mes INTEGER NOT NULL CHECK (mes BETWEEN 1 AND 13),
    ano INTEGER NOT NULL,
    tipo_folha VARCHAR(20) NOT NULL,
    data_inicio DATE NOT NULL,
    data_fim DATE NOT NULL,
    data_pagamento DATE NOT NULL,
    situacao VARCHAR(20) DEFAULT 'ABERTA',
    data_calculo TIMESTAMP WITH TIME ZONE,
    data_fechamento TIMESTAMP WITH TIME ZONE,
    protocolo_fechamento VARCHAR(50),
    recibo_fechamento VARCHAR(50),
    data_envio_esocial TIMESTAMP WITH TIME ZONE,
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(empresa_id, mes, ano, tipo_folha)
);

CREATE TABLE eventos_remuneracao (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    periodo_folha_id UUID NOT NULL REFERENCES periodos_folha(id) ON DELETE CASCADE,
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    tipo_evento VARCHAR(20) NOT NULL,
    total_vencimentos DECIMAL(10,2) DEFAULT 0,
    total_descontos DECIMAL(10,2) DEFAULT 0,
    valor_liquido DECIMAL(10,2) DEFAULT 0,
    base_inss DECIMAL(10,2) DEFAULT 0,
    valor_inss DECIMAL(10,2) DEFAULT 0,
    base_irrf DECIMAL(10,2) DEFAULT 0,
    valor_irrf DECIMAL(10,2) DEFAULT 0,
    base_fgts DECIMAL(10,2) DEFAULT 0,
    valor_fgts DECIMAL(10,2) DEFAULT 0,
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP WITH TIME ZONE,
    status_esocial VARCHAR(20),
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(periodo_folha_id, vinculo_id, tipo_evento)
);

CREATE TABLE itens_remuneracao (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    evento_remuneracao_id UUID NOT NULL REFERENCES eventos_remuneracao(id) ON DELETE CASCADE,
    rubrica_id UUID NOT NULL REFERENCES rubricas(id),
    quantidade DECIMAL(10,4) DEFAULT 1,
    valor_unitario DECIMAL(10,2),
    valor_total DECIMAL(10,2) NOT NULL,
    tipo_item CHAR(1) NOT NULL,
    referencia VARCHAR(100),
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE auditoria (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    tabela VARCHAR(100) NOT NULL,
    operacao VARCHAR(10) NOT NULL,
    registro_id UUID NOT NULL,
    dados_anteriores JSONB,
    dados_novos JSONB,
    usuario_nome VARCHAR(255),
    data_hora TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    ip_origem VARCHAR(45)
);
