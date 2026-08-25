# Plan: Gerar Estruturas XML para Eventos S-1000 e S-1005 (Domínio Empregador)

**TL;DR**: Criar estruturas Go para os eventos S-1000 (Informações do Empregador) e S-1005 (Dados do Estabelecimento) na pasta `empregador/xml/`, reutilizando os tipos validados de `pkg/types` (CNPJ, CPF, DataApenas, SimNao, DecimalValor) e seguindo a arquitetura existente do projeto.

---

## Descobertas

- **Tipos disponíveis**: CNPJ (alfanumérico 14 dígitos com dígito verificador), CPF (11 dígitos), DataApenas (YYYY-MM-DD), SimNao (1/2), DecimalValor (com casas decimais)
- **Estrutura do projeto**: `esocial/internal/model/empregador/xml/` existe mas está vazia
- **Padrão**: Projeto usa Go com `encoding/xml` para serialização
- **Documentação existente**: S1000_S1005_ESTRUTURA.md descreve campos, obrigatoriedade e validações detalhadas

---

## Steps

### 1. Criar entidades Go para S-1000
**Arquivo**: `esocial/internal/model/empregador/xml/s1000_info_empregador.go`

**Estrutura hierárquica**:
```
EvtInfoEmpregador
├── IdeEmpregador
│   ├── TpInsc (int: 1=CNPJ, 2=CPF, 3=CAEPF, 4=CNO)
│   └── NrInsc (CNPJ ou CPF)
└── InfoCadastro
    ├── DtCadastro (DataApenas)
    ├── InfoPJ
    │   ├── RazaoSocial (string, 4-180 chars)
    │   ├── NomeFantasia (string, opcional, 0-180 chars)
    │   ├── DsCategEmpregador (string, 2 chars: 01-10)
    │   ├── ClassTributaria (string, 2 chars: 01-11)
    │   ├── IndConstrutora (SimNao, opcional)
    │   ├── IndDesFolha (SimNao, opcional)
    │   └── IndCooperativa (SimNao, opcional)
    ├── EnderEmpregador (tipo Endereco)
    ├── InfoContato (opcional)
    └── InfoComplementar (opcional)
```

**Obrigatoriedade**:
- ✅ Obrigatório: TpInsc, NrInsc, DtCadastro, RazaoSocial, DsCategEmpregador, ClassTributaria
- ✅ Obrigatório (Endereço): TpLograd, DscLograd, NrLograd, Bairro, CepLograd, CodMun, UF
- ❌ Opcional: NomeFantasia, CmplLograd, InfoContato, indicadores, InfoComplementar

**Tags XML**: camelCase conforme leiaute oficial (tpInsc, nrInsc, razaoSocial, dsCategEmpregador, classTributaria, etc.)

---

### 2. Criar entidades Go para S-1005
**Arquivo**: `esocial/internal/model/empregador/xml/s1005_info_estab.go`

**Estrutura hierárquica**:
```
EvtInfoEstab
├── IdeEmpregador
│   ├── TpInsc (deve corresponder a S-1000)
│   └── NrInsc (deve corresponder a S-1000)
├── InfoCadastro
│   ├── IdeEstab
│   │   ├── TpInscEstab (int: 1=CNPJ, 2=CEI)
│   │   └── NrInscEstab (CNPJ ou CEI)
│   ├── DadosEstab
│   │   ├── NomeEstab (string, 1-180 chars)
│   │   ├── DsCategEmpregador (string, 2 chars: 01-10)
│   │   ├── AtivEcon (string, 7 dígitos - CNAE, ex: 6201500)
│   │   ├── IndDesFolha (SimNao, opcional)
│   │   ├── IndConstrutora (SimNao, opcional)
│   │   └── IndIntermediaria (SimNao, opcional)
│   ├── EnderEstab (tipo Endereco)
│   ├── InfoSoftHouse (opcional)
│   └── Alteracao (opcional)
└── ContatoPJ (opcional)
```

**Obrigatoriedade**:
- ✅ Obrigatório (S-1005): TpInscEstab, NrInscEstab, NomeEstab, DsCategEmpregador, AtivEcon
- ✅ Obrigatório (Referência S-1000): TpInsc, NrInsc (IDÊNTICOS ao S-1000)
- ✅ Obrigatório (Endereço): TpLograd, DscLograd, NrLograd, Bairro, CepLograd, CodMun, UF
- ❌ Opcional: CmplLograd, IndDesFolha, IndConstrutora, IndIntermediaria, InfoSoftHouse, ContatoPJ, Alteracao

**Regra de Precedência**: S-1000 deve ser transmitido ANTES de S-1005. Os campos tpInsc e nrInsc de S-1005 devem corresponder exatamente aos de S-1000.

---

### 3. Criar tipos auxiliares compartilhados
**Arquivo**: `esocial/internal/model/empregador/xml/types.go`

**Tipos a criar**:

#### 3.1 Endereco (tipo comum para S-1000 e S-1005)
```go
type Endereco struct {
    TpLograd   string // RUA, AV, TRAV, PÇ, etc. (obrigatório)
    DscLograd  string // Nome do logradouro (obrigatório, 1-100 chars)
    NrLograd   string // Número (obrigatório, 1-10 chars)
    CmplLograd string `xml:",omitempty"` // Complemento (opcional, 0-30 chars)
    Bairro     string // Bairro (obrigatório, 1-90 chars)
    CepLograd  string // CEP (obrigatório, 8 dígitos sem formatação)
    CodMun     string // Código IBGE do município (obrigatório, 7 dígitos)
    UF         string // Sigla UF (obrigatório, 2 chars: SP, RJ, MG, etc.)
}
```

#### 3.2 Contato (tipo reutilizável para S-1000 e S-1005)
```go
type Contato struct {
    NmContato   string `xml:",omitempty"` // Nome (opcional, 1-70 chars)
    CpfContato  CPF    `xml:",omitempty"` // CPF (opcional, 11 dígitos)
    Telefone    string `xml:",omitempty"` // Telefone (opcional, 6-15 chars)
    Email       string `xml:",omitempty"` // E-mail (opcional, 0-120 chars)
}
```

#### 3.3 InfoSoftHouse (para S-1005, opcional)
```go
type InfoSoftHouse struct {
    CnpjSoftHouse     CNPJ   // CNPJ da software house (14 dígitos)
    NmSoftHouse       string // Nome (1-80 chars)
    VersaoSoftHouse   string // Versão (1-20 chars)
}
```

#### 3.4 Constantes
```go
const (
    // TpInsc - Tipo de Inscrição
    TpInscCNPJ   = 1
    TpInscCPF    = 2
    TpInscCAEPF  = 3
    TpInscCNO    = 4
    
    // TpInscEstab - Tipo de Inscrição do Estabelecimento
    TpInscEstabCNPJ = 1
    TpInscEstabCEI  = 2
    
    // DsCategEmpregador - Categorias (Tabela 07)
    CategEmpPrivada            = "01"
    CategAdmFederal            = "02"
    CategAdmEstadual           = "03"
    CategAdmMunicipal          = "04"
    CategOrgaoPodJudiciario    = "05"
    CategOrgaoPodLegislativo   = "06"
    CategSindicatTrabalhadores = "07"
    CategSindicatEmpregadores  = "08"
    CategEntidadeSemFinsLucro  = "09"
    CategProdutorRural         = "10"
    
    // ClassTributaria - Classificação Tributária (Tabela 52)
    ClassTribMEI              = "01"
    ClassTribEconomyMista     = "02"
    ClassTribSocieNomeColetivo = "03"
    ClassTribSocieComandita    = "04"
    ClassTribSocieAnonima      = "05"
    ClassTribEPP               = "06"
    ClassTribPessoaFisica      = "07"
    ClassTribLucroReal         = "08"
    ClassTribLucroPresumido    = "09"
    ClassTribLucroArbitrado    = "10"
    ClassTribSimplesNacional   = "11"
)
```

---

### 4. Criar exemplos e funções helper
**Arquivo**: `esocial/internal/model/empregador/xml/examples.go`

**Funções de exemplo**:

#### 4.1 ExemploS1000Empresa()
Retorna uma instância de `EvtInfoEmpregador` com dados preenchidos:
- Empresa: "TECH SYSTEMS LTDA"
- CNPJ: 34028316000152
- Categoria: "01" (Empresa privada)
- Classificação: "06" (EPP)
- Endereço: Rua das Flores, 123, Centro, São Paulo, SP (CEP 01310100, município IBGE 3550308)
- Contato: João da Silva, CPF, telefone, e-mail

#### 4.2 ExemploS1005Matriz()
Retorna uma instância de `EvtInfoEstab` para a matriz (mesmo CNPJ da empresa):
- CNPJ Matriz: 34028316000152 (mesmo de S-1000)
- Nome: "TECH SYSTEMS - MATRIZ"
- CNAE: "6201500" (Desenvolvimento de software customizado)
- Endereço: Rua das Flores, 123, Centro, São Paulo, SP (CEP 01310100)
- Software House: SOFTWARETECH SOLUTIONS, versão 2.1.0

#### 4.3 ExemploS1005Filial()
Retorna uma instância de `EvtInfoEstab` para uma filial:
- Empregador: CNPJ 34028316000152 (referência a S-1000)
- Estabelecimento: CNPJ 34028316000160 (CNPJ da filial)
- Nome: "TECH SYSTEMS - FILIAL RIO"
- CNAE: "6201500"
- Endereço: Avenida Paulista, 1000, Bela Vista, Rio de Janeiro, RJ (CEP 01311100, município IBGE 3300100)

---

### 5. Criar testes unitários
**Arquivos**: 
- `esocial/internal/model/empregador/xml/s1000_info_empregador_test.go`
- `esocial/internal/model/empregador/xml/s1005_info_estab_test.go`

**Testes para S-1000**:
1. `TestS1000MarshalXML_campos_obrigatorios` — Validar que marshal gera XML com campos obrigatórios
2. `TestS1000MarshalXML_campos_opcionais_omitempty` — Validar que campos vazios têm tag omitempty
3. `TestS1000MarshalXML_tipos_especiais` — Validar serialização de CNPJ, DataApenas, SimNao
4. `TestS1000MarshalXML_ordem_elementos` — Validar ordem dos elementos XML (ideEmpregador antes de infoCadastro)
5. `TestS1000ExemploValido` — Marshal do exemplo S-1000 e validar output

**Testes para S-1005**:
1. `TestS1005MarshalXML_campos_obrigatorios` — Validar campos obrigatórios
2. `TestS1005MarshalXML_referencia_S1000` — Validar que ideEmpregador de S-1005 corresponde a S-1000
3. `TestS1005MarshalXML_tipos_especiais` — Validar serialização de tipos
4. `TestS1005MarshalXML_omitempty` — Validar omitempty para campos opcionais
5. `TestS1005ExemploMatrizValido` — Marshal do exemplo S-1005 (Matriz)
6. `TestS1005ExemploFilialValido` — Marshal do exemplo S-1005 (Filial)

**Validações comuns**:
- Verificar que tags XML estão em camelCase correto
- Verificar que valores com tipos especiais (CNPJ, DataApenas) são serializados corretamente
- Comparar XML gerado com expected (usando regexp ou XML parser)
- Verificar que nenhum campo vazio aparece no XML (se tiver tag `omitempty`)

---

## Relevant Files

| Arquivo | Propósito |
|---------|-----------|
| [esocial/internal/model/empregador/xml/](esocial/internal/model/empregador/xml/) | Pasta destino (vazia atualmente) |
| [esocial/pkg/types/cnpj.go](esocial/pkg/types/cnpj.go) | Tipo CNPJ com validação alfanumérica (14 dígitos) |
| [esocial/pkg/types/cpf.go](esocial/pkg/types/cpf.go) | Tipo CPF com validação (11 dígitos) |
| [esocial/pkg/types/dataapenas.go](esocial/pkg/types/dataapenas.go) | Tipo Data em formato YYYY-MM-DD |
| [esocial/pkg/types/simnao.go](esocial/pkg/types/simnao.go) | Tipo SimNao (1=Sim, 2=Não) |
| [esocial/pkg/types/decimalvalor.go](esocial/pkg/types/decimalvalor.go) | Tipo para valores monetários com casas decimais |
| [S1000_S1005_ESTRUTURA.md](S1000_S1005_ESTRUTURA.md) | Referência de campos, obrigatoriedade, validações e erros comuns |

---

## Verification Checklist

- [ ] Estruturas compilam sem erros: `go build ./esocial/...`
- [ ] Marshal de exemplo S-1000 gera XML válido com tags corretas (tpInsc, nrInsc, razaoSocial, etc.)
- [ ] Marshal de exemplo S-1005 gera XML válido
- [ ] Campos `omitempty` funcionam (campos vazios não aparecem no XML)
- [ ] Ordem de elementos segue o leiaute oficial (ideEmpregador antes de infoCadastro, etc.)
- [ ] Tipos especiais (CNPJ, DataApenas, SimNao) serializam corretamente
- [ ] Testes passam: `go test ./esocial/internal/model/empregador/xml/...`
- [ ] Cobertura de testes >= 80%

---

## Decisions

| Decisão | Justificativa |
|---------|---------------|
| **Reutilizar pkg/types** | CNPJ/CPF/DataApenas/SimNao oferecem validação tipada e serialização XML consistente |
| **Tipos auxiliares em types.go** | Facilita reutilização de Endereco e Contato entre S-1000 e S-1005 |
| **Um arquivo por evento** | Mantém organização limpa; s1000_*.go e s1005_*.go são facilmente localizáveis |
| **Examples em arquivo separado** | Evita poluição de testes com dados; exemplo é útil para docs/integrações |
| **Tags XML em camelCase** | Segue convenção oficial do leiaute eSocial |
| **Estruturas aninhadas** | Reflete hierarquia XML; facilita navegação e validação |

---

## Further Considerations

1. **Campos opcionais com `omitempty`?**  
   ✅ **Sim**, todos os campos não-obrigatórios devem ter `xml:",omitempty"` para não aparecer no XML quando vazios.

2. **CNAE como tipo dedicado ou string?**  
   ✅ **String simples** "6201500" (7 dígitos numéricos). Validação será feita no construtor ou via função de validação separada.

3. **Namespace XML na raiz?**  
   ⚠️ **Verificar leiaute oficial**. eSocial pode exigir `xmlns` na raiz:  
   ```xml
   <EvtInfoEmpregador xmlns="http://www.esocial.gov.br/schema/evt/evtInfoEmpregador/02.05.00">
   ```
   Por enquanto, manter estrutura sem namespace; adicionar se validações falharem.

4. **ID do evento (atributo `id`)?**  
   ⚠️ **Verificar leiaute**. Alguns eventos eSocial requerem atributo `id`:  
   ```xml
   <EvtInfoEmpregador id="ID001">
   ```
   Adicionar como campo struct com tag XML se necessário: `Id string \`xml:"id,attr"\``

5. **Validação de CNPJ/CEP/Município em tempo de marshal?**  
   ✅ **Não** — deixar para o construtor/validador. Marshal apenas serializa o que já foi validado.

---

## Implementation Order

1. ✅ Criar `types.go` (tipos auxiliares: Endereco, Contato, constantes)
2. ✅ Criar `s1000_info_empregador.go` (estrutura S-1000)
3. ✅ Criar `s1005_info_estab.go` (estrutura S-1005)
4. ✅ Criar `examples.go` (funções ExemploS1000Empresa, ExemploS1005Matriz, ExemploS1005Filial)
5. ✅ Criar `s1000_info_empregador_test.go` (testes S-1000)
6. ✅ Criar `s1005_info_estab_test.go` (testes S-1005)
7. ✅ Validar compilação e testes
8. ✅ Documentar (comentários Go e README se necessário)

---

## Notes

- Campos de data usam `DataApenas` que formata automaticamente para "2006-01-02"
- Campos com indicadores (S/N) usam `SimNao` que serializa como 1 ou 2
- CNPJ/CPF são tipos `string` com métodos MarshalXML que removem formatação
- Endereço é tipo reutilizável; não duplicar em S-1000 e S-1005
- Exemplos servem tanto para testes quanto para documentação de uso

