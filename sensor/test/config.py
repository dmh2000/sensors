import math
import sys

# frequency in Hz (cycles/sec): must be float d.d+
frequency = float(sys.argv[1])
# amplitude in m : must be float d.d+
amplitude = float(sys.argv[2])

# sample rate in samples/second : must be float d.d+
rate = 100
# delta t
dt = 1.0 / rate
y = 0.0
t = 0.0
pi2 = 2 * math.pi
period = 1.0 / frequency
for i in range(0, rate * 4):
    tx = math.fmod(t,pi2) # to avoid loss of precision when t is large
    y = amplitude * math.sin(tx * frequency * pi2)
    print(f"{t:.2f},{y:.2f},{dt:.2f}")
    t = t + dt
