CREATE TABLE empresas (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    cnpj VARCHAR(14) UNIQUE NOT NULL,
    razao_social VARCHAR(255) NOT NULL,
    nome_fantasia VARCHAR(255),
    tipo_inscricao CHAR(1) NOT NULL,
    classificacao_tributaria VARCHAR(2) NOT NULL,
    natureza_juridica VARCHAR(4),
    ind_cooperativa CHAR(1) DEFAULT 'N',
    ind_construtora CHAR(1) DEFAULT 'N',
    ind_desoneracao CHAR(1) DEFAULT 'N',
    situacao VARCHAR(20) DEFAULT 'ATIVA',
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    usuario_cadastro VARCHAR(100),
    telefone VARCHAR(20),
    email VARCHAR(255),
    logradouro VARCHAR(255),
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    uf CHAR(2),
    cep VARCHAR(8),
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP WITH TIME ZONE,
    status_esocial VARCHAR(20),
    CONSTRAINT chk_cnpj_valido CHECK (LENGTH(cnpj) = 14)
);

CREATE TABLE estabelecimentos (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    tipo_inscricao CHAR(1) NOT NULL,
    numero_inscricao VARCHAR(15) NOT NULL,
    cnae_principal VARCHAR(7),
    cnae_preparatorio VARCHAR(7),
    alvara_funcionamento VARCHAR(50),
    data_inicio_atividades DATE,
    ind_obra CHAR(1) DEFAULT 'N',
    logradouro VARCHAR(255),
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    uf CHAR(2),
    cep VARCHAR(8),
    situacao VARCHAR(20) DEFAULT 'ATIVO',
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(empresa_id, numero_inscricao)
);

CREATE TABLE rubricas (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    natureza_rubrica VARCHAR(4) NOT NULL,
    tipo_rubrica CHAR(1) NOT NULL,
    cod_inccp VARCHAR(2),
    cod_incirrf VARCHAR(2),
    cod_incfgts VARCHAR(2),
    cod_incsind VARCHAR(2),
    ativa BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE,
    UNIQUE(empresa_id, codigo)
);

CREATE TABLE lotacoes (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    tipo_lotacao VARCHAR(2) NOT NULL,
    fpas VARCHAR(3),
    cod_terceiros VARCHAR(4),
    ativa BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE,
    UNIQUE(empresa_id, codigo)
);

CREATE TABLE cargos (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    cbo VARCHAR(6) NOT NULL,
    ativo BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE,
    UNIQUE(empresa_id, codigo)
);
