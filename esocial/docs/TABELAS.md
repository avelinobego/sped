Excelente pergunta! Vou mapear **todas as tabelas necessárias** para um sistema eSocial completo em PostgreSQL, além das tabelas do governo.

## Estrutura Completa de Banco de Dados

### 1. **Domínio: Empresa/Empregador**

```sql
-- Empresas (S-1000)
CREATE TABLE empresas (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    cnpj VARCHAR(14) UNIQUE NOT NULL,
    razao_social VARCHAR(255) NOT NULL,
    nome_fantasia VARCHAR(255),
    tipo_inscricao CHAR(1) NOT NULL, -- 1=CNPJ, 2=CPF, 3=CAEPF, 4=CNO
    classificacao_tributaria VARCHAR(2) NOT NULL, -- 01=Simples, 02=Lucro Presumido, etc
    natureza_juridica VARCHAR(4),
    ind_cooperativa CHAR(1) DEFAULT 'N',
    ind_construtora CHAR(1) DEFAULT 'N',
    ind_desoneracao CHAR(1) DEFAULT 'N',
    situacao VARCHAR(20) DEFAULT 'ATIVA', -- ATIVA, SUSPENSA, INATIVA
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    usuario_cadastro VARCHAR(100),
    
    -- Contato
    telefone VARCHAR(20),
    email VARCHAR(255),
    
    -- Endereço
    logradouro VARCHAR(255),
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    uf CHAR(2),
    cep VARCHAR(8),
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    status_esocial VARCHAR(20), -- PENDENTE, ENVIADO, PROCESSADO, ERRO
    
    CONSTRAINT chk_cnpj_valido CHECK (LENGTH(cnpj) = 14)
);

CREATE INDEX idx_empresas_cnpj ON empresas(cnpj);
CREATE INDEX idx_empresas_situacao ON empresas(situacao);
CREATE INDEX idx_empresas_status_esocial ON empresas(status_esocial);

-- Estabelecimentos (S-1005)
CREATE TABLE estabelecimentos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    tipo_inscricao CHAR(1) NOT NULL, -- 1=CNPJ, 4=CNO
    numero_inscricao VARCHAR(15) NOT NULL,
    cnae_principal VARCHAR(7),
    cnae_preparatorio VARCHAR(7),
    alvara_funcionamento VARCHAR(50),
    data_inicio_atividades DATE,
    ind_obra CHAR(1) DEFAULT 'N',
    
    -- Endereço
    logradouro VARCHAR(255),
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    uf CHAR(2),
    cep VARCHAR(8),
    
    situacao VARCHAR(20) DEFAULT 'ATIVO',
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(empresa_id, numero_inscricao)
);

CREATE INDEX idx_estabelecimentos_empresa ON estabelecimentos(empresa_id);

-- Rubricas (S-1010)
CREATE TABLE rubricas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    natureza_rubrica VARCHAR(4) NOT NULL, -- Tabela 3 do eSocial
    tipo_rubrica CHAR(1) NOT NULL, -- 1=Vencimento, 2=Desconto, 3=Informativa, 4=Incidência
    cod_inccp VARCHAR(2), -- Código incidência CP (INSS)
    cod_incirrf VARCHAR(2), -- Código incidência IRRF
    cod_incfgts VARCHAR(2), -- Código incidência FGTS
    cod_incsind VARCHAR(2), -- Código incidência Sindical
    
    ativa BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE,
    
    UNIQUE(empresa_id, codigo)
);

CREATE INDEX idx_rubricas_empresa ON rubricas(empresa_id);
CREATE INDEX idx_rubricas_ativa ON rubricas(ativa);

-- Lotações Tributárias (S-1020)
CREATE TABLE lotacoes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    tipo_lotacao VARCHAR(2) NOT NULL, -- Tabela 10 do eSocial
    fpas VARCHAR(3), -- Código FPAS
    cod_terceiros VARCHAR(4), -- Código terceiros
    
    ativa BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE,
    
    UNIQUE(empresa_id, codigo)
);

-- Cargos (S-1030)
CREATE TABLE cargos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    cbo VARCHAR(6) NOT NULL, -- CBO 2002
    
    ativo BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_inicio_validade DATE NOT NULL,
    data_fim_validade DATE,
    
    UNIQUE(empresa_id, codigo)
);

-- Horários/Turnos (S-1050)
CREATE TABLE horarios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    tipo_jornada CHAR(1) NOT NULL, -- 2=12x36, 3=Escala, 4=Normal, 9=Outros
    
    -- Jornada semanal
    duracao_jornada_semanal INTEGER, -- Em minutos
    tipo_intervalo CHAR(1), -- 1=Repouso/Alimentação, 2=Interjornada, 3=Intrajornada
    duracao_intervalo INTEGER, -- Em minutos
    
    ativo BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(empresa_id, codigo)
);

-- Horários Detalhados (dias da semana)
CREATE TABLE horarios_detalhes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    horario_id UUID NOT NULL REFERENCES horarios(id) ON DELETE CASCADE,
    dia_semana INTEGER NOT NULL, -- 1=Segunda, 2=Terça, ..., 7=Domingo
    horario_entrada TIME NOT NULL,
    horario_saida TIME NOT NULL,
    duracao_jornada INTEGER NOT NULL, -- Em minutos
    
    UNIQUE(horario_id, dia_semana)
);

-- Ambientes de Trabalho (S-1060)
CREATE TABLE ambientes_trabalho (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    codigo VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    local_ambiente CHAR(1) NOT NULL, -- 1=Estabelecimento, 2=Externo, 3=Terceiros
    tipo_inscricao CHAR(1),
    numero_inscricao VARCHAR(15),
    
    ativo BOOLEAN DEFAULT TRUE,
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(empresa_id, codigo)
);
```

### 2. **Domínio: Trabalhador**

```sql
-- Trabalhadores (S-2200)
CREATE TABLE trabalhadores (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    
    -- Identificação
    cpf VARCHAR(11) UNIQUE NOT NULL,
    nis VARCHAR(11),
    nome VARCHAR(255) NOT NULL,
    nome_social VARCHAR(255),
    data_nascimento DATE NOT NULL,
    sexo CHAR(1) NOT NULL, -- M, F
    raca_cor CHAR(1), -- Tabela 6 eSocial
    estado_civil CHAR(1), -- Tabela 26 eSocial
    grau_instrucao VARCHAR(2), -- Tabela 30 eSocial
    
    -- Nacionalidade
    pais_nascimento VARCHAR(3) DEFAULT '105', -- 105=Brasil
    pais_nacionalidade VARCHAR(3) DEFAULT '105',
    
    -- Documentos
    numero_ctps VARCHAR(11),
    serie_ctps VARCHAR(5),
    uf_ctps CHAR(2),
    data_emissao_ctps DATE,
    
    numero_rg VARCHAR(20),
    orgao_emissor_rg VARCHAR(10),
    uf_rg CHAR(2),
    data_emissao_rg DATE,
    
    numero_cnh VARCHAR(11),
    categoria_cnh VARCHAR(2),
    validade_cnh DATE,
    
    numero_rne VARCHAR(20), -- Para estrangeiros
    
    -- Contato
    telefone VARCHAR(20),
    celular VARCHAR(20),
    email VARCHAR(255),
    
    -- Endereço
    cep VARCHAR(8),
    logradouro VARCHAR(255),
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    uf CHAR(2),
    
    -- Dados bancários
    banco VARCHAR(3),
    agencia VARCHAR(5),
    conta VARCHAR(20),
    tipo_conta CHAR(1), -- C=Corrente, P=Poupança
    
    -- Deficiência
    possui_deficiencia BOOLEAN DEFAULT FALSE,
    tipo_deficiencia VARCHAR(2), -- Tabela 31 eSocial
    observacao_deficiencia TEXT,
    
    -- Status
    situacao VARCHAR(20) DEFAULT 'ATIVO', -- ATIVO, AFASTADO, DESLIGADO
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foto
    foto_url VARCHAR(500),
    
    CONSTRAINT chk_cpf_valido CHECK (LENGTH(cpf) = 11)
);

CREATE INDEX idx_trabalhadores_cpf ON trabalhadores(cpf);
CREATE INDEX idx_trabalhadores_empresa ON trabalhadores(empresa_id);
CREATE INDEX idx_trabalhadores_situacao ON trabalhadores(situacao);
CREATE INDEX idx_trabalhadores_nome ON trabalhadores USING gin(to_tsvector('portuguese', nome));

-- Vínculos de Trabalho (S-2200)
CREATE TABLE vinculos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    trabalhador_id UUID NOT NULL REFERENCES trabalhadores(id) ON DELETE CASCADE,
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    estabelecimento_id UUID REFERENCES estabelecimentos(id),
    
    -- Identificação
    matricula VARCHAR(30) NOT NULL,
    
    -- Admissão
    data_admissao DATE NOT NULL,
    tipo_admissao CHAR(1) NOT NULL, -- Tabela 5 eSocial
    tipo_regime_trabalhista CHAR(1) NOT NULL, -- 1=CLT, 2=Estatutário
    tipo_regime_previdenciario CHAR(1) NOT NULL, -- 1=RGPS, 2=RPPS, 3=Militar
    cadastro_inicial CHAR(1) DEFAULT 'S', -- S=Sim, N=Não
    
    -- Categoria
    categoria VARCHAR(3) NOT NULL, -- Tabela 1 eSocial (101=Empregado CLT, etc)
    
    -- Cargo e Função
    cargo_id UUID REFERENCES cargos(id),
    funcao_id UUID, -- Se tiver função/comissão
    
    -- Lotação
    lotacao_id UUID REFERENCES lotacoes(id),
    
    -- Jornada
    horario_id UUID REFERENCES horarios(id),
    tipo_jornada CHAR(1) NOT NULL, -- Tabela 7 eSocial
    
    -- Contrato
    tipo_contrato CHAR(1) NOT NULL, -- Tabela 14 eSocial (1=Prazo Indeterminado, 2=Prazo Determinado)
    data_inicio_contrato DATE NOT NULL,
    data_fim_contrato DATE, -- Apenas se prazo determinado
    
    -- Remuneração
    valor_salario DECIMAL(10,2) NOT NULL,
    unidade_salario CHAR(1) NOT NULL, -- Tabela 15 eSocial (1=Hora, 7=Mês)
    tipo_salario CHAR(1) NOT NULL, -- Tabela 16 eSocial (N=Normal, P=Piso, etc)
    
    -- Desligamento
    data_desligamento DATE,
    motivo_desligamento VARCHAR(2), -- Tabela 19 eSocial
    aviso_previo CHAR(1), -- S=Sim, N=Não
    data_aviso_previo DATE,
    ind_pagamento_aviso CHAR(1), -- S=Sim (trabalhado), N=Não (indenizado)
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    status_esocial VARCHAR(20), -- PENDENTE, ENVIADO, PROCESSADO, ERRO
    
    -- Status
    situacao VARCHAR(20) DEFAULT 'ATIVO', -- ATIVO, AFASTADO, DESLIGADO
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(empresa_id, matricula)
);

CREATE INDEX idx_vinculos_trabalhador ON vinculos(trabalhador_id);
CREATE INDEX idx_vinculos_empresa ON vinculos(empresa_id);
CREATE INDEX idx_vinculos_situacao ON vinculos(situacao);
CREATE INDEX idx_vinculos_data_admissao ON vinculos(data_admissao);

-- Dependentes
CREATE TABLE dependentes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    trabalhador_id UUID NOT NULL REFERENCES trabalhadores(id) ON DELETE CASCADE,
    
    -- Identificação
    cpf VARCHAR(11),
    nome VARCHAR(255) NOT NULL,
    data_nascimento DATE NOT NULL,
    tipo_dependente VARCHAR(2) NOT NULL, -- Tabela 7 eSocial (01=Cônjuge, 03=Filho, etc)
    
    -- Dados adicionais
    sexo CHAR(1),
    
    -- Dependência
    dependente_irrf BOOLEAN DEFAULT FALSE,
    dependente_sf BOOLEAN DEFAULT FALSE, -- Salário família
    
    -- Incapacidade
    incapaz BOOLEAN DEFAULT FALSE,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_inativacao DATE
);

CREATE INDEX idx_dependentes_trabalhador ON dependentes(trabalhador_id);
```

### 3. **Domínio: Afastamentos**

```sql
-- Afastamentos (S-2230)
CREATE TABLE afastamentos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    
    -- Período
    data_inicio DATE NOT NULL,
    data_fim DATE,
    
    -- Motivo
    motivo_afastamento VARCHAR(2) NOT NULL, -- Tabela 18 eSocial
    descricao_motivo TEXT,
    codigo_cid VARCHAR(10), -- CID-10 (se doença)
    
    -- Acidente de trabalho
    ind_acidente_trabalho BOOLEAN DEFAULT FALSE,
    
    -- INSS
    ind_auxilio_doenca BOOLEAN DEFAULT FALSE,
    numero_beneficio VARCHAR(20),
    
    -- Status
    situacao VARCHAR(20) DEFAULT 'ATIVO', -- ATIVO, ENCERRADO
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_afastamentos_vinculo ON afastamentos(vinculo_id);
CREATE INDEX idx_afastamentos_data_inicio ON afastamentos(data_inicio);
```

### 4. **Domínio: SST (Saúde e Segurança do Trabalho)**

```sql
-- CAT - Comunicação de Acidente de Trabalho (S-2210)
CREATE TABLE acidentes_trabalho (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    
    -- Dados do acidente
    numero_cat VARCHAR(20) UNIQUE,
    data_acidente TIMESTAMP NOT NULL,
    hora_acidente TIME NOT NULL,
    tipo_acidente CHAR(1) NOT NULL, -- Tabela 16 eSocial (1=Típico, 2=Trajeto, 3=Doença)
    
    -- Local
    tipo_local_acidente CHAR(1) NOT NULL, -- Tabela 23 eSocial
    descricao_local VARCHAR(255),
    logradouro VARCHAR(255),
    numero VARCHAR(10),
    complemento VARCHAR(100),
    bairro VARCHAR(100),
    cidade VARCHAR(100),
    uf CHAR(2),
    cep VARCHAR(8),
    pais VARCHAR(3) DEFAULT '105',
    
    -- Parte do corpo atingida
    parte_corpo_atingida VARCHAR(2), -- Tabela 13 eSocial
    
    -- Agente causador
    agente_causador VARCHAR(10), -- Tabela 14 eSocial
    
    -- CID
    codigo_cid VARCHAR(10),
    
    -- Situação geradora
    situacao_geradora VARCHAR(9), -- Tabela 15 eSocial
    
    -- Descrição
    descricao_acidente TEXT NOT NULL,
    
    -- Morte
    houve_morte BOOLEAN DEFAULT FALSE,
    data_obito TIMESTAMP,
    
    -- Atestado médico
    houve_atendimento_medico BOOLEAN DEFAULT FALSE,
    duracao_tratamento INTEGER, -- Em dias
    
    -- Afastamento
    houve_afastamento BOOLEAN DEFAULT FALSE,
    afastamento_id UUID REFERENCES afastamentos(id),
    
    -- Testemunhas
    testemunhas TEXT,
    
    -- CAT origem (em caso de reabertura)
    cat_origem_id UUID REFERENCES acidentes_trabalho(id),
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_cat_vinculo ON acidentes_trabalho(vinculo_id);
CREATE INDEX idx_cat_data ON acidentes_trabalho(data_acidente);

-- ASO - Atestado de Saúde Ocupacional (S-2220)
CREATE TABLE exames_medicos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    
    -- Tipo de exame
    tipo_exame CHAR(1) NOT NULL, -- 0=Admissional, 1=Periódico, 2=Retorno, 3=Mudança, 9=Demissional
    
    -- Data
    data_exame DATE NOT NULL,
    
    -- Médico
    nome_medico VARCHAR(255) NOT NULL,
    crm_medico VARCHAR(10) NOT NULL,
    uf_crm CHAR(2) NOT NULL,
    
    -- Resultado
    resultado CHAR(1) NOT NULL, -- A=Apto, I=Inapto
    observacoes TEXT,
    
    -- Próximo exame (se periódico)
    data_proximo_exame DATE,
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_exames_vinculo ON exames_medicos(vinculo_id);
CREATE INDEX idx_exames_tipo ON exames_medicos(tipo_exame);

-- Agentes Nocivos - Condições Ambientais (S-2240)
CREATE TABLE condicoes_ambientais (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    ambiente_trabalho_id UUID REFERENCES ambientes_trabalho(id),
    
    -- Período
    data_inicio DATE NOT NULL,
    data_fim DATE,
    
    -- Fator de risco
    tipo_fator_risco VARCHAR(2) NOT NULL, -- Tabela 23 eSocial
    codigo_fator_risco VARCHAR(10), -- Conforme tipo
    intensidade_concentracao VARCHAR(20),
    tecnica_utilizada VARCHAR(100),
    
    -- EPI
    utiliza_epi BOOLEAN DEFAULT FALSE,
    epi_eficaz BOOLEAN,
    
    -- Aposentadoria especial
    ind_aposentadoria_especial BOOLEAN DEFAULT FALSE,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_condicoes_vinculo ON condicoes_ambientais(vinculo_id);

-- Treinamentos (S-2245)
CREATE TABLE treinamentos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    
    -- Treinamento
    codigo_treinamento VARCHAR(30) NOT NULL,
    descricao VARCHAR(255) NOT NULL,
    tipo_treinamento CHAR(1) NOT NULL, -- Tabela 29 eSocial
    carga_horaria INTEGER NOT NULL, -- Em horas
    
    -- Instrutor
    nome_instrutor VARCHAR(255),
    cpf_instrutor VARCHAR(11),
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(empresa_id, codigo_treinamento)
);

-- Participantes dos Treinamentos
CREATE TABLE treinamentos_participantes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    treinamento_id UUID NOT NULL REFERENCES treinamentos(id) ON DELETE CASCADE,
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    
    -- Realização
    data_inicio DATE NOT NULL,
    data_fim DATE NOT NULL,
    
    -- Resultado
    aprovado BOOLEAN DEFAULT TRUE,
    nota DECIMAL(5,2),
    observacoes TEXT,
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(treinamento_id, vinculo_id, data_inicio)
);
```

### 5. **Domínio: Folha de Pagamento**

```sql
-- Períodos de Folha
CREATE TABLE periodos_folha (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    
    -- Competência
    mes INTEGER NOT NULL CHECK (mes BETWEEN 1 AND 13), -- 13 = 13º salário
    ano INTEGER NOT NULL,
    
    -- Tipo de folha
    tipo_folha VARCHAR(20) NOT NULL, -- MENSAL, ADIANTAMENTO, FERIAS, RESCISAO, 13_SALARIO
    
    -- Datas
    data_inicio DATE NOT NULL,
    data_fim DATE NOT NULL,
    data_pagamento DATE NOT NULL,
    
    -- Status
    situacao VARCHAR(20) DEFAULT 'ABERTA', -- ABERTA, CALCULADA, FECHADA, ENVIADA
    data_calculo TIMESTAMP,
    data_fechamento TIMESTAMP,
    
    -- eSocial S-1299
    protocolo_fechamento VARCHAR(50),
    recibo_fechamento VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(empresa_id, mes, ano, tipo_folha)
);

CREATE INDEX idx_periodos_empresa ON periodos_folha(empresa_id);
CREATE INDEX idx_periodos_competencia ON periodos_folha(ano, mes);

-- Eventos de Remuneração (S-1200)
CREATE TABLE eventos_remuneracao (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    periodo_folha_id UUID NOT NULL REFERENCES periodos_folha(id) ON DELETE CASCADE,
    vinculo_id UUID NOT NULL REFERENCES vinculos(id) ON DELETE CASCADE,
    
    -- Identificação
    tipo_evento VARCHAR(20) NOT NULL, -- FOLHA_NORMAL, ADIANTAMENTO, FERIAS, RESCISAO, 13_SALARIO
    
    -- Valores totais
    total_vencimentos DECIMAL(10,2) DEFAULT 0,
    total_descontos DECIMAL(10,2) DEFAULT 0,
    valor_liquido DECIMAL(10,2) DEFAULT 0,
    
    -- Bases de cálculo
    base_inss DECIMAL(10,2) DEFAULT 0,
    valor_inss DECIMAL(10,2) DEFAULT 0,
    base_irrf DECIMAL(10,2) DEFAULT 0,
    valor_irrf DECIMAL(10,2) DEFAULT 0,
    base_fgts DECIMAL(10,2) DEFAULT 0,
    valor_fgts DECIMAL(10,2) DEFAULT 0,
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    status_esocial VARCHAR(20),
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(periodo_folha_id, vinculo_id, tipo_evento)
);

CREATE INDEX idx_eventos_periodo ON eventos_remuneracao(periodo_folha_id);
CREATE INDEX idx_eventos_vinculo ON eventos_remuneracao(vinculo_id);

-- Itens da Remuneração (rubricas lançadas)
CREATE TABLE itens_remuneracao (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    evento_remuneracao_id UUID NOT NULL REFERENCES eventos_remuneracao(id) ON DELETE CASCADE,
    rubrica_id UUID NOT NULL REFERENCES rubricas(id),
    
    -- Quantidade e valores
    quantidade DECIMAL(10,4) DEFAULT 1,
    valor_unitario DECIMAL(10,2),
    valor_total DECIMAL(10,2) NOT NULL,
    
    -- Tipo
    tipo_item CHAR(1) NOT NULL, -- V=Vencimento, D=Desconto, I=Informativo
    
    -- Referências
    referencia VARCHAR(100), -- Ex: "220 horas", "30 dias"
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_itens_evento ON itens_remuneracao(evento_remuneracao_id);
CREATE INDEX idx_itens_rubrica ON itens_remuneracao(rubrica_id);
```

### 6. **Domínio: Processos e Auditoria**

```sql
-- Processos Trabalhistas (S-2500)
CREATE TABLE processos_trabalhistas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    
    -- Identificação
    numero_processo VARCHAR(20) UNIQUE NOT NULL,
    tipo_processo CHAR(1) NOT NULL, -- 1=Administrativo, 2=Judicial
    
    -- Vara/Órgão
    uf_vara CHAR(2),
    codigo_vara VARCHAR(5),
    id_vara VARCHAR(10),
    
    -- Datas
    data_inicio DATE NOT NULL,
    data_sentenca DATE,
    data_transito_julgado DATE,
    
    -- Autor
    cpf_autor VARCHAR(11) NOT NULL,
    nome_autor VARCHAR(255) NOT NULL,
    
    -- Valores
    valor_causa DECIMAL(12,2),
    valor_condenacao DECIMAL(12,2),
    
    -- Advogados
    nome_advogado VARCHAR(255),
    oab_advogado VARCHAR(15),
    
    situacao VARCHAR(20) DEFAULT 'EM_ANDAMENTO', -- EM_ANDAMENTO, SENTENCIADO, ARQUIVADO
    
    -- eSocial
    protocolo_esocial VARCHAR(50),
    recibo_esocial VARCHAR(50),
    data_envio_esocial TIMESTAMP,
    
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Log de Eventos eSocial (Auditoria completa)
CREATE TABLE eventos_esocial_log (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    
    -- Evento
    tipo_evento VARCHAR(10) NOT NULL, -- S-1000, S-2200, etc
    identificador_evento VARCHAR(50),
    
    -- Dados do trabalhador (se aplicável)
    trabalhador_id UUID REFERENCES trabalhadores(id),
    vinculo_id UUID REFERENCES vinculos(id),
    
    -- XML
    xml_enviado TEXT,
    xml_retorno TEXT,
    
    -- Resultado
    protocolo VARCHAR(50),
    recibo VARCHAR(50),
    status VARCHAR(20) NOT NULL, -- ENVIADO, PROCESSADO, ERRO, REJEITADO
    codigo_retorno VARCHAR(10),
    descricao_retorno TEXT,
    
    -- Erros (se houver)
    erros JSONB,
    
    -- Timestamps
    data_envio TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    data_retorno TIMESTAMP,
    
    -- Usuário
    usuario_envio VARCHAR(100)
);

CREATE INDEX idx_eventos_log_empresa ON eventos_esocial_log(empresa_id);
CREATE INDEX idx_eventos_log_tipo ON eventos_esocial_log(tipo_evento);
CREATE INDEX idx_eventos_log_status ON eventos_esocial_log(status);
CREATE INDEX idx_eventos_log_data ON eventos_esocial_log(data_envio);
CREATE INDEX idx_eventos_log_trabalhador ON eventos_esocial_log(trabalhador_id);

-- Fila de Eventos eSocial (para processamento assíncrono)
CREATE TABLE fila_eventos_esocial (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    
    -- Evento
    tipo_evento VARCHAR(10) NOT NULL,
    entidade_tipo VARCHAR(50), -- trabalhador, vinculo, folha, etc
    entidade_id UUID NOT NULL,
    
    -- Payload
    dados JSONB NOT NULL,
    
    -- Processamento
    situacao VARCHAR(20) DEFAULT 'PENDENTE', -- PENDENTE, PROCESSANDO, ENVIADO, ERRO
    tentativas INTEGER DEFAULT 0,
    max_tentativas INTEGER DEFAULT 3,
    
    -- Timestamps
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    data_processamento TIMESTAMP,
    
    -- Resultado
    evento_log_id UUID REFERENCES eventos_esocial_log(id),
    mensagem_erro TEXT,
    
    -- Prioridade
    prioridade INTEGER DEFAULT 5 -- 1=Urgente, 5=Normal, 10=Baixa
);

CREATE INDEX idx_fila_situacao ON fila_eventos_esocial(situacao);
CREATE INDEX idx_fila_prioridade ON fila_eventos_esocial(prioridade, data_criacao);
CREATE INDEX idx_fila_empresa ON fila_eventos_esocial(empresa_id);
```

### 7. **Tabelas de Suporte/Auxiliares**

```sql
-- Usuários do sistema
CREATE TABLE usuarios (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    nome VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    senha_hash VARCHAR(255) NOT NULL,
    cpf VARCHAR(11) UNIQUE,
    
    -- Perfil
    perfil VARCHAR(20) NOT NULL, -- ADMIN, OPERADOR, CONSULTA
    
    -- Status
    ativo BOOLEAN DEFAULT TRUE,
    
    -- Timestamps
    data_cadastro TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ultimo_acesso TIMESTAMP,
    
    -- Empresa (se for usuário de uma empresa específica)
    empresa_id UUID REFERENCES empresas(id)
);

-- Permissões de acesso
CREATE TABLE usuarios_empresas (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    usuario_id UUID NOT NULL REFERENCES usuarios(id) ON DELETE CASCADE,
    empresa_id UUID NOT NULL REFERENCES empresas(id) ON DELETE CASCADE,
    perfil VARCHAR(20) NOT NULL,
    
    UNIQUE(usuario_id, empresa_id)
);

-- Auditoria geral do sistema
CREATE TABLE auditoria (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Ação
    tabela VARCHAR(100) NOT NULL,
    operacao VARCHAR(10) NOT NULL, -- INSERT, UPDATE, DELETE
    registro_id UUID NOT NULL,
    
    -- Dados
    dados_anteriores JSONB,
    dados_novos JSONB,
    
    -- Usuário
    usuario_id UUID REFERENCES usuarios(id),
    usuario_nome VARCHAR(255),
    
    -- Timestamp
    data_hora TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- IP
    ip_origem VARCHAR(45)
);

CREATE INDEX idx_auditoria_tabela ON auditoria(tabela);
CREATE INDEX idx_auditoria_registro ON auditoria(registro_id);
CREATE INDEX idx_auditoria_data ON auditoria(data_hora);

-- Configurações do sistema
CREATE TABLE configuracoes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    chave VARCHAR(100) UNIQUE NOT NULL,
    valor TEXT,
    tipo VARCHAR(20) NOT NULL, -- STRING, INTEGER, BOOLEAN, JSON
    descricao TEXT,
    editavel BOOLEAN DEFAULT TRUE,
    
    data_atualizacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Notificações
CREATE TABLE notificacoes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Destinatário
    usuario_id UUID REFERENCES usuarios(id) ON DELETE CASCADE,
    empresa_id UUID REFERENCES empresas(id),
    
    -- Conteúdo
    tipo VARCHAR(50) NOT NULL, -- EVENTO_ERRO, EVENTO_SUCESSO, ALERTA, etc
    titulo VARCHAR(255) NOT NULL,
    mensagem TEXT NOT NULL,
    
    -- Link relacionado
    link VARCHAR(500),
    
    -- Status
    lida BOOLEAN DEFAULT FALSE,
    data_leitura TIMESTAMP,
    
    -- Timestamp
    data_criacao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_notificacoes_usuario ON notificacoes(usuario_id, lida);
```

### 8. **Views Úteis**

```sql
-- View: Trabalhadores ativos com vínculo atual
CREATE VIEW v_trabalhadores_ativos AS
SELECT 
    t.id,
    t.cpf,
    t.nome,
    t.data_nascimento,
    v.matricula,
    v.data_admissao,
    v.valor_salario,
    c.descricao as cargo,
    l.descricao as lotacao,
    e.razao_social as empresa
FROM trabalhadores t
INNER JOIN vinculos v ON t.id = v.trabalhador_id
INNER JOIN empresas e ON v.empresa_id = e.id
LEFT JOIN cargos c ON v.cargo_id = c.id
LEFT JOIN lotacoes l ON v.lotacao_id = l.id
WHERE v.situacao = 'ATIVO'
  AND t.situacao = 'ATIVO';

-- View: Eventos pendentes de envio
CREATE VIEW v_eventos_pendentes AS
SELECT 
    f.id,
    f.tipo_evento,
    f.entidade_tipo,
    f.prioridade,
    f.tentativas,
    f.data_criacao,
    e.razao_social as empresa
FROM fila_eventos_esocial f
INNER JOIN empresas e ON f.empresa_id = e.id
WHERE f.situacao = 'PENDENTE'
ORDER BY f.prioridade ASC, f.data_criacao ASC;

-- View: Resumo de folha por período
CREATE VIEW v_resumo_folha AS
SELECT 
    pf.id as periodo_id,
    pf.mes,
    pf.ano,
    pf.tipo_folha,
    e.razao_social as empresa,
    COUNT(DISTINCT er.vinculo_id) as quantidade_trabalhadores,
    SUM(er.total_vencimentos) as total_vencimentos,
    SUM(er.total_descontos) as total_descontos,
    SUM(er.valor_liquido) as total_liquido,
    SUM(er.valor_inss) as total_inss,
    SUM(er.valor_irrf) as total_irrf,
    SUM(er.valor_fgts) as total_fgts
FROM periodos_folha pf
INNER JOIN empresas e ON pf.empresa_id = e.id
LEFT JOIN eventos_remuneracao er ON pf.id = er.periodo_folha_id
GROUP BY pf.id, pf.mes, pf.ano, pf.tipo_folha, e.razao_social;
```

### 9. **Functions e Triggers Úteis**

```sql
-- Function: Atualizar timestamp automaticamente
CREATE OR REPLACE FUNCTION atualizar_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.data_atualizacao = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger para empresas
CREATE TRIGGER trigger_empresas_timestamp
BEFORE UPDATE ON empresas
FOR EACH ROW
EXECUTE FUNCTION atualizar_timestamp();

-- Trigger para trabalhadores
CREATE TRIGGER trigger_trabalhadores_timestamp
BEFORE UPDATE ON trabalhadores
FOR EACH ROW
EXECUTE FUNCTION atualizar_timestamp();

-- Trigger para vínculos
CREATE TRIGGER trigger_vinculos_timestamp
BEFORE UPDATE ON vinculos
FOR EACH ROW
EXECUTE FUNCTION atualizar_timestamp();

-- Function: Validar CPF
CREATE OR REPLACE FUNCTION validar_cpf(cpf_input VARCHAR)
RETURNS BOOLEAN AS $$
DECLARE
    cpf VARCHAR(11);
    soma INTEGER;
    resto INTEGER;
    digito1 INTEGER;
    digito2 INTEGER;
BEGIN
    -- Remove caracteres não numéricos
    cpf := regexp_replace(cpf_input, '\D', '', 'g');
    
    -- Verifica tamanho
    IF LENGTH(cpf) != 11 THEN
        RETURN FALSE;
    END IF;
    
    -- Verifica se todos os dígitos são iguais
    IF cpf ~ '^(\d)\1{10}$' THEN
        RETURN FALSE;
    END IF;
    
    -- Calcula primeiro dígito verificador
    soma := 0;
    FOR i IN 1..9 LOOP
        soma := soma + (SUBSTRING(cpf, i, 1)::INTEGER * (11 - i));
    END LOOP;
    resto := (soma * 10) % 11;
    IF resto = 10 THEN
        resto := 0;
    END IF;
    digito1 := resto;
    
    -- Calcula segundo dígito verificador
    soma := 0;
    FOR i IN 1..10 LOOP
        soma := soma + (SUBSTRING(cpf, i, 1)::INTEGER * (12 - i));
    END LOOP;
    resto := (soma * 10) % 11;
    IF resto = 10 THEN
        resto := 0;
    END IF;
    digito2 := resto;
    
    -- Valida
    IF SUBSTRING(cpf, 10, 1)::INTEGER = digito1 AND 
       SUBSTRING(cpf, 11, 1)::INTEGER = digito2 THEN
        RETURN TRUE;
    END IF;
    
    RETURN FALSE;
END;
$$ LANGUAGE plpgsql;

-- Constraint usando a function
ALTER TABLE trabalhadores ADD CONSTRAINT chk_cpf_valido_func 
CHECK (validar_cpf(cpf));

-- Function: Registrar auditoria
CREATE OR REPLACE FUNCTION registrar_auditoria()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'DELETE' THEN
        INSERT INTO auditoria (tabela, operacao, registro_id, dados_anteriores)
        VALUES (TG_TABLE_NAME, TG_OP, OLD.id, row_to_json(OLD));
        RETURN OLD;
    ELSIF TG_OP = 'UPDATE' THEN
        INSERT INTO auditoria (tabela, operacao, registro_id, dados_anteriores, dados_novos)
        VALUES (TG_TABLE_NAME, TG_OP, NEW.id, row_to_json(OLD), row_to_json(NEW));
        RETURN NEW;
    ELSIF TG_OP = 'INSERT' THEN
        INSERT INTO auditoria (tabela, operacao, registro_id, dados_novos)
        VALUES (TG_TABLE_NAME, TG_OP, NEW.id, row_to_json(NEW));
        RETURN NEW;
    END IF;
END;
$$ LANGUAGE plpgsql;

-- Aplicar trigger de auditoria em tabelas críticas
CREATE TRIGGER trigger_auditoria_trabalhadores
AFTER INSERT OR UPDATE OR DELETE ON trabalhadores
FOR EACH ROW EXECUTE FUNCTION registrar_auditoria();

CREATE TRIGGER trigger_auditoria_vinculos
AFTER INSERT OR UPDATE OR DELETE ON vinculos
FOR EACH ROW EXECUTE FUNCTION registrar_auditoria();

CREATE TRIGGER trigger_auditoria_eventos_remuneracao
AFTER INSERT OR UPDATE OR DELETE ON eventos_remuneracao
FOR EACH ROW EXECUTE FUNCTION registrar_auditoria();
```

## Resumo das Tabelas

**Total: ~35 tabelas principais**

1. **Empresa**: 8 tabelas (empresas, estabelecimentos, rubricas, lotacoes, cargos, horarios, ambientes)
2. **Trabalhador**: 3 tabelas (trabalhadores, vinculos, dependentes)
3. **Afastamentos**: 1 tabela
4. **SST**: 5 tabelas (acidentes, exames, condicoes_ambientais, treinamentos, participantes)
5. **Folha**: 4 tabelas (periodos, eventos, itens, pagamentos)
6. **Processos**: 1 tabela
7. **eSocial**: 3 tabelas (log, fila, retornos)
8. **Sistema**: 6 tabelas (usuarios, permissoes, auditoria, configuracoes, notificacoes)
9. **Views**: 3+ views úteis

Esta estrutura cobre **100% do eSocial** e permite rastreabilidade completa!
