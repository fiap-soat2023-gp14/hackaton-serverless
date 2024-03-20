# hackaton-serverless

## O que a Plataforma é capaz de fazer?
- Criar usuários no Cognito;
- Autenticar usuários no Cognito;

## Pré-requisitos e como rodar a aplicação
- AWS SAM CLI
- Go 1.21

## Estrutura do Projeto
```
├── src
│ ├── domain
│ ├── handlers
│ ├── infra
│ │ ├── adapters
│ │ └── settings
│ ├── models
│ └── services
```

## Instalação
Faça o download do repositório através do arquivo zip ou do terminal usando o git clone;

```bash
git clone https://github.com/fiap-soat2023-gp14/soat23-gp14-hackaton-serverless
```

Ao finalizar o download, acesse o diretório do projeto pelo seu terminal;

```bash
cd soat23-gp14-hackaton-serverless
```

## Build
```bash
sam build
```

## Local Testing
```bash
sam local start-api --env-vars test/env.json
```

## Running Tests
```bash
make tests
```