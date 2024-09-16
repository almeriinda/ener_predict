import joblib
import pandas as pd
import sys
import os

# Função para carregar o modelo treinado
def load_model(model_path="models/energy_forecast_model.pkl"):
    if os.path.exists(model_path):
        model = joblib.load(model_path)
        print(f"Modelo carregado de {model_path}")
    else:
        raise FileNotFoundError("Modelo não encontrado")
    return model

# Função para executar a previsão
def run_forecast(model, new_data):
    # Converter dados para DataFrame
    data = pd.DataFrame([new_data])
    forecast = model.predict(data)
    return forecast[0]

# Função principal
if __name__ == "__main__":
    if len(sys.argv) != 5:
        print("Uso: python run_forecast.py <month> <day> <temperature> <usage_hours>")
        sys.exit(1)

    # Coletar dados de entrada a partir dos argumentos da linha de comando
    month = int(sys.argv[1])
    day = int(sys.argv[2])
    temperature = float(sys.argv[3])
    usage_hours = float(sys.argv[4])

    # Dados fornecidos pelo usuário
    new_data = {
        "month": month,
        "day": day,
        "temperature": temperature,
        "usage_hours": usage_hours
    }

    # Carregar o modelo treinado
    model = load_model()

    # Executar a previsão
    forecast = run_forecast(model, new_data)
    print(f"Previsão de consumo de energia: {forecast:.2f} kWh")
