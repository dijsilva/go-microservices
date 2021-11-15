import pickle
import pandas as pd
import os
from scipy.signal import savgol_filter

class SpectraPrediction:
    def pre_processing_samples(self, rows):
        # Order rows
        rows_ordered = sorted(rows, key = lambda key : key['Row'])
        if not rows_ordered[0].get('IsHeader', False):
            raise ValueError('First line not is header')
        
        lines = [line['Values'] for line in rows_ordered]
        matrix = []
        for line in lines:
            col_values = [col['Value'] for col in line]
            matrix.append(col_values)
        
        
        matrix = pd.DataFrame(matrix[1:], columns=matrix[0])
        return matrix
        
    def predict(self, input_data):
        
        rows = input_data.get('Rows', None)

        if (not rows) or (len(rows) <= 1):
            raise ValueError('Data of spectra not defiend')

        matrix = self.pre_processing_samples(rows)

        if matrix.shape[0] >= 2:
            matrix = matrix.mean(axis=0)
        
        print('Loading model')

        dir_path = os.path.dirname(os.path.realpath(__file__))
        with open(f"{dir_path}/model/randomforest_1.0.pkl", 'rb') as m:
            model = pickle.load(m)

        x = savgol_filter(matrix.values, polyorder=4, deriv=2, window_length=11)

        predictions = model.predict(x)
        predicao = self.getStringClassified(predictions[0])
        return predicao, predictions[0]

    def getStringClassified(self, prediction: int) -> str:
        classes = {
            1: 'Saudável',
            2: 'Câncer no endométrio',
            3: 'Câncer de ovário'
        }

        return classes[prediction]