import sys
import csv
import matplotlib.pyplot as plt

wave = sys.argv[1]
x = []
y = []
z = []
with open(wave) as file:
    lines = csv.reader(file)

    for line in lines:
        z.append(float(line[1]))        
        x.append(float(line[2]))
        y.append(float(line[3]))


plt.plot(x,y)
plt.grid()
plt.show()