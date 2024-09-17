import joblib
import pandas as pd
import sys
import os

def load_model(model_path="models/energy_forecast_model.pkl"):
    if os.path.exists(model_path):
        model = joblib.load(model_path)
        print(f"Modelo carregado de {model_path}")
    else:
        raise FileNotFoundError("Modelo não encontrado")
    return model

def run_forecast(model, new_data):
    data = pd.DataFrame([new_data])
    forecast = model.predict(data)
    return forecast[0]

if __name__ == "__main__":
    if len(sys.argv) != 5:
        print("Uso: python run_forecast.py <month> <day> <temperature> <usage_hours>")
        sys.exit(1)

    month = int(sys.argv[1])
    day = int(sys.argv[2])
    temperature = float(sys.argv[3])
    usage_hours = float(sys.argv[4])

    new_data = {
        "month": month,
        "day": day,
        "temperature": temperature,
        "usage_hours": usage_hours
    }

    model = load_model()

    forecast = run_forecast(model, new_data)
    print(f"Previsão de consumo de energia: {forecast:.2f} kWh")
