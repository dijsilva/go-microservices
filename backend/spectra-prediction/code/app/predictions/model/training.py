import pickle
from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import train_test_split
import pandas as pd
from scipy.signal import savgol_filter
import numpy as np

df = pd.read_csv("", sep=";", decimal=",").dropna()

y = df["SPECTRA CLASS"]
x = df.iloc[:,1:]
x = savgol_filter(x, polyorder=4, deriv=2, window_length=11)

x_train, x_test, y_train, y_test = train_test_split(x, y, test_size=0.33, random_state=123)

randomforest = RandomForestClassifier(n_estimators=1000, random_state=123)
randomforest.fit(X=x_train, y=y_train)

file_name = 'randomforest_1.0.pkl'
with open('./randomforest_1.0.pkl', 'wb') as f:
    pickle.dump(randomforest, f)