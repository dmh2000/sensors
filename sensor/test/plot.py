import sys
import csv
import matplotlib.pyplot as plt

wave = sys.argv[1]
fname = sys.argv[2]
y = []
t = []
with open(wave) as file:
    lines = csv.reader(file)

    for line in lines:
        t.append(float(line[0]))        
        y.append(float(line[1]))


plt.plot(t,y)
plt.grid()
plt.savefig(fname)