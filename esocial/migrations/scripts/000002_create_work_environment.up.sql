CREATE TABLE horarios (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    tipo_jornada CHAR(1) NOT NULL,
    duracao_jornada_semanal INTEGER,
    tipo_intervalo CHAR(1),
    duracao_intervalo INTEGER,
    ativo BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(empresa_id, codigo)
);

CREATE TABLE horarios_detalhes (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    horario_id UUID NOT NULL REFERENCES horarios(id) ON DELETE CASCADE,
    dia_semana INTEGER NOT NULL,
    horario_entrada TIME NOT NULL,
    horario_saida TIME NOT NULL,
    duracao_jornada INTEGER NOT NULL,
    UNIQUE(horario_id, dia_semana)
);

CREATE TABLE ambientes_trabalho (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    local_ambiente CHAR(1) NOT NULL,
    tipo_inscricao CHAR(1),
    numero_inscricao VARCHAR(15),
    ativo BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(empresa_id, codigo)
);
