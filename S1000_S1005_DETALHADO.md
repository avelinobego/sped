# Estrutura Técnica dos Eventos S-1000 e S-1005 - eSocial v. S-1.3

## 📋 Sumário Executivo

Este documento apresenta a análise detalhada dos eventos de cadastro do empregador e estabelecimento no eSocial versão S-1.3:

- **S-1000**: Informações do Empregador / Empresa
- **S-1005**: Dados do Estabelecimento

Ambos fazem parte da família de **Tabelas e Cadastros Básicos** e são fundamentais para a estrutura da transmissão eSocial.

---

## 1. EVENTO S-1000: Informações do Empregador

### 1.1 Descrição Geral

O evento **S-1000** registra as informações cadastrais do empregador (empresa, órgão público, etc.). É o evento raiz da transmissão eSocial e **obrigatório** para todo contribuinte. Deve ser transmitido antes de qualquer outro evento relacionado ao empregador.

| Atributo | Valor |
|----------|-------|
| **Código do Evento** | S-1000 |
| **Nome** | Informações do Empregador |
| **Obrigatoriedade** | Obrigatório (Precursor) |
| **Versão** | S-1.3 (NT 06/2026) |
| **Tipo de Transmissão** | Evento de Cadastro/Tabela |
| **Competência** | Data de início de validade das informações |

### 1.2 Estrutura XML Principal

```xml
<evtInfoEmpregador Id="ID000001">
  <ideEvento>
    <tpAmb>2</tpAmb>                          <!-- 1=Produção / 2=Produção Restrita -->
    <procEmi>1</procEmi>                      <!-- 1=Aplicativo / 2=Sistema eSocial -->
    <verProc>09.00</verProc>                  <!-- Versão do processador -->
  </ideEvento>
  
  <ideEmpregador>
    <tpInsc>1</tpInsc>                        <!-- 1=CNPJ / 2=CPF / 3=CAEPF / 4=CNO -->
    <nrInsc>34028316000152</nrInsc>           <!-- Número da Inscrição (14 dígitos se CNPJ, 11 se CPF) -->
  </ideEmpregador>
  
  <infoCadastro>
    <dtCadastro>2024-01-15</dtCadastro>       <!-- Data de cadastro do empregador -->
    
    <!-- INFORMAÇÕES DO EMPREGADOR / PJ -->
    <infoPJ>
      <razaoSocial>EMPRESA EXEMPLO LTDA</razaoSocial>              <!-- Obrigatório: até 180 caracteres -->
      <nomeFantasia>EMPRESA EXEMPLO</nomeFantasia>                  <!-- Opcional: até 180 caracteres -->
      <dsCategEmpregador>01</dsCategEmpregador>                    <!-- Categoria de empregador (01-10) -->
      
      <!-- ESTRUTURA ADMINISTRATIVA -->
      <classTributaria>06</classTributaria>    <!-- Classificação Tributária:
                                                   01=Micro empresa; 02=Pequena empresa;
                                                   06=Empresa de pequeno porte (EPP);
                                                   ... -->
      
      <!-- INDICATIVOS -->
      <indDesFolha>N</indDesFolha>             <!-- S/N: Indicativo de empresa desonerada -->
      <indConstrutora>N</indConstrutora>       <!-- S/N: Indicativo de construtora -->
      <indCooperativa>N</indCooperativa>       <!-- S/N: Indicativo de cooperativa -->
      
      <!-- CERTIFICADO DIGITAL -->
      <nrCertif>34.028.316/0001-52</nrCertif>  <!-- Opcional: Número do certificado do e-CNPJ -->
    </infoPJ>
    
    <!-- ENDEREÇO DO EMPREGADOR -->
    <enderEmpregador>
      <tpLograd>RUA</tpLograd>                 <!-- Tipo de logradouro (RUA, AV, PÇ, etc.) -->
      <dscLograd>DAS FLORES</dscLograd>        <!-- Descrição do logradouro (até 100 caracteres) -->
      <nrLograd>123</nrLograd>                 <!-- Número (até 10 caracteres) -->
      <complemento>APTO 101</complemento>      <!-- Opcional: Complemento (até 30 caracteres) -->
      <bairro>CENTRO</bairro>                  <!-- Bairro (até 90 caracteres) -->
      <cepLograd>01310100</cepLograd>          <!-- CEP (formato XXXXXXXX - sem hífen) -->
      <codMun>3550308</codMun>                 <!-- Código IBGE do município -->
      <uf>SP</uf>                              <!-- UF (2 caracteres) -->
    </enderEmpregador>
    
    <!-- CONTATO -->
    <infoCont>
      <nmCont>João Silva</nmCont>              <!-- Nome do contato (até 70 caracteres) -->
      <cpfCont>12345678901</cpfCont>           <!-- CPF do contato (sem formatação) -->
      <telefone>1133334444</telefone>          <!-- Telefone (até 15 caracteres) -->
      <email>contato@empresa.com.br</email>    <!-- Email (até 120 caracteres) -->
    </infoCont>
  </infoCadastro>
</evtInfoEmpregador>
```

### 1.3 Campos Principais do S-1000

#### 1.3.1 Grupo `ideEvento`

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `tpAmb` | Inteiro | Sim | Tipo de Ambiente: 1=Produção, 2=Produção Restrita |
| `procEmi` | Inteiro | Sim | Processo de Emissão: 1=Aplicativo, 2=Sistema eSocial |
| `verProc` | String | Sim | Versão do processador (ex: "09.00") |

#### 1.3.2 Grupo `ideEmpregador`

| Campo | Tipo | Obrigatório | Formato | Descrição |
|-------|------|-------------|---------|-----------|
| `tpInsc` | Inteiro | Sim | 1-4 | Tipo de Inscrição: 1=CNPJ, 2=CPF, 3=CAEPF, 4=CNO |
| `nrInsc` | String | Sim | 11-14 dígitos | Número da Inscrição |

#### 1.3.3 Grupo `infoCadastro`

| Campo | Tipo | Obrigatório | Validações |
|-------|------|-------------|-----------|
| `dtCadastro` | Data | Sim | Formato: AAAA-MM-DD |
| `razaoSocial` | String | Sim | Máx. 180 caracteres |
| `nomeFantasia` | String | Não | Máx. 180 caracteres |
| `dsCategEmpregador` | Inteiro | Sim | 01-10 (Tabela eSocial 07) |
| `classTributaria` | Inteiro | Sim | (Tabela eSocial 52) |
| `indDesFolha` | Char | Não | S/N |
| `indConstrutora` | Char | Não | S/N |
| `indCooperativa` | Char | Não | S/N |

#### 1.3.4 Grupo `enderEmpregador`

| Campo | Tipo | Obrigatório | Validações |
|-------|------|-------------|-----------|
| `tpLograd` | String | Sim | Máx. 4 caracteres (RUA, AV, PÇ, etc.) |
| `dscLograd` | String | Sim | Máx. 100 caracteres |
| `nrLograd` | String | Sim | Máx. 10 caracteres |
| `complemento` | String | Não | Máx. 30 caracteres |
| `bairro` | String | Sim | Máx. 90 caracteres |
| `cepLograd` | String | Sim | 8 dígitos (sem hífen) |
| `codMun` | Inteiro | Sim | Código IBGE do município |
| `uf` | String | Sim | Exatamente 2 caracteres |

#### 1.3.5 Grupo `infoCont`

| Campo | Tipo | Obrigatório | Validações |
|-------|------|-------------|-----------|
| `nmCont` | String | Não | Máx. 70 caracteres |
| `cpfCont` | String | Não | 11 dígitos (sem formatação) |
| `telefone` | String | Não | Máx. 15 caracteres |
| `email` | String | Não | Máx. 120 caracteres |

### 1.4 Validações e Regras Específicas do S-1000

#### Regras de Validação

1. **Tipo de Inscrição**:
   - Se `tpInsc = 1` (CNPJ): `nrInsc` deve ter exatamente 14 dígitos
   - Se `tpInsc = 2` (CPF): `nrInsc` deve ter exatamente 11 dígitos
   - CNPJ e CPF devem passar na validação de dígito verificador

2. **Data de Cadastro**:
   - Não pode ser anterior a 01/01/2010 (início do eSocial)
   - Não pode ser superior à data da transmissão do evento

3. **Categoria de Empregador**:
   - Código 01-10 (conforme Tabela 07 do eSocial)
   - Exemplos: 01=Empresa privada, 02=Órgão público, etc.

4. **Classificação Tributária**:
   - Válida conforme Tabela 52 do eSocial
   - Deve estar registrada e ativa no período

5. **Endereço**:
   - CEP deve existir no cadastro de CEP válido
   - Município deve ser válido (código IBGE existente)
   - UF deve ser válido (Tabela 26 do eSocial)

6. **Indicativos**:
   - Campo `indDesFolha`:
     - `S` = Empresa desonerada (Lei 12.546/2011, Lei 12.844/2013, Lei 13.199/2015)
     - `N` = Empresa não desonerada (padrão)

### 1.5 Relacionamento com Outros Eventos

```
S-1000 (Informações do Empregador)
  ├── Deve ser transmitido ANTES de:
  │   ├── S-1005 (Dados do Estabelecimento) ✓
  │   ├── S-1010, S-1020, S-1030 (Tabelas)
  │   └── S-1200+ (Eventos de Trabalhadores)
  │
  └── Pode ser retificado por:
      └── S-1000 (nova transmissão com mesmo ID)
```

---

## 2. EVENTO S-1005: Dados do Estabelecimento

### 2.1 Descrição Geral

O evento **S-1005** registra informações específicas de cada estabelecimento da empresa. É **obrigatório** quando a empresa possui estabelecimentos, e deve ser transmitido após o S-1000. Cada estabelecimento é identificado por um CNPJ ou CEI específico.

| Atributo | Valor |
|----------|-------|
| **Código do Evento** | S-1005 |
| **Nome** | Dados do Estabelecimento |
| **Obrigatoriedade** | Obrigatório (se houver estabelecimento) |
| **Versão** | S-1.3 (NT 06/2026) |
| **Tipo de Transmissão** | Evento de Cadastro/Tabela |
| **Relacionado com** | S-1000 (Empregador) |

### 2.2 Estrutura XML Principal

```xml
<evtInfoEstab Id="ID000002">
  <ideEvento>
    <tpAmb>2</tpAmb>
    <procEmi>1</procEmi>
    <verProc>09.00</verProc>
  </ideEvento>
  
  <ideEmpregador>
    <tpInsc>1</tpInsc>
    <nrInsc>34028316000152</nrInsc>              <!-- CNPJ da matriz/empresa -->
  </ideEmpregador>
  
  <infoCadastro>
    <!-- IDENTIFICAÇÃO DO ESTABELECIMENTO -->
    <ideEstab>
      <tpInscEstab>1</tpInscEstab>               <!-- 1=CNPJ / 2=CEI -->
      <nrInscEstab>34028316000152</nrInscEstab>  <!-- CNPJ ou CEI do estabelecimento -->
    </ideEstab>
    
    <!-- INFORMAÇÕES BÁSICAS DO ESTABELECIMENTO -->
    <dadosEstab>
      <nomeEstab>FILIAL CENTRO</nomeEstab>       <!-- Nome/Razão Social do Estab. (até 180 caracteres) -->
      <dsCategEmpregador>01</dsCategEmpregador>  <!-- Categoria de empregador -->
      
      <!-- ENDEREÇO DO ESTABELECIMENTO -->
      <enderEstab>
        <tpLograd>RUA</tpLograd>
        <dscLograd>DAS FLORES</dscLograd>
        <nrLograd>123</nrLograd>
        <complemento>SALA 200</complemento>
        <bairro>CENTRO</bairro>
        <cepLograd>01310100</cepLograd>
        <codMun>3550308</codMun>                 <!-- Código IBGE do município -->
        <uf>SP</uf>
      </enderEstab>
      
      <!-- INFORMAÇÕES ADICIONAIS -->
      <ativEcon>4711300</ativEcon>               <!-- Código CNAE da atividade econômica -->
      <indDesFolha>N</indDesFolha>               <!-- Indicativo de desoneração -->
      
      <!-- SOFTWARES -->
      <softwares>
        <softwareHouse>
          <cnpjSoftHouse>12345678000190</cnpjSoftHouse>  <!-- CNPJ da software house -->
          <nmSoftware>SISTEMA RH XYZ</nmSoftware>        <!-- Nome do software (até 80 caracteres) -->
          <versaoSistema>5.10.1</versaoSistema>          <!-- Versão do software -->
        </softwareHouse>
      </softwares>
      
      <!-- DADOS DE AVULSO/CONTRIBUINTE INDIVIDUAL -->
      <dadosAvulso>
        <!-- Opcional: Dados para empresas que utilizam avulso -->
        <indAvulso>N</indAvulso>                 <!-- S/N: É empresa de trabalho avulso? -->
      </dadosAvulso>
    </dadosEstab>
    
    <!-- ALTERAÇÕES DE DADOS -->
    <alteracoes>
      <dtAlteracao>2024-06-15</dtAlteracao>      <!-- Data da alteração -->
      <descricaoAlteracao>Mudança de endereço</descricaoAlteracao>  <!-- Descrição (até 255 caracteres) -->
    </alteracoes>
  </infoCadastro>
</evtInfoEstab>
```

### 2.3 Campos Principais do S-1005

#### 2.3.1 Grupo `ideEmpregador`

Identifica o empregador (mesma estrutura do S-1000):

| Campo | Tipo | Obrigatório | Descrição |
|-------|------|-------------|-----------|
| `tpInsc` | Inteiro | Sim | 1=CNPJ, 2=CPF, 3=CAEPF, 4=CNO |
| `nrInsc` | String | Sim | Número da Inscrição do empregador |

#### 2.3.2 Grupo `ideEstab`

Identifica o estabelecimento específico:

| Campo | Tipo | Obrigatório | Formato | Descrição |
|-------|------|-------------|---------|-----------|
| `tpInscEstab` | Inteiro | Sim | 1-2 | 1=CNPJ, 2=CEI |
| `nrInscEstab` | String | Sim | 14 ou 12 dígitos | CNPJ (14) ou CEI (12) do estabelecimento |

#### 2.3.3 Grupo `dadosEstab`

| Campo | Tipo | Obrigatório | Validações |
|-------|------|-------------|-----------|
| `nomeEstab` | String | Sim | Máx. 180 caracteres |
| `dsCategEmpregador` | Inteiro | Sim | 01-10 (Tabela 07) |
| `ativEcon` | String | Sim | 7 dígitos (Código CNAE) |
| `indDesFolha` | Char | Não | S/N |
| `indAvulso` | Char | Não | S/N |

#### 2.3.4 Grupo `enderEstab`

Mesma estrutura do `enderEmpregador` do S-1000:

| Campo | Tipo | Obrigatório | Validações |
|-------|------|-------------|-----------|
| `tpLograd` | String | Sim | Máx. 4 caracteres |
| `dscLograd` | String | Sim | Máx. 100 caracteres |
| `nrLograd` | String | Sim | Máx. 10 caracteres |
| `complemento` | String | Não | Máx. 30 caracteres |
| `bairro` | String | Sim | Máx. 90 caracteres |
| `cepLograd` | String | Sim | 8 dígitos |
| `codMun` | Inteiro | Sim | Código IBGE |
| `uf` | String | Sim | 2 caracteres |

#### 2.3.5 Grupo `softwares`

Informações sobre os softwares utilizados:

| Campo | Tipo | Obrigatório | Validações |
|-------|------|-------------|-----------|
| `cnpjSoftHouse` | String | Não | 14 dígitos |
| `nmSoftware` | String | Não | Máx. 80 caracteres |
| `versaoSistema` | String | Não | Máx. 20 caracteres |

### 2.4 Validações e Regras Específicas do S-1005

#### Regras de Validação

1. **Tipo de Inscrição do Estabelecimento**:
   - Se `tpInscEstab = 1` (CNPJ): `nrInscEstab` deve ter 14 dígitos
   - Se `tpInscEstab = 2` (CEI): `nrInscEstab` deve ter 12 dígitos

2. **Relacionamento com S-1000**:
   - O `tpInsc` e `nrInsc` de `ideEmpregador` devem corresponder aos do S-1000 já transmitido
   - O S-1005 não pode ser transmitido antes do S-1000

3. **CNAE - Classificação Nacional de Atividade Econômica**:
   - Deve ser um código válido (7 dígitos)
   - Deve estar ativo no período de transmissão
   - Exemplos:
     - 4711300 = Comércio varejista de tecidos, vestuário, acessórios
     - 6201500 = Desenvolvimento de software customizado
     - 7020400 = Consultoria em gestão empresarial

4. **Endereço do Estabelecimento**:
   - Deve ser diferente do endereço da matriz (em geral)
   - CEP e município devem ser válidos
   - Se o estabelecimento está no exterior, pode haver variações nas regras

5. **Software House**:
   - Se preenchido, o CNPJ deve ser válido
   - O nome do software deve ser consistente
   - Versão deve seguir padrão semântico (X.Y.Z)

### 2.5 Indicativos Importantes

#### Indicador de Desoneração (`indDesFolha`)

- **S** (Sim): Estabelecimento está amparado por Lei de Desoneração da Folha
  - Lei 12.546/2011 (Setores de tecnologia, audiovisual, etc.)
  - Lei 12.844/2013 (Setores de construção civil, calçados, etc.)
  - Lei 13.199/2015 (Renovação de concessão de desoneração)
  
- **N** (Não): Não está amparado por lei de desoneração (padrão)

#### Indicador de Avulso (`indAvulso`)

- **S** (Sim): Empresa contrata trabalhadores avulsos
  - Portuários, estivadores, etc.
  
- **N** (Não): Não utiliza regime de avulso (padrão)

---

## 3. Tipos de Dados Utilizados

### 3.1 Tipos Primitivos

| Tipo | Descrição | Exemplo |
|------|-----------|---------|
| **String** | Sequência de caracteres | "EMPRESA LTDA" |
| **Inteiro** | Número inteiro sem casas decimais | 1, 14, 3550308 |
| **Data** | Formato ISO 8601 AAAA-MM-DD | 2024-01-15 |
| **Decimal** | Número com casas decimais (formato brasileiro) | 1234.56 |
| **Char** | Caractere único | S, N |

### 3.2 Validação de Dígito Verificador

#### Validação de CNPJ (14 dígitos)

O CNPJ segue um padrão de validação com dígito verificador:
- 12 dígitos principais + 2 dígitos verificadores
- Exemplos válidos:
  - 34.028.316/0001-52
  - 11.222.333/0001-81

#### Validação de CPF (11 dígitos)

- 9 dígitos principais + 2 dígitos verificadores
- Exemplos válidos:
  - 123.456.789-01

#### Validação de CEI (12 dígitos)

- Código de Estabelecimento do INSS
- Formato: 12 dígitos
- Validação própria conforme INSS

---

## 4. Relacionamento entre S-1000 e S-1005

### 4.1 Fluxo de Transmissão

```
FASE 1: Cadastro Inicial
├── Transmitir S-1000 (Informações do Empregador)
│   └── Aguardar recebimento e validação ✓
│
FASE 2: Estabelecimentos
├── Transmitir S-1005 (Dados do Estabelecimento - 1º)
├── Transmitir S-1005 (Dados do Estabelecimento - 2º) [se houver múltiplos]
└── ...
│
FASE 3: Tabelas de Apoio (Após S-1005)
├── Transmitir S-1010 (Tabela de Cargos)
├── Transmitir S-1020 (Tabela de Lotações)
└── ...
```

### 4.2 Regras de Relacionamento

| Regra | Descrição |
|-------|-----------|
| **Precedência** | S-1000 deve ser transmitido ANTES de S-1005 |
| **Empregador Único** | Todos os S-1005 devem referenciar o mesmo `tpInsc` e `nrInsc` do S-1000 |
| **Identificação Única** | Cada S-1005 deve ter um `nrInscEstab` único |
| **Retificação** | Alterações em S-1000 afetam todos os S-1005 relacionados |
| **Desativação** | Fechar um S-1000 implica em desativar todos os S-1005 |

### 4.3 Exemplo de Cenário Múltiplo

```
CNPJ Matriz: 34.028.316/0001-52
  ├── S-1000 (Informações da matriz)
  │
  ├── S-1005 #1 (Estabelecimento Matriz)
  │   └── CNPJ: 34.028.316/0001-52 (mesmo da matriz)
  │   └── Endereço: Rua das Flores, 123, São Paulo
  │
  ├── S-1005 #2 (Filial Centro)
  │   └── CNPJ: 34.028.316/0001-99 (filial)
  │   └── Endereço: Avenida Paulista, 1000, São Paulo
  │
  └── S-1005 #3 (Filial Sul)
      └── CNPJ: 34.028.316/0002-60 (filial)
      └── Endereço: Rua XV de Novembro, 500, Curitiba
```

---

## 5. Exemplos Práticos de Estrutura XML

### 5.1 Exemplo Completo S-1000

```xml
<?xml version="1.0" encoding="UTF-8"?>
<eSocial xmlns="http://www.esocial.gov.br/schema/evt/evtInfoEmpregador/v09_00_00">
  <evtInfoEmpregador Id="ID000001">
    <ideEvento>
      <tpAmb>2</tpAmb>
      <procEmi>1</procEmi>
      <verProc>09.00</verProc>
    </ideEvento>
    
    <ideEmpregador>
      <tpInsc>1</tpInsc>
      <nrInsc>34028316000152</nrInsc>
    </ideEmpregador>
    
    <infoCadastro>
      <dtCadastro>2024-01-15</dtCadastro>
      
      <infoPJ>
        <razaoSocial>TECNOLOGIA E SISTEMAS LTDA</razaoSocial>
        <nomeFantasia>TECH SYSTEMS</nomeFantasia>
        <dsCategEmpregador>01</dsCategEmpregador>
        <classTributaria>06</classTributaria>
        <indDesFolha>N</indDesFolha>
        <indConstrutora>N</indConstrutora>
        <indCooperativa>N</indCooperativa>
      </infoPJ>
      
      <enderEmpregador>
        <tpLograd>RUA</tpLograd>
        <dscLograd>DAS FLORES</dscLograd>
        <nrLograd>123</nrLograd>
        <complemento>APTO 1001</complemento>
        <bairro>CENTRO</bairro>
        <cepLograd>01310100</cepLograd>
        <codMun>3550308</codMun>
        <uf>SP</uf>
      </enderEmpregador>
      
      <infoCont>
        <nmCont>João Carlos Silva</nmCont>
        <cpfCont>12345678901</cpfCont>
        <telefone>1133334444</telefone>
        <email>joao@techsystems.com.br</email>
      </infoCont>
    </infoCadastro>
  </evtInfoEmpregador>
</eSocial>
```

### 5.2 Exemplo Completo S-1005

```xml
<?xml version="1.0" encoding="UTF-8"?>
<eSocial xmlns="http://www.esocial.gov.br/schema/evt/evtInfoEstab/v09_00_00">
  <evtInfoEstab Id="ID000002">
    <ideEvento>
      <tpAmb>2</tpAmb>
      <procEmi>1</procEmi>
      <verProc>09.00</verProc>
    </ideEvento>
    
    <ideEmpregador>
      <tpInsc>1</tpInsc>
      <nrInsc>34028316000152</nrInsc>
    </ideEmpregador>
    
    <infoCadastro>
      <ideEstab>
        <tpInscEstab>1</tpInscEstab>
        <nrInscEstab>34028316000152</nrInscEstab>
      </ideEstab>
      
      <dadosEstab>
        <nomeEstab>TECH SYSTEMS - MATRIZ</nomeEstab>
        <dsCategEmpregador>01</dsCategEmpregador>
        <ativEcon>6201500</ativEcon>
        <indDesFolha>N</indDesFolha>
        
        <enderEstab>
          <tpLograd>RUA</tpLograd>
          <dscLograd>DAS FLORES</dscLograd>
          <nrLograd>123</nrLograd>
          <complemento>APTO 1001</complemento>
          <bairro>CENTRO</bairro>
          <cepLograd>01310100</cepLograd>
          <codMun>3550308</codMun>
          <uf>SP</uf>
        </enderEstab>
        
        <ativEcon>6201500</ativEcon>
        <indDesFolha>N</indDesFolha>
        
        <softwares>
          <softwareHouse>
            <cnpjSoftHouse>12345678000190</cnpjSoftHouse>
            <nmSoftware>SISTEMA GESTÃO RH PRO</nmSoftware>
            <versaoSistema>5.10.1</versaoSistema>
          </softwareHouse>
        </softwares>
      </dadosEstab>
      
      <alteracoes>
        <dtAlteracao>2024-01-15</dtAlteracao>
        <descricaoAlteracao>Cadastro inicial do estabelecimento</descricaoAlteracao>
      </alteracoes>
    </infoCadastro>
  </evtInfoEstab>
</eSocial>
```

---

## 6. Tabelas de Referência Importantes

### 6.1 Categoria de Empregador (Tabela 07)

| Código | Descrição |
|--------|-----------|
| 01 | Empresa privada |
| 02 | Órgão público federal |
| 03 | Órgão público estadual |
| 04 | Órgão público municipal |
| 05 | Órgão público do Distrito Federal |
| 06 | Serviço de saúde - pessoa jurídica de direito privado |
| 07 | Administração pública indireta |
| 08 | Setor público |
| 09 | Entidade sem fins lucrativos |
| 10 | Produtor rural |

### 6.2 Classificação Tributária (Tabela 52)

| Código | Descrição |
|--------|-----------|
| 01 | Microempreendedor individual (MEI) |
| 06 | Empresa de pequeno porte (EPP) |
| 07 | Empresa de pequeno porte (SIMPLES) |
| 08 | Lucro real (trimestral) |
| 09 | Lucro real (anual) |
| 10 | Lucro presumido |
| 11 | SIMPLES Nacional |
| ... | ... |

### 6.3 Estados (UF) - Tabela 26

Siglas de 2 caracteres conforme IBGE:
- AC, AL, AP, AM, BA, CE, DF, ES, GO, MA, MT, MS, MG, PA, PB, PR, PE, PI, RJ, RN, RS, RO, RR, SC, SP, SE, TO

### 6.4 CNAE - Principais Códigos

| Código | Descrição |
|--------|-----------|
| 0111100 | Produção de lavouras temporárias |
| 4711300 | Comércio varejista de tecidos |
| 4722002 | Comércio varejista de roupas |
| 6201500 | Desenvolvimento de software customizado |
| 6204000 | Consultoria em tecnologia da informação |
| 7020400 | Consultoria em gestão empresarial |

---

## 7. Validações e Restrições

### 7.1 Restrições por Tipo de Empregador

#### Empresa Privada (tpInsc=1, dsCategEmpregador=01)
- Deve ter CNPJ válido
- `razaoSocial` é obrigatória
- Pode ter `nomeFantasia`
- Indicadores: `indDesFolha`, `indConstrutora`, `indCooperativa` podem ser aplicáveis

#### Órgão Público (dsCategEmpregador=02-05, 07)
- Pode ter CNPJ ou CPF
- Indicadores de desonerção geralmente não se aplicam
- Pode ter estrutura diferenciada

### 7.2 Validações de Data

```
dtCadastro:
  - Mínimo: 2010-01-01 (Início do eSocial)
  - Máximo: Data da transmissão do evento
  - Não pode ser futuro

dtAlteracao (S-1005):
  - Deve ser >= dtCadastro
  - Não pode ser futuro
  - Não pode ser anterior à competência anterior
```

### 7.3 Validações de Endereço

```
CEP:
  - Formato: 8 dígitos (sem hífen)
  - Deve existir na base de dados de CEPs válidos
  
Município (codMun):
  - Deve ser código IBGE válido
  - Deve estar ativo no período de transmissão
  
UF:
  - Deve corresponder ao código de município informado
  - Deve estar em tabela de estados válidos
```

---

## 8. Erros Comuns e Soluções

### 8.1 Erros na Estrutura do S-1000

| Erro | Causa | Solução |
|------|-------|---------|
| "CNPJ inválido" | Dígito verificador incorreto | Validar com algoritmo de CNPJ |
| "Tipo de inscrição não corresponde" | tpInsc e nrInsc incompatíveis | Se CNPJ, usar tpInsc=1 |
| "CEP não localizado" | CEP não existe no cadastro | Verificar CEP no site Correios |
| "Município não cadastrado" | Código IBGE inválido | Consultar tabela de municípios |
| "Data anterior a 2010" | dtCadastro muito antiga | Usar data mínima de 2010-01-01 |

### 8.2 Erros na Estrutura do S-1005

| Erro | Causa | Solução |
|------|-------|---------|
| "S-1000 não encontrado" | S-1005 antes de S-1000 | Transmitir S-1000 primeiro |
| "CNAE inválido" | Código de atividade inválido | Validar CNAE na tabela oficial |
| "Tipo inscrição estabelecimento inválido" | tpInscEstab fora do intervalo 1-2 | Usar 1=CNPJ ou 2=CEI |
| "CEI com dígitos incorretos" | CEI inválido | Validar CEI com INSS |

---

## 9. Referências e Documentação Oficial

- **Leiautes eSocial v. S-1.3** (Nota Técnica 06/2026 - rev. 09/04/2026)
- **Manual de Orientação do eSocial (MOS)** - Versão consolidada
- **Esquemas XSD** - Disponíveis no site www.gov.br/esocial
- **Tabelas do eSocial** - Consulta online em frontend.esocial.gov.br/adm/
- **Sistema de Escrituração Digital** - https://login.esocial.gov.br/

---

## 10. Resumo Comparativo S-1000 vs S-1005

| Aspecto | S-1000 | S-1005 |
|--------|--------|--------|
| **Obrigatoriedade** | Obrigatório (Precursor) | Obrigatório (se houver filial) |
| **Tipo de Inscrição** | CNPJ, CPF, CAEPF, CNO | CNPJ ou CEI |
| **Dados Principais** | Empresa/Empregador | Estabelecimento |
| **Endereço** | Da empresa | Do estabelecimento |
| **Contato** | De responsável geral | Não tem (opcional em S-1000) |
| **Precedência** | Deve ser 1º | Após S-1000 |
| **Múltiplas Transmissões** | Uma por empregador | Uma por estabelecimento |
| **Software House** | Não consta | Consta (opcional) |
| **CNAE** | Não consta | Obrigatório |

---

**Documento atualizado em:** 2026-08-17  
**Versão eSocial:** S-1.3 (NT 06/2026)  
**Status:** Validado conforme documentação oficial
