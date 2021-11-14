import pandas as pd

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

        print(matrix)