# EnerPredict

EnerPredict é uma aplicação projetada para prever o consumo de energia com base em dados históricos e variáveis específicas. A aplicação possui um backend em Go com Gin e um sistema de previsão utilizando Python.

# Funcionalidades

- Cadastro de Usuários: Permite que novos usuários se registrem e façam login na aplicação.
- Registro de Consumo: Usuários podem adicionar dados de consumo de energia.
- Previsão de Consumo: Gera previsões de consumo de energia com base em dados fornecidos.
- Consulta de Dados: Visualiza dados de consumo e informações do usuário.

# Tecnologias Utilizadas

- Backend: Go, Gin
- Banco de Dados: PostgreSQL
- Modelagem e Previsão: Python, scikit-learn, joblib

# Configuração do Ambiente

Configure as variáveis de ambiente no seu arquivo .env:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=ener_predict

# Configuração do Banco de Dados

- Crie e configure seu banco de dados PostgreSQL.
- Execute as migrações para criar as tabelas necessárias:

```go
go run main.go migrate
```

# Instalação e Execução

- Clone o repositório
- Configure e ative o ambiente virtual Python:

```python
python3 -m venv .venv
source .venv/bin/activate
```
- Instale as dependências Python:

```python
pip install -r requirements.txt
```
- Instale as dependências Go:

```go
go mod tidy
```
- Treine o modelo:

```python
python3 scripts/train_model.py

```

- Execute a aplicação:

```go
go run main.go
```

# Contribuindo

Se você deseja contribuir para o projeto, por favor, siga estas etapas ;) <3

- Faça um fork do repositório.
- Crie uma branch para sua feature (git checkout -b feature/nome-da-feature).
- Faça suas alterações e adicione testes, se necessário.

