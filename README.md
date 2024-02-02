# micropython-pico-presentations
Presentations on micropython and raspberrpy pico

## Part 1 - MicroPython and Pi Pico

**Evaluating Features and Performance**

**(Can it replace Arduino?)**

Python on a $4 microcontroller sounds promising - but how good is it?
Make a list of features to test, attempt some quick demos starting from the
offical examples, measure performance, investigate where slow.
From blinking an LED with GPIO to an interrupt driven quadrature decoder,
this presentation covers the process of evaluating MicroPython on the pi pico
focusing on performance and developer effort.


## Part 2 - PIO Peripheral of the Raspberry Pi Pico

**Offloading GPIO Bit-Banging to the PIO**

When your processor does not implement (or have enough of) a peripheral,
sometimes you can make your own by toggling GPIO pins with the correct timing,
commonly referred to as "bit-banging".
The Raspberry Pi Pico includes a simple PIO (Programmable Input Output)
processor dedicated to such tasks (and at very high speed).
This presentation starts with the architecture and programming model of the PIO,
then switches to MicroPython for creating PIO programs, controlling a PIO,
and performing input and output with a PIO.

Note that while Part 1 includes demonstrating a couple of the official
PIO examples, all details of how it works are left to Part 2.



## View online

* [part 1](https://htmlpreview.github.io/?https://github.com/nludban/micropython-pico-presentations/blob/master/presentation-1.html).
* [part 2](https://htmlpreview.github.io/?https://github.com/nludban/micropython-pico-presentations/blob/master/presentation-2.html)
