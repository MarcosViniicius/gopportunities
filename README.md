# Gopportunities - Job Opportunities API

<div align="center">

![Gopportunities](assets\golang.png)

![Go](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=flat-square&logo=go)
![Gin](https://img.shields.io/badge/Gin-1.11.0-00ADD8?style=flat-square)
![SQLite](https://img.shields.io/badge/SQLite-3-003B57?style=flat-square&logo=sqlite)
![Swagger](https://img.shields.io/badge/Swagger-2.0-85EA2D?style=flat-square&logo=swagger)

A simple REST API to manage job opportunities. Create, read, update, and delete job openings easily.

**[Português](#português) | [English](#english)**

</div>

---

## English

### What is Gopportunities?

Gopportunities is a REST API built with Go that helps you manage job opportunities. You can create new job openings, list all opportunities, search for specific jobs, update job information, and delete outdated openings.

### Main Features

- ✅ Create new job opportunities
- ✅ List all job opportunities
- ✅ Find a specific job by ID
- ✅ Update job information
- ✅ Delete job opportunities
- ✅ Automatic API documentation with Swagger
- ✅ SQLite database (no external database needed)
- ✅ Simple and clean code

### Technologies Used

| Technology | Purpose |
|-----------|---------|
| **Go 1.25.5** | Programming language |
| **Gin** | Web framework for building HTTP servers |
| **GORM** | Database ORM for easy database operations |
| **SQLite** | Simple file-based database |
| **Swagger** | Automatic API documentation |

### How to Run

#### Step 1: Install Go

Download and install Go from [golang.org](https://golang.org/dl/)

#### Step 2: Clone or Download the Project

```bash
git clone https://github.com/MarcosViniicius/gopportunities
cd gopportunities
```

#### Step 3: Install Dependencies

```bash
go mod download
```

#### Step 4: Run the Project

Run with Swagger documentation generation:

```bash
make run-with-docs
```

Or run directly:

```bash
go run main.go
```

The API will start at: `http://localhost:8080`

#### Step 5: View API Documentation

Open your browser and go to:

```
http://localhost:8080/swagger/index.html
```

You will see all available endpoints with examples.

### API Endpoints

All endpoints start with `/api/v1/`

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/opening` | Create a new job opportunity |
| GET | `/opening?id=1` | Get a specific job opportunity by ID |
| GET | `/openings` | List all job opportunities |
| PUT | `/opening?id=1` | Update a job opportunity |
| DELETE | `/opening?id=1` | Delete a job opportunity |

### Example: Create a New Job Opportunity

Using curl:

```bash
curl -X POST http://localhost:8080/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Go Developer",
    "company": "Tech Company",
    "location": "São Paulo, Brazil",
    "remote": true,
    "link": "https://company.com/jobs/go-developer",
    "salary": 120000
  }'
```

Response:

```json
{
  "message": "operation from handler: create-opening successful",
  "data": {
    "id": 1,
    "createdAt": "2026-01-10T12:00:00Z",
    "updatedAt": "2026-01-10T12:00:00Z",
    "role": "Go Developer",
    "company": "Tech Company",
    "location": "São Paulo, Brazil",
    "remote": true,
    "link": "https://company.com/jobs/go-developer",
    "salary": 120000
  }
}
```

### Available Commands

| Command | Description |
|---------|-------------|
| `make run-with-docs` | Run the project with Swagger documentation generation |
| `make run` | Run the project without generating docs |
| `make build` | Build a binary file (executable) |
| `make docs` | Generate Swagger documentation only |
| `make test` | Run all tests |
| `make clean` | Remove binary and docs folder |

### Project Structure

```
gopportunities/
├── main.go              # Application entry point
├── config/              # Configuration files
│   ├── config.go
│   ├── logger.go
│   └── sqlite.go
├── router/              # HTTP routes
│   ├── router.go
│   └── routes.go
├── handler/             # API request handlers
│   ├── createOpening.go
│   ├── listOpenings.go
│   ├── showOpening.go
│   ├── updateOpening.go
│   ├── deleteOpening.go
│   ├── request.go
│   └── response.go
├── schemas/             # Database models
│   └── opening.go
├── docs/                # Generated API documentation
├── db/                  # Database file (auto-created)
└── makefile             # Build commands
```

### Database

The project uses SQLite, which automatically creates a database file at `./db/main.db` when you first run the application. No configuration needed.

### Troubleshooting

**Problem: Port 8080 is already in use**

Solution: Change the port using the PORT environment variable:

```bash
export PORT=3000
go run main.go
```

**Problem: Swagger documentation not showing**

Solution: Make sure to run:

```bash
make run-with-docs
```

This command regenerates the Swagger documentation before starting the API.

### License

This project is open source and available under the MIT License.

---

## Português

### O que é Gopportunities?

Gopportunities é uma API REST feita em Go que ajuda você a gerenciar oportunidades de emprego. Você pode criar novas vagas, listar todas as oportunidades, procurar trabalhos específicos, atualizar informações de vagas e deletar oportunidades desatualizadas.

### Principais Recursos

- ✅ Criar novas oportunidades de emprego
- ✅ Listar todas as oportunidades de emprego
- ✅ Encontrar uma vaga específica pelo ID
- ✅ Atualizar informações de vagas
- ✅ Deletar oportunidades de emprego
- ✅ Documentação automática da API com Swagger
- ✅ Banco de dados SQLite (sem necessidade de banco de dados externo)
- ✅ Código simples e limpo

### Tecnologias Utilizadas

| Tecnologia | Propósito |
|-----------|-----------|
| **Go 1.25.5** | Linguagem de programação |
| **Gin** | Framework web para construir servidores HTTP |
| **GORM** | ORM de banco de dados para operações fáceis |
| **SQLite** | Banco de dados simples baseado em arquivo |
| **Swagger** | Documentação automática da API |

### Como Executar

#### Passo 1: Instale Go

Baixe e instale Go de [golang.org](https://golang.org/dl/)

#### Passo 2: Clone ou Baixe o Projeto

```bash
git clone https://github.com/MarcosViniicius/gopportunities
cd gopportunities
```

#### Passo 3: Instale as Dependências

```bash
go mod download
```

#### Passo 4: Execute o Projeto

Execute com geração de documentação Swagger:

```bash
make run-with-docs
```

Ou execute diretamente:

```bash
go run main.go
```

A API iniciará em: `http://localhost:8080`

#### Passo 5: Veja a Documentação da API

Abra seu navegador e vá para:

```
http://localhost:8080/swagger/index.html
```

Você verá todos os endpoints disponíveis com exemplos.

### Endpoints da API

Todos os endpoints começam com `/api/v1/`

| Método | Endpoint | Descrição |
|--------|----------|-----------|
| POST | `/opening` | Criar uma nova oportunidade de emprego |
| GET | `/opening?id=1` | Obter uma oportunidade específica pelo ID |
| GET | `/openings` | Listar todas as oportunidades de emprego |
| PUT | `/opening?id=1` | Atualizar uma oportunidade de emprego |
| DELETE | `/opening?id=1` | Deletar uma oportunidade de emprego |

### Exemplo: Criar uma Nova Oportunidade de Emprego

Usando curl:

```bash
curl -X POST http://localhost:8080/api/v1/opening \
  -H "Content-Type: application/json" \
  -d '{
    "role": "Desenvolvedor Go",
    "company": "Empresa de Tecnologia",
    "location": "São Paulo, Brasil",
    "remote": true,
    "link": "https://empresa.com/vagas/desenvolvedor-go",
    "salary": 120000
  }'
```

Resposta:

```json
{
  "message": "operation from handler: create-opening successful",
  "data": {
    "id": 1,
    "createdAt": "2026-01-10T12:00:00Z",
    "updatedAt": "2026-01-10T12:00:00Z",
    "role": "Desenvolvedor Go",
    "company": "Empresa de Tecnologia",
    "location": "São Paulo, Brasil",
    "remote": true,
    "link": "https://empresa.com/vagas/desenvolvedor-go",
    "salary": 120000
  }
}
```

### Comandos Disponíveis

| Comando | Descrição |
|---------|-----------|
| `make run-with-docs` | Executar o projeto com geração de documentação Swagger |
| `make run` | Executar o projeto sem gerar docs |
| `make build` | Construir um arquivo binário (executável) |
| `make docs` | Gerar documentação Swagger apenas |
| `make test` | Executar todos os testes |
| `make clean` | Remover arquivo binário e pasta docs |

### Estrutura do Projeto

```
gopportunities/
├── main.go              # Ponto de entrada da aplicação
├── config/              # Arquivos de configuração
│   ├── config.go
│   ├── logger.go
│   └── sqlite.go
├── router/              # Rotas HTTP
│   ├── router.go
│   └── routes.go
├── handler/             # Handlers de requisições da API
│   ├── createOpening.go
│   ├── listOpenings.go
│   ├── showOpening.go
│   ├── updateOpening.go
│   ├── deleteOpening.go
│   ├── request.go
│   └── response.go
├── schemas/             # Modelos de banco de dados
│   └── opening.go
├── docs/                # Documentação da API gerada
├── db/                  # Arquivo do banco de dados (auto-criado)
└── makefile             # Comandos de construção
```

### Banco de Dados

O projeto usa SQLite, que cria automaticamente um arquivo de banco de dados em `./db/main.db` quando você executa a aplicação pela primeira vez. Nenhuma configuração necessária.

### Solução de Problemas

**Problema: A porta 8080 já está em uso**

Solução: Mude a porta usando a variável de ambiente PORT:

```bash
export PORT=3000
go run main.go
```

**Problema: A documentação do Swagger não está aparecendo**

Solução: Certifique-se de executar:

```bash
make run-with-docs
```

Este comando regenera a documentação Swagger antes de iniciar a API.

### Licença

Este projeto é de código aberto e está disponível sob a Licença MIT.

---

<div align="center">

Made with ❤️ by [Marcos Vinícius](https://github.com/MarcosViniicius)

</div>
