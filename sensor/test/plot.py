import sys
import csv
import matplotlib.pyplot as plt


fig,axs = plt.subplots(len(sys.argv)-1)
for i in range(1, len(sys.argv)):
    y = []
    t = []
    with open(sys.argv[i]) as file:
        lines = csv.reader(file)
        for line in lines:
            t.append(float(line[0]))        
            y.append(float(line[1]))
    axs[i-1].plot(t,y)
    plt.grid()

plt.show()