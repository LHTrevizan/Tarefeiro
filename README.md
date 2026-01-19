#  📝 Tarefeiro — Gerenciador de Tarefas via CLI (Go)
Aplicação CLI desenvolvida em Go para gerenciamento de tarefas pessoais.
O objetivo do projeto é demonstrar domínio em:

- Desenvolvimento de CLIs em Go

- Uso da biblioteca Cobra

- Organização de código e separação de responsabilidades

- Persistência em arquivo JSON

- Validações e tratamento de erros

## ⚙️ Requisitos

- Go 1.22 ou superior

- Funciona em Windows, Linux e macOS

- Não utiliza CGO

📁 Estrutura do Projeto
```
tarefeiro/
├── cmd/
│   └── tarefeiro/          # CLI (Cobra)
│       ├── main.go
│       ├── add.go
│       ├── list.go
│       ├── show.go
│       ├── edit.go
│       ├── complete.go
│       └── delete.go
│
├── internal/
│   ├── model/              # Entidades e validações
│   │   └── task.go
│   │
│   ├── service/            # Regras de negócio
│   │   └── task_service.go
│   │
│   ├── repository/         # Acesso a dados do domínio
│   │   └── task_repository.go
│   │
│   └── infra/              # Infraestrutura / persistência
│       └── json_storage.go
│
├── data/
│   └── tasks.json          # Criado automaticamente
│
├── go.mod
└── README.md
```
## 📦 Instalação
### ✅ Importante

**Todos os comandos abaixo devem ser executados a partir da raiz do projeto (tarefeiro/).** 

## Opção 1 — Instalação recomendada (Go way 💙)
``` 
cd tarefeiro
go install ./cmd/tarefeiro 
```

Esse comando irá:

- Compilar a aplicação

- Instalar o binário tarefeiro em $GOPATH/bin

Após isso, o comando estará disponível globalmente:
```
tarefeiro --help
```
## Observação para Windows

Certifique-se de que o diretório abaixo está no PATH do sistema:
```
C:\Users\<seu-usuario>\go\bin
```

Caso contrário, o comando tarefeiro não será reconhecido no terminal.

## Opção 2 — Build manual (sem instalar)
```
cd tarefeiro
go build -o tarefeiro ./cmd/tarefeiro
```

**Windows**
```
.\tarefeiro add "Estudar Go"
```
**Linux / macOS**
```
./tarefeiro add "Estudar Go"
```
## 🧠 Persistência de Dados

As tarefas são armazenadas localmente em:
```
data/tasks.json
```
O diretório e o arquivo são criados automaticamente na primeira execução.

## 🚀 Uso da CLI
### ➕ Adicionar tarefa
```
tarefeiro add "Estudar Go" --priority high --tags dev,estudos
```

Flags disponíveis:

```--priority```: low | medium | high (default: medium)

```--tags```: lista separada por vírgula

### 📋 Listar tarefas
```
tarefeiro list
```
Com filtros:
```
tarefeiro list --status pending
tarefeiro list --priority high
tarefeiro list --text "titulo"
tarefeiro list --text "tags"
```
### 🔍 Exibir detalhes de uma tarefa
```
tarefeiro show <id>
```
### ✏️ Editar tarefa
```
tarefeiro edit <id> --title "Novo título" --priority low --tags nova,lista
```
### ✅ Marcar tarefa como concluída
```
tarefeiro complete <id>
```
### ❌ Remover tarefa
```
tarefeiro delete <id>
```
## 🧪 Testes
```
go test ./...
```
## 🛡️ Validações e Regras
- ```title``` é obrigatório
- ```priority``` aceita apenas: low, medium, high
- ```status``` é controlado internamente:
    - ```pending``` ao criar
    - ```done``` ao completar
- IDs são gerados como UUIDs
- Datas de criação, atualização e conclusão são controladas pela aplicação

## 🐳 Docker (opcional)

O projeto pode ser executado via Docker, caso desejado:
```
docker compose build
docker compose run tarefeiro add "estudar go"
docker compose run tarefeiro list
```

Docker é opcional e não é necessário para executar a aplicação.

## 🧠 Observações Técnicas

- O projeto não utiliza CGO, evitando problemas de build no Windows
- Separação clara de responsabilidades:
    - ```cmd```: interface CLI
    - ```service```: regras de negócio
    - ```infra```: persistência
    - ```model```: domínio e validações
- Estrutura segue o padrão recomendado pela comunidade Go para CLIs

## ✨ Autor

Projeto desenvolvido como desafio técnico em Go, com foco em clareza, simplicidade e boas práticas.
