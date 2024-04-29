import sys
import os
import time
from queue import Queue
import datetime as dt
import matplotlib.pyplot as plt
import matplotlib.animation as animation
import threading
import pika


def rabbitmq(q):
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='localhost'))
    channel = connection.channel()

    channel.queue_declare(queue='waveforms')

    def callback(ch, method, properties, body):
        q.put(body.decode("utf-8"))

    channel.basic_consume(queue='waveforms', on_message_callback=callback, auto_ack=True)

    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()


# This function is called periodically from FuncAnimation
def create_animate(q):
    def animate(i, xs, ys):

        # Read temperature (Celsius) from TMP102
        s = q.get()
        v = s.split(",")
        temp_c = float(v[2])

        # Add x and y to lists
        xs.append(dt.datetime.now().strftime('%H:%M:%S.%f'))
        ys.append(temp_c)

        # Limit x and y lists to 20 items
        xs = xs[-20:]
        ys = ys[-20:]

        # Draw x and y lists
        ax.clear()
        ax.plot(xs, ys)

        # Format plot
        plt.xticks(rotation=45, ha='right')
        plt.subplots_adjust(bottom=0.30)
        plt.title('val')
        plt.ylabel('Temperature (deg C)')
    return animate


if __name__ == '__main__':
    q = Queue()
    t = threading.Thread(target=rabbitmq, args=(q,))
    t.start()

    # Create figure for plotting
    fig, ax = plt.subplots()
    xs = []
    ys = []

        # Set up plot to call animate() function periodically
    a = create_animate(q)
    ani = animation.FuncAnimation(fig, a, fargs=(xs, ys), interval=100)
    plt.show()  
