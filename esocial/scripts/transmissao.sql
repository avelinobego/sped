-- 1. Tabela de Lotes de Transmissão (Agrupador de Eventos - ex: S-1299, S-1000)
CREATE TABLE lote_transmissao (
    id_lote SERIAL PRIMARY KEY,
    empregador_id INT NOT NULL, -- Referência à tabela de empregador
    protocolo_envio VARCHAR(50) UNIQUE, -- Protocolo retornado pelo ambiente do eSocial
    data_hora_envio TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_lote VARCHAR(30) NOT NULL -- 'ENVIADO', 'PROCESSADO', 'ERRO'
);

-- 2. Tabela de Eventos Transmitidos (Detalhe individual de cada evento no lote)
CREATE TABLE evento_transmissao (
    id_evento SERIAL PRIMARY KEY,
    lote_id INT REFERENCES lote_transmissao(id_lote) ON DELETE CASCADE,
    id_evento_esocial VARCHAR(36) UNIQUE NOT NULL, -- ID único gerado no XML do evento (ex: ID1... )
    tipo_evento VARCHAR(10) NOT NULL, -- Ex: 'S-1000', 'S-2200', 'S-1200'
    xml_envio TEXT NOT NULL, -- Conteúdo XML assinado enviado
    status_processamento VARCHAR(30) NOT NULL, -- 'SUCESSO', 'REJEITADO'
    recibo_evento VARCHAR(44), -- Número do recibo oficial retornado pelo eSocial (se sucesso)
    data_hora_processamento TIMESTAMP
);

-- 3. Tabela de Ocorrências / Erros de Retorno (Validações do eSocial)
CREATE TABLE ocorrencia_transmissao (
    id_ocorrencia SERIAL PRIMARY KEY,
    evento_id INT REFERENCES evento_transmissao(id_evento) ON DELETE CASCADE,
    codigo_erro INT NOT NULL, -- Código da crítica/erro do eSocial
    descricao_erro TEXT NOT NULL, -- Mensagem descritiva da inconsistência
    tipo_ocorrencia VARCHAR(10) NOT NULL -- 'ADVERTENCIA', 'ERRO'
);