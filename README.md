# Pirâmide de Testes em Go
Decisões técnicas e arquiteturais
Golang foi a linguagem escolhida para conduzir o desenvolvimento deste projeto devido a uma apresentação para demonstrar o uso da teoria da pirâmide de testes em Go, além de oferecer a criação de executáveis de código para diversas plataformas. Com dessas características, o Go apresenta atributos gerais, como simplicidade e desempenho, que se destacam como elementos-chave em uma variedade de aplicações.

Ports and Adapter é a arquitetura utilizada como design para o projeto por trazer simplicidade, fácil manutenibilidade e a capacidade de extensão para futuras funcionalidades. Um padrão flexível e desacoplado leva a separação de responsabilidades, tornando o software coeso e manutenível.

# Dependências
- gomock: https://github.com/golang/mock
    - Gomock é uma estrutura de simulação para linguagem Go.
    - Utilizado na geração de código que simula objetos para testes.
- testify: https://github.com/stretchr/testify
    - Fornece muitas ferramentas para testar o código conforme a necessidade do desenvolvedor.
    - Utilizado para fazer asserções nos testes.
- testcontainers: https://github.com/testcontainers/testcontainers-go
    - Fornece o ambiente para realizarmos o teste de integração com o banco de dados.
    - Utilizado para criar um container de teste.

# Executando projeto
 - 1º Construíndo build da aplicação conteinerizada
  
    ``make run``

- 2º Executando aplicação linha a linha ou através de um arquivo
  
  ``make run ou make run < input.txt``

- 3º Executando todos os testes

  ``make test``

- 4° Verificação da cobertura de testes de unidade

  ``make test-cov``

# Outros comandos
- Gerando/atualizando mocks
 
  ``make build-mocks``

- Download das dependências
  
  ``make dependencies``

- Gerando binário local

  ``make build``

# Notas adicionais
- A cobertura de teste pode ser visualizada através de um arquivo **cover.html** na raiz do projeto
- O arquivo Makefile é um script usado para automatizar a compilação e outras tarefas relacionadas ao desenvolvimento
- Para rodar o projeto localmente é necessário instalar o Golang: https://go.dev/doc/install
- Para mais informações sobre o projeto segue o link da [apresentação](https://docs.google.com/presentation/d/1hgRWaVf9-1eew_uEl5JfSvNB9S18pBGmvhMaPhJ1bcA/edit?usp=sharing)
