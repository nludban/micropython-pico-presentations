# micropython-pico-presentations
Presentations on micropython and raspberrpy pico

## Part 1 - MicroPython and Pi Pico

**Replacing the Arduino**

A Python sized for a microcontroller, running on a processor sized to run it...
is this the silver bullet that could end development with the Arduino?
This presentation is an overview of the features of micropython and the
Pi Pico, focusing on development time and performance.
From blinking an LED to interrupt driven quadrature decoding,
can it keep up with the demands of a real-time embedde system?


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
