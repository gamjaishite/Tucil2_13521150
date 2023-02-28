import matplotlib.pyplot as plt
import pandas as pd

# Read data from CSV
df = pd.read_csv('points.csv')
df2 = pd.read_csv('ppoints.csv')

# Initializing the frame
plt.figure(figsize=(15, 13))
ax = plt.axes(projection='3d')

# Drawing point to the frame
for i in range(len(df.index)):
    inside = False
    for j in range(len(df2.index)):
        if (str(df['X'][i]) == str(df2['X'][j]) and str(df['Y'][i]) == str(df2['Y'][j]) and str(df['Z'][i]) == str(df2['Z'][j])):
            inside = True
            break
    if inside:
        ax.scatter(df['X'][i], df['Y'][i], df['Z'][i], color='red')
    else:
        ax.scatter(df['X'][i], df['Y'][i], df['Z'][i], color='blue')
for i in range(1, len(df2.index), 2):
    ax.plot(df2['X'][i-1:i+1], df2['Y'][i-1:i+1],
            df2['Z'][i-1:i+1], color='red')

# Saving plotting result to file
plt.savefig("../../output/result.png")

# Showing the visualizer
plt.show()
