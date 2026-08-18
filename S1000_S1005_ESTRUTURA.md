# Estrutura dos Eventos S-1000 e S-1005 (eSocial v. S-1.3)

## Visão Geral

Os eventos **S-1000** (Informações do Empregador) e **S-1005** (Dados do Estabelecimento) formam a base inicial de qualquer transmissão no eSocial. São eventos precursores que devem ser transmitidos antes de qualquer outro evento.

| Propriedade | S-1000 | S-1005 |
|-------------|--------|--------|
| **Descrição** | Informações do Empregador | Dados do Estabelecimento |
| **Obrigatoriedade** | ✅ Obrigatório (1º) | ✅ Obrigatório (2º) |
| **Ordem de Transmissão** | Primeiro | Depois de S-1000 |
| **Frequência** | Esporádico (alterações) | Esporádico (alterações) |
| **Evento Precursor** | Sim | Depende de S-1000 |

---

## 1. S-1000: Informações do Empregador

### 1.1 Propósito
Registrar dados cadastrais da empresa/empregador (matriz). É o evento inicial que deve ser transmitido para inicializar o ambiente eSocial.

### 1.2 Estrutura Hierárquica

```
evtInfoEmpregador
├── ideEmpregador
│   ├── tpInsc (Tipo de Inscrição)
│   └── nrInsc (Número de Inscrição)
├── infoCadastro
│   ├── dtCadastro (Data do Cadastro)
│   ├── infoPJ (Dados PJ) ou infoPF (Dados PF)
│   ├── enderEmpregador (Endereço)
│   ├── infoContato (Contato) [Opcional]
│   └── infoIndicadores (Indicadores Especiais) [Opcional]
└── infoComplementar [Opcional]
    ├── indAprendiz (Indicador Aprendiz)
    ├── nrRegPJ (Registro no CRJ)
    └── outras informações complementares
```

### 1.3 Campos Obrigatórios

#### Identificação do Empregador
| Campo | Tipo | Tamanho | Descrição |
|-------|------|--------|-----------|
| `tpInsc` | Numérico | 1 | 1=CNPJ, 2=CPF, 3=CAEPF, 4=CNO |
| `nrInsc` | Numérico | 11-14 | Número de inscrição sem formatação |

#### Dados da Empresa (infoPJ - para PJ)
| Campo | Tipo | Tamanho | Descrição | Obrigatório |
|-------|------|--------|-----------|------------|
| `razaoSocial` | String | 4-180 | Razão social da empresa | ✅ |
| `nomeFantasia` | String | 0-180 | Nome fantasia (se houver) | ❌ |
| `dsCategEmpregador` | Char | 2 | Categoria do empregador (01-10) | ✅ |
| `classTributaria` | Char | 2 | Classificação tributária (01-11) | ✅ |
| `indConstrutora` | Char | 1 | Construtora de imóveis (S/N) | ❌ |
| `indDesFolha` | Char | 1 | Desonerado (S/N) | ❌ |
| `indCooperativa` | Char | 1 | Cooperativa (S/N) | ❌ |

#### Endereço (enderEmpregador)
| Campo | Tipo | Tamanho | Descrição | Obrigatório |
|-------|------|--------|-----------|------------|
| `tpLograd` | String | - | Tipo logradouro (RUA, AV, PÇ, etc.) | ✅ |
| `dscLograd` | String | 1-100 | Nome do logradouro | ✅ |
| `nrLograd` | String | 1-10 | Número | ✅ |
| `cmplLograd` | String | 0-30 | Complemento | ❌ |
| `bairro` | String | 1-90 | Nome do bairro | ✅ |
| `cepLograd` | Numérico | 8 | CEP sem formatação | ✅ |
| `codMun` | Numérico | 7 | Código IBGE do município | ✅ |
| `uf` | Char | 2 | Sigla UF (SP, RJ, etc.) | ✅ |

#### Contato (infoContato) - Opcional
| Campo | Tipo | Tamanho | Descrição |
|-------|------|--------|-----------|
| `nmContato` | String | 1-70 | Nome do contato |
| `cpfContato` | Numérico | 11 | CPF do contato (sem formatação) |
| `telefone` | String | 6-15 | Telefone |
| `email` | String | 0-120 | E-mail |

### 1.4 Validações Principais

| Campo | Regra de Validação |
|-------|-------------------|
| **CNPJ** | 14 dígitos com dígito verificador válido |
| **CPF** | 11 dígitos com dígito verificador válido |
| **CEP** | 8 dígitos, deve existir na base de endereços |
| **Razão Social** | Não vazia, caracteres alfanuméricos |
| **Categoria Empregador** | Valor entre 01 e 10 (conforme tabela eSocial) |
| **Classificação Tributária** | Valor entre 01 e 11 (conforme tabela eSocial) |
| **Município** | Código IBGE válido e ativo |
| **Data do Cadastro** | >= 2010-01-01, não pode ser data futura |

### 1.5 Exemplo XML - S-1000

```xml
<?xml version="1.0" encoding="UTF-8"?>
<eSocial>
  <evtInfoEmpregador id="ID001">
    <ideEmpregador>
      <tpInsc>1</tpInsc>
      <nrInsc>34028316000152</nrInsc>
    </ideEmpregador>
    <infoCadastro>
      <dtCadastro>2024-01-15</dtCadastro>
      <infoPJ>
        <razaoSocial>TECH SYSTEMS LTDA</razaoSocial>
        <nomeFantasia>TECH SOLUTIONS</nomeFantasia>
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
        <cmplLograd>APTO 45</cmplLograd>
        <bairro>CENTRO</bairro>
        <cepLograd>01310100</cepLograd>
        <codMun>3550308</codMun>
        <uf>SP</uf>
      </enderEmpregador>
      <infoContato>
        <nmContato>JOÃO DA SILVA</nmContato>
        <cpfContato>12345678901</cpfContato>
        <telefone>1133334444</telefone>
        <email>contato@techsystems.com.br</email>
      </infoContato>
    </infoCadastro>
  </evtInfoEmpregador>
</eSocial>
```

---

## 2. S-1005: Dados do Estabelecimento

### 2.1 Propósito
Registrar informações específicas de cada filial/estabelecimento da empresa. Para cada estabelecimento (incluindo matriz), é necessário transmitir um evento S-1005.

### 2.2 Estrutura Hierárquica

```
evtInfoEstab
├── ideEmpregador
│   ├── tpInsc (Tipo de Inscrição)
│   └── nrInsc (Número de Inscrição)
├── infoCadastro
│   ├── ideEstab (Identificação do Estabelecimento)
│   │   ├── tpInscEstab (Tipo de Inscrição)
│   │   └── nrInscEstab (Número de Inscrição)
│   ├── dadosEstab (Dados do Estabelecimento)
│   │   ├── nomeEstab (Nome)
│   │   ├── dsCategEmpregador (Categoria)
│   │   ├── ativEcon (CNAE - Atividade Econômica)
│   │   ├── enderEstab (Endereço)
│   │   ├── infoSoftHouse (Info Software House) [Opcional]
│   │   ├── infoIntermediaria (Info Intermediária) [Opcional]
│   │   └── infoComplementar (Info Complementar) [Opcional]
│   └── alteracao (Alteração) [Opcional]
└── contatoPJ (Contato PJ) [Opcional]
```

### 2.3 Campos Obrigatórios

#### Identificação do Empregador
| Campo | Tipo | Tamanho | Descrição |
|-------|------|--------|-----------|
| `tpInsc` | Numérico | 1 | Deve corresponder ao S-1000 |
| `nrInsc` | Numérico | 11-14 | Deve corresponder ao S-1000 |

#### Identificação do Estabelecimento
| Campo | Tipo | Tamanho | Descrição | Obrigatório |
|-------|------|--------|-----------|------------|
| `tpInscEstab` | Numérico | 1 | 1=CNPJ, 2=CEI | ✅ |
| `nrInscEstab` | Numérico | 12-14 | CNPJ ou CEI sem formatação | ✅ |

#### Dados do Estabelecimento (dadosEstab)
| Campo | Tipo | Tamanho | Descrição | Obrigatório |
|-------|------|--------|-----------|------------|
| `nomeEstab` | String | 1-180 | Nome do estabelecimento | ✅ |
| `dsCategEmpregador` | Char | 2 | Categoria (01-10) | ✅ |
| `ativEcon` | Numérico | 7 | CNAE (ex: 6201500) | ✅ |
| `indDesFolha` | Char | 1 | Desonerado (S/N) | ❌ |
| `indConstrutora` | Char | 1 | Construtora (S/N) | ❌ |
| `indIntermediaria` | Char | 1 | Intermediária (S/N) | ❌ |

#### Endereço (enderEstab)
Mesma estrutura do S-1000 (ver seção 1.3)

#### Contato (contatoPJ) - Opcional
| Campo | Tipo | Tamanho | Descrição |
|-------|------|--------|-----------|
| `nmContato` | String | 1-70 | Nome do contato |
| `cpfContato` | Numérico | 11 | CPF do contato |
| `telefone` | String | 6-15 | Telefone |
| `email` | String | 0-120 | E-mail |

#### Software House (infoSoftHouse) - Opcional
| Campo | Tipo | Tamanho | Descrição |
|-------|------|--------|-----------|
| `cnpjSoftHouse` | Numérico | 14 | CNPJ da software house |
| `nmSoftHouse` | String | 1-80 | Nome da software house |
| `versaoSoftHouse` | String | 1-20 | Versão do software |

### 2.4 Validações Principais

| Campo | Regra de Validação |
|-------|-------------------|
| **tpInsc/nrInsc** | Deve corresponder exatamente ao S-1000 (precedência) |
| **CNPJ Estab. (tpInscEstab=1)** | 14 dígitos com dígito verificador válido |
| **CEI (tpInscEstab=2)** | 12 dígitos, formato específico |
| **CNAE** | 7 dígitos válido (deve existir na tabela eSocial) |
| **Nome Estabelecimento** | Não vazio, diferente da razão social geralmente |
| **Categoria** | 01-10 (deve ser compatível com S-1000) |
| **Endereço** | Mesmas regras do S-1000 |
| **Precedência** | S-1000 deve ser transmitido antes |

### 2.5 Exemplo XML - S-1005 (Matriz)

```xml
<?xml version="1.0" encoding="UTF-8"?>
<eSocial>
  <evtInfoEstab id="ID002">
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
        <indConstrutora>N</indConstrutora>
      </dadosEstab>
      <enderEstab>
        <tpLograd>RUA</tpLograd>
        <dscLograd>DAS FLORES</dscLograd>
        <nrLograd>123</nrLograd>
        <cmplLograd>APTO 45</cmplLograd>
        <bairro>CENTRO</bairro>
        <cepLograd>01310100</cepLograd>
        <codMun>3550308</codMun>
        <uf>SP</uf>
      </enderEstab>
      <infoSoftHouse>
        <cnpjSoftHouse>12345678000199</cnpjSoftHouse>
        <nmSoftHouse>SOFTWARETECH SOLUTIONS</nmSoftHouse>
        <versaoSoftHouse>2.1.0</versaoSoftHouse>
      </infoSoftHouse>
    </infoCadastro>
    <contatoPJ>
      <nmContato>MARIA SANTOS</nmContato>
      <cpfContato>12345678902</cpfContato>
      <telefone>1133334444</telefone>
      <email>maria@techsystems.com.br</email>
    </contatoPJ>
  </evtInfoEstab>
</eSocial>
```

### 2.6 Exemplo XML - S-1005 (Filial)

```xml
<?xml version="1.0" encoding="UTF-8"?>
<eSocial>
  <evtInfoEstab id="ID003">
    <ideEmpregador>
      <tpInsc>1</tpInsc>
      <nrInsc>34028316000152</nrInsc>
    </ideEmpregador>
    <infoCadastro>
      <ideEstab>
        <tpInscEstab>1</tpInscEstab>
        <nrInscEstab>34028316000160</nrInscEstab>
      </ideEstab>
      <dadosEstab>
        <nomeEstab>TECH SYSTEMS - FILIAL RIO</nomeEstab>
        <dsCategEmpregador>01</dsCategEmpregador>
        <ativEcon>6201500</ativEcon>
        <indDesFolha>N</indDesFolha>
      </dadosEstab>
      <enderEstab>
        <tpLograd>AVENIDA</tpLograd>
        <dscLograd>PAULISTA</dscLograd>
        <nrLograd>1000</nrLograd>
        <bairro>BELA VISTA</bairro>
        <cepLograd>01311100</cepLograd>
        <codMun>3300100</codMun>
        <uf>RJ</uf>
      </enderEstab>
    </infoCadastro>
  </evtInfoEstab>
</eSocial>
```

---

## 3. Tabelas de Referência

### 3.1 Categoria de Empregador (Tabela 07)

| Código | Descrição |
|--------|-----------|
| 01 | Empresa privada |
| 02 | Órgão público da administração federal |
| 03 | Órgão público da administração estadual ou DF |
| 04 | Órgão público da administração municipal |
| 05 | Órgão do poder judiciário |
| 06 | Órgão do poder legislativo |
| 07 | Sindicato de trabalhadores |
| 08 | Sindicato de empregadores |
| 09 | Entidade sem fins lucrativos |
| 10 | Produtor rural |

### 3.2 Classificação Tributária (Tabela 52)

| Código | Descrição |
|--------|-----------|
| 01 | MEI (Microempreendedor Individual) |
| 02 | Empresa de economia mista |
| 03 | Sociedade em nome coletivo |
| 04 | Sociedade em comandita simples |
| 05 | Sociedade limitada |
| 06 | EPP (Empresa de Pequeno Porte) |
| 07 | Pessoa física |
| 08 | Lucro real |
| 09 | Lucro presumido |
| 10 | Lucro arbitrado |
| 11 | SIMPLES Nacional |

### 3.3 CNAE - Atividades Principais

| Código | Descrição |
|--------|-----------|
| 0111100 | Cultivo de cereais |
| 4711300 | Comércio varejista de tecidos |
| 5611201 | Restaurante com serviço à la carte |
| 6201500 | Desenvolvimento de software customizado |
| 6204000 | Consultoria em tecnologia da informação |
| 6309100 | Tratamento de dados, provedoria de serviços |
| 7020400 | Consultoria em gestão empresarial |
| 8011100 | Atividades de administração pública |
| 8532100 | Educação superior |
| 8711300 | Atividades de vigilância e segurança |

### 3.4 Tipo de Logradouro

| Abreviatura | Descrição |
|-------------|-----------|
| RUA | Rua |
| AV | Avenida |
| TRAV | Travessa |
| PÇ | Praça |
| PRAD | Parada |
| ESTR | Estrada |
| ALC | Alcançara |
| ROD | Rodovia |
| LARGO | Largo |
| BECO | Beco |

---

## 4. Regras de Negócio

### 4.1 Ordem de Transmissão

```
1. S-1000 (Empregador) - OBRIGATÓRIO
   ↓
2. S-1005 (Matriz) - OBRIGATÓRIO
   ↓
3. S-1005 (Filial 1) - Se houver
   ↓
4. S-1005 (Filial 2) - Se houver
   ↓
5. S-1005+ (N filiais) - Quantas forem necessárias
   ↓
6. S-1010+ (Tabelas e demais eventos)
```

### 4.2 Relacionamento entre S-1000 e S-1005

- **S-1005 depende de S-1000**: Não é possível transmitir S-1005 sem que S-1000 tenha sido transmitido antes
- **Mesmo empregador**: Os campos `tpInsc` e `nrInsc` de S-1005 devem ser idênticos ao S-1000
- **CNPJ/CEI diferentes**: O `nrInscEstab` pode ser diferente (CNPJ da filial, por exemplo)

### 4.3 Alterações

#### S-1000 - Alteração de Dados
Se dados da empresa forem alterados (razão social, endereço, etc.), deve-se:
1. Respeitar o campo `dtCadastro` (data não muda)
2. Atualizar os campos modificados
3. Re-transmitir o evento com os dados atualizados

#### S-1005 - Alteração de Dados
Se dados de estabelecimento forem alterados, deve-se:
1. Incluir informação na seção `alteracao` (se aplicável)
2. Atualizar os campos modificados
3. Re-transmitir o evento com os dados atualizados

### 4.4 Indicadores Especiais

| Indicador | Significado | Valores |
|-----------|-----------|---------|
| `indDesFolha` | Empresa desonerada (Lei nº 12.546/2011) | S/N |
| `indConstrutora` | Construtora de imóveis | S/N |
| `indCooperativa` | Cooperativa de trabalho | S/N |
| `indIntermediaria` | Intermediária de mão de obra | S/N |

---

## 5. Erros Comuns e Soluções

| Erro | Causa Provável | Solução |
|------|---|---|
| **CNPJ inválido** | Dígito verificador incorreto | Validar o cálculo do dígito verificador |
| **CEP não localizado** | CEP inexistente ou incorreto | Verificar em www.correios.com.br |
| **S-1000 não encontrado** | Tentar transmitir S-1005 antes de S-1000 | Transmitir S-1000 primeiro |
| **Campos de empregador não correspondem** | tpInsc/nrInsc de S-1005 ≠ S-1000 | Copiar exatamente os mesmos valores |
| **CNAE inválido** | Código de atividade não existe | Consultar tabela de CNAE do eSocial |
| **Município não cadastrado** | Código IBGE incorreto ou inativo | Verificar código IBGE válido |
| **Data futura no cadastro** | dtCadastro com data maior que hoje | Usar data atual ou passada |
| **Categoria não compatível** | dsCategEmpregador incompatível com tributação | Verificar compatibilidade na tabela 07 |

---

## 6. Tipos de Dados

| Tipo | Descrição | Tamanho | Formato | Exemplo |
|------|-----------|---------|---------|---------|
| **String** | Sequência de caracteres | Variável | Alfanumérico | "EMPRESA LTDA" |
| **Numérico** | Números inteiros | Variável | Somente dígitos | 34028316000152 |
| **Decimal** | Números com casas decimais | Variável | XX.XX ou XX,XX | 1234.56 |
| **Char** | Um caractere único | 1 | S, N, 0-9 | S |
| **Data** | Data no formato ISO | 10 | YYYY-MM-DD | 2024-01-15 |

---

## 7. Checklist de Implementação

### Antes de Transmitir S-1000
- [ ] CNPJ/CPF validado (dígito verificador correto)
- [ ] Razão social preenchida (4-180 caracteres)
- [ ] Categoria de empregador válida (01-10)
- [ ] Classificação tributária válida (01-11)
- [ ] Endereço completo (logradouro, número, bairro, CEP, município, UF)
- [ ] CEP existe e está correto
- [ ] Código IBGE do município é válido
- [ ] Data do cadastro não é futura
- [ ] Tipo de logradouro preenchido

### Antes de Transmitir S-1005
- [ ] S-1000 foi transmitido com sucesso
- [ ] tpInsc e nrInsc correspondem exatamente ao S-1000
- [ ] CNPJ/CEI do estabelecimento validado
- [ ] Nome do estabelecimento preenchido (1-180 caracteres)
- [ ] Categoria do empregador preenchida (01-10)
- [ ] CNAE válido (7 dígitos)
- [ ] Endereço completo do estabelecimento
- [ ] CEP existe e está correto
- [ ] Código IBGE do município é válido
- [ ] Se software house, CNPJ validado e nome/versão preenchidos

---

## 8. Documentação Relacionada

Para complementar este documento, consulte:
- Leiaute oficial do eSocial v. S-1.3
- Norma Técnica (NT) 06/2026
- Tabelas oficiais de CNAE
- Tabelas de municípios IBGE
- Manual técnico de validações

---

**Última atualização**: 2024-08-17  
**Versão do eSocial**: S-1.3  
**Status**: Documentação Ativa
