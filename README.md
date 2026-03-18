# Sistema Público de Escrituração Digital (SPED)

O **SPED** é um pilar fundamental da modernização da relação entre o fisco e os contribuintes no Brasil. Instituído pelo Decreto nº 6.022/2007, ele substitui a emissão de livros e documentos em papel por arquivos digitais, padronizando o envio de informações fiscais, previdenciárias e contábeis.

## 🚀 Objetivos do Projeto

  * **Integração:** Unificar o compartilhamento de informações entre as esferas federal, estadual e municipal.
  * **Racionalização:** Reduzir custos acessórios para as empresas (armazenamento físico e burocracia).
  * **Fiscalização:** Aumentar a eficácia da auditoria através do cruzamento de dados em tempo real.

## 📦 Módulos do SPED

O ecossistema é dividido em frentes que abrangem documentos fiscais, escrituração contábil e obrigações trabalhistas.

### 1\. Documentos Fiscais Eletrônicos (DF-e)

  * **NF-e (Nota Fiscal Eletrônica):** Modelo 55, substitui as notas de mercadorias modelos 1 e 1A.
  * **NFS-e (Nota Fiscal de Serviços Eletrônica):** Focada na incidência de ISS (âmbito municipal).
  * **CT-e (Conhecimento de Transporte Eletrônico):** Para prestação de serviços de transporte de cargas.
  * **BP-e (Bilhete de Passagem Eletrônico):** Transporte de passageiros.
  * **MDF-e (Manifesto Eletrônico de Documentos Fiscais):** Vincula documentos de carga transportados em um veículo.

### 2\. Escriturações Digitais

  * **ECD (Escrituração Contábil Digital):** Substitui os livros Diário e Razão.
  * **ECF (Escrituração Contábil Fiscal):** Substitui a DIPJ, focada na base de cálculo do IRPJ e CSLL.
  * **EFD ICMS IPI:** Registra a movimentação de mercadorias e a apuração de impostos estaduais e federais.
  * **EFD Contribuições:** Focada no PIS/Pasep e na Cofins.
  * **EFD-Reinf:** Escrituração de retenções e outras informações fiscais (complementar ao eSocial).

### 3\. Esfera Trabalhista e Previdenciária

  * **eSocial:** Unifica o envio de dados sobre empregados, folha de pagamento, acidentes de trabalho e aviso prévio.

### 4\. Outros Projetos

  * **e-Financeira:** Informações sobre operações financeiras junto à Receita Federal.
  * **Central de Balanços:** Repositório público de demonstrações contábeis.

-----

## 🔗 Links Úteis e Documentação

Para desenvolvedores e analistas fiscais, o acompanhamento dos manuais de orientação e esquemas XML (XSD) é indispensável:

  * [**Portal Oficial do SPED**](http://sped.rfb.gov.br/): Ponto central para manuais, guias práticos e downloads de validadores (PVA).
  * [**Portal da Nota Fiscal Eletrônica (NF-e)**](http://www.nfe.fazenda.gov.br/): Especificações técnicas, web services e consultas de chave de acesso.
  * [**Portal do eSocial**](https://www.gov.br/esocial/): Documentação técnica sobre os eventos de S-1000 a S-5013.
  * [**ENAT (Encontro Nacional de Administradores Tributários)**](https://www.google.com/search?q=https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/enat): Discussões sobre a integração nacional do sistema.

## 🛠️ Aspectos Técnicos

A implementação do SPED exige conformidade com:

1.  **Certificação Digital:** Uso de certificados ICP-Brasil (A1 ou A3) para assinatura digital.
2.  **Padrão XML:** A maioria das transmissões utiliza SOAP Web Services e arquivos estruturados em XML.
3.  **Layouts:** Cada módulo possui um Guia Prático com a definição de blocos, registros e campos obrigatórios.
