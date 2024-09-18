import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error
import joblib
import psycopg2
from sqlalchemy import create_engine


DB_HOST = 'localhost'
DB_PORT = '5432'
DB_USER = 'postgres'
DB_PASSWORD = 'password'
DB_NAME = 'ener_predict'

def load_data():
    connection_string = f'postgresql://{DB_USER}:{DB_PASSWORD}@{DB_HOST}:{DB_PORT}/{DB_NAME}'
    engine = create_engine(connection_string)
    
    query = """
    SELECT
        EXTRACT(MONTH FROM consumption_date) AS month,
        EXTRACT(DAY FROM consumption_date) AS day,
        temperature,
        usage_hours,
        amount AS energy_consumption
    FROM consumptions
    """
    
    data = pd.read_sql(query, engine)
    return data

def train_model(data):
    X = data[["month", "day", "temperature", "usage_hours"]]
    y = data["energy_consumption"]

    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

    model = LinearRegression()
    model.fit(X_train, y_train)

    predictions = model.predict(X_test)
    mse = mean_squared_error(y_test, predictions)
    print(f"Mean Squared Error: {mse}")

    return model

def save_model(model, model_path="models/energy_forecast_model.pkl"):
    joblib.dump(model, model_path)
    print(f"Modelo salvo em {model_path}")

if __name__ == "__main__":
    data = load_data()
    model = train_model(data)
    save_model(model)
