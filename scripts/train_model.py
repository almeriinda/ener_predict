import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error
import joblib
import os

# Função para carregar dados de consumo de energia do banco de dados ou de um CSV
def load_data(file_path="data/consumption_data.csv"):
    if os.path.exists(file_path):
        data = pd.read_csv(file_path)
    else:
        raise FileNotFoundError("Arquivo de dados não encontrado")
    return data

# Função para treinar o modelo de previsão
def train_model(data):
    # Separar dados de entrada (X) e saída (y)
    X = data[["month", "day", "temperature", "usage_hours"]]
    y = data["energy_consumption"]

    # Dividir os dados em treino e teste
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    # Treinar o modelo de regressão linear
    model = LinearRegression()
    model.fit(X_train, y_train)

    # Avaliar o modelo
    predictions = model.predict(X_test)
    mse = mean_squared_error(y_test, predictions)
    print(f"Mean Squared Error: {mse}")

    return model

# Função para salvar o modelo treinado
def save_model(model, model_path="models/energy_forecast_model.pkl"):
    joblib.dump(model, model_path)
    print(f"Modelo salvo em {model_path}")

# Função principal
if __name__ == "__main__":
    # Carregar os dados
    data = load_data()

    # Treinar o modelo
    model = train_model(data)

    # Salvar o modelo
    save_model(model)
