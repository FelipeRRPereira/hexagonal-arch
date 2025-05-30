# Arquitetura Hexagonal (Ports and Adapters)

> Conteúdo baseado na apresentação de Wesley Willians

## Introdução

A **arquitetura hexagonal** propõe uma separação clara entre a complexidade **de negócio** e a complexidade **técnica**. Quanto mais desacoplado for o sistema, maior será a liberdade para realizar mudanças técnicas ao longo do tempo, sem impactar o núcleo da aplicação.

---

## Ciclo de Vida de um Projeto

É comum que projetos negligenciem a arquitetura nos estágios iniciais. Porém, mesmo que o escopo e a importância de um software mudem com o tempo, é essencial, em algum momento, refletir sobre:

- **Quais problemas o software resolve**
- **Quais são as partes mais importantes do sistema**

Isso garante um **crescimento sustentável**, **menor dívida técnica** e mais clareza no futuro do projeto.

---

## Princípios e Reflexões

### Pontos importantes sobre arquitetura:
- Crescimento sustentável
- Software precisa se pagar com o tempo
- Software deve ser **desenhado por você**, não pelo seu framework
- Peças devem se **encaixar** e poder ser **substituídas**

> "Arquitetura diz respeito com o futuro do seu software. CRUD qualquer um faz!"

### Reflexões importantes:
- **Visão de Futuro**: Comece um projeto pensando além de uma simples POC
- **Limites bem definidos**: Evite sobrecarga de responsabilidades
- **Troca de componentes**: O core deve ser estável, adaptadores mutáveis
- **Escalabilidade**: Planeje escala horizontal desde o início
- **Otimizações frequentes**: Mantenha o projeto saudável ao longo do tempo
- **Preparado para mudanças**: Mudanças técnicas sem comprometer o domínio

---

## Questões Importantes para se Fazer:

- Está sendo **doloroso** desenvolver melhorias?
- Isso poderia ter sido evitado?
- O software **está se pagando**?
- O cliente será impactado negativamente por mudanças bruscas?
- **Em que momento tudo se perdeu?**
- Se você fosse novo na equipe, **julgaria o código existente**?

---

## Arquitetura vs Design

> “Atividades relacionadas à arquitetura de software são sempre de design. Entretanto, nem toda atividade de design é arquitetural.” — Elemar Jr.

- **Arquitetura**: Trata das **decisões estruturais** que impactam a longevidade, escalabilidade e alinhamento com os objetivos do negócio
- **Design**: Cuida da qualidade interna das unidades de software. Ex: princípios **SOLID**
- **Exemplo**: O SOLID melhora a qualidade de classes, mas não trata diretamente da arquitetura de um sistema

---

## Arquitetura Hexagonal (Ports and Adapters)

A **arquitetura hexagonal**, também chamada de **Ports and Adapters**, organiza a aplicação ao redor de seu **núcleo de negócio**, isolando-o de detalhes técnicos como bancos de dados, APIs, filas, etc.

### Conceitos-chave:

- **Ports (Portas)**: Interfaces que representam as ações que o domínio pode realizar ou receber
- **Adapters (Adaptadores)**: Implementações técnicas que conversam com o mundo externo (como bancos, HTTP, filas)

---

## Benefícios

- **Desacoplamento** entre domínio e tecnologia
- **Facilidade de testes**
- **Substituição de dependências técnicas** com baixo impacto
- **Pronto para microserviços** no futuro

---

## Exemplos de Componentes Externos

- Banco de Dados
- Uploads
- Logs
- Cache
- APIs (REST/GraphQL)
- Filas de Mensagens
- Interfaces de CLI, Web ou Mobile

---

## Princípio de Inversão de Dependência (DIP)

> Parte do S.O.L.I.D. — base para arquitetura hexagonal

- Módulos de alto nível **não devem depender** de módulos de baixo nível
- Ambos devem depender de **abstrações**
- Abstrações **não devem depender de detalhes**
- Detalhes devem depender de **abstrações**

---

## Conclusão

Adotar a arquitetura hexagonal é uma decisão que promove longevidade e flexibilidade no software. Ao proteger a complexidade de negócio da complexidade técnica, você garante um projeto saudável, sustentável e pronto para crescer.
