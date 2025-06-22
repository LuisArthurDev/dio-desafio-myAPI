
# 🐹 API REST em Go (Golang) com Gorilla Mux e JSON como Banco de Dados

Este é um projeto simples de **API RESTful** feito em **Go**, utilizando o **Gorilla Mux** como roteador e um **arquivo JSON local (`users.json`) para persistência de dados**.

---

## 🚀 Funcionalidades da API

| Método | Endpoint         | Descrição                    |
|--------|------------------|------------------------------|
| GET    | `/users`         | Lista todos os usuários       |
| GET    | `/users/{id}`    | Retorna um usuário por ID     |
| POST   | `/users`         | Cria um novo usuário          |
| PUT    | `/users/{id}`    | Atualiza um usuário existente |
| DELETE | `/users/{id}`    | Deleta um usuário             |

---

## 🗃️ Estrutura do Usuário (JSON)

Cada usuário tem esta estrutura:

```json
{
  "id": 1,
  "name": "Luis",
  "age": 28
}
```

---

## 🧑‍💻 Como Rodar o Projeto

### 1. Clone o repositório:

```bash
git clone https://github.com/seunome/seu-repo.git
cd seu-repo
```

### 2. Instale as dependências:

> É necessário ter o **Go instalado** na sua máquina.

```bash
go mod init nome-do-seu-modulo
go get github.com/gorilla/mux
```

---

### 3. Rode o servidor:

```bash
go run main.go
```

O servidor ficará rodando em:

```
http://localhost:8000
```

---

## 🧪 Exemplos de Requisições HTTP

### Criar um novo usuário (POST):

#### Linux / Mac / Git Bash:

```bash
curl -X POST http://localhost:8000/users -H "Content-Type: application/json" -d '{"name":"Luis","age":28}'
```

#### Windows CMD:

```bash
curl -X POST http://localhost:8000/users -H "Content-Type: application/json" -d "{\"name\":\"Luis\",\"age\":28}"
```

#### Windows PowerShell:

```powershell
curl -X POST http://localhost:8000/users -H "Content-Type: application/json" -Body '{"name":"Luis","age":28}' -ContentType "application/json"
```

Ou usando Invoke-RestMethod no PowerShell:

```powershell
Invoke-RestMethod -Uri http://localhost:8000/users -Method POST -Body '{"name":"Luis","age":28}' -ContentType "application/json"
```

---

### Listar usuários (GET):

```bash
curl http://localhost:8000/users
```

---

### Buscar usuário por ID (GET):

```bash
curl http://localhost:8000/users/1
```

---

### Atualizar usuário (PUT):

```bash
curl -X PUT http://localhost:8000/users/1 -H "Content-Type: application/json" -d '{"name":"Luis Updated","age":29}'
```

---

### Deletar usuário (DELETE):

```bash
curl -X DELETE http://localhost:8000/users/1
```

---

## 💾 Persistência de Dados

- Os dados dos usuários são salvos no arquivo `users.json` na raiz do projeto.
- O arquivo é criado automaticamente após a primeira alteração (POST, PUT ou DELETE).

---

## ✅ Requisitos

- Go 1.16 ou superior
- Gorilla Mux

---

## 📌 Próximos passos (melhorias futuras)

- Adicionar CORS para permitir acessos de front-end externos
- Usar banco de dados real (MySQL, PostgreSQL, SQLite)
- Criar um front-end para consumir essa API

---

## 🛠️ Desenvolvido por
**Luis Arthur**  
[LinkedIn](https://www.linkedin.com/in/luisarthurrib/)

---

## 📝 Licença

Este projeto é livre para uso educacional.

---
