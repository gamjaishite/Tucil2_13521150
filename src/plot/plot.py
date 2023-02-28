import matplotlib.pyplot as plt
import numpy as np
import pandas as pd

df = pd.read_csv('points.csv')
df2 = pd.read_csv('ppoints.csv')

plt.figure(figsize=(15, 13))
ax = plt.axes(projection='3d')
ax.scatter(df['X'], df['Y'], df['Z'])
for i in range(1, len(df2.index), 2):
    ax.plot(df2['X'][i-1:i+1], df2['Y'][i-1:i+1],
            df2['Z'][i-1:i+1], color='red')
    ax.scatter(df2['X'][i-1:i+1], df2['Y'][i-1:i+1],
               df2['Z'][i-1:i+1], color='red')
plt.savefig("result.png")
plt.show()
