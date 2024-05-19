import sys
import csv
import matplotlib.pyplot as plt

wave = sys.argv[1]
t = []
y = []
with open(wave) as file:
    lines = csv.reader(file)

    for line in lines:
        t.append(float(line[0]))        
        y.append(float(line[1]))


plt.plot(t,y)
plt.grid()
plt.show()