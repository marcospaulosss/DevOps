[![CircleCI](https://circleci.com/gh/estrategiaconcursos/backend/tree/master.svg?style=svg&circle-token=f81b67b0e0f74c75ad72c67def6161f1cec5682f)](https://circleci.com/gh/estrategiaconcursos/backend/tree/master)

# Backend

## Serviços

- **severino**: API Gateway/Backend for Front-end (BFF)
- **elearning**: Audio (albums, faixas, tags e seções)
- **accounts**: Autenticação, Autorização e contas de usuários
- **ecommerce**: Produtos, Pedidos e pagamentos

## Setup

### Baixe o binário do protobuf

Faça o download do compilador `protoc` e coloque no `$PATH`

```sh
# MacOS
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-osx-x86_64.zip

# Linux
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip

unzip protoc-3.7.1*.zip
sudo mv bin/protoc /usr/local/bin
```

Instale as libs necessárias para trabalhar com gRPC:

```sh
go get google.golang.org/grpc
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/vektra/mockery/cmd/mockery
go get -u github.com/onsi/ginkgo/ginkgo
```

**Crie ou atualize os arquivos proto**

Você deve executar os comandos abaixo sempre que houver mudanças nos arquivos `.proto`.

```sh
make proto
```

### Arquivo de configuração

As configurações estão dentro do arquivo `config.yaml` dentro da pasta de cada serviço.

### Instale o plugin Editorconfig

Esse plugin garante a formatação do código após salva-lo no editor.

### Executando

#### Cada aplicação de forma individual (RECOMENDADO)

Abra um terminal para cada aplicação dentro de `apps` e execute o `make run`. Ex.:

```sh
cd severino
make run
```

#### Como containers

Ideal para o frontend subir toda estrutura localmente sem precisar alterar o código ou instalar o Go e suas dependências.

```
docker-compose up -d
```

## Documentação da API 

Sempre que houver mudança na API do Severino, execute o comando `make docs` dentro de `severino` para gerar o index.html da documentação. Esse arquivo é copiado para o S3.

- http://apidocs.estrategia.dev/index.html

## Libs

Libs adicionadas dentro de `backend/libs`:

- **json**: Utiliza a lib json iterator que possui uma performance melhor do que a stdlib `encoding/json`. Também adiciona suporte ao formato de timestamptz do Posrtgres.
- **logger**: Utiliza o logrus e adiciona suporte para RSyslog.
- **testing**: Executa queries contidas em arquivos SQL.
- **configuration**: Utiliza o viper para obter dados a partir de arquivos de configuração e variáveis de ambiente.
- **httpvalidator**: Encapsula o validador utilizado pelo Echo. 
- **remoteprocedurecall**: Para conexões gRPC
- **databases**: Implementação do Postgres

## Mais informações

Veja o `README.md` dentro da pasta de cada serviço.
