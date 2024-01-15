#! micropython
# color_setup.py - see also setup_examples/*.py
# Must export SSD class and ssd instance.

from machine import Pin, PWM, SPI
import gc

from drivers.st7789.st7789_4bit import ST7789 as SSD

# Backlight control
bl_pin = Pin(6, Pin.OUT, value=1)
bl_pwm = PWM(bl_pin)
bl_pwm.freq(1_000)
bl_pwm.duty_u16(45_000)

dc_pin = Pin(7, Pin.OUT, value=0)
rst_pin = Pin(14, Pin.OUT, value=1)  # n/c (SDCS)
cs_pin = Pin(9, Pin.OUT, value=1)

gc.collect()

spi = SPI(1,
          sck=Pin(10),
          mosi=Pin(11),
          miso=Pin(8),
          baudrate=60_000_000)

ssd = SSD(spi, dc=dc_pin, cs=cs_pin, rst=rst_pin,
          width=240, height=320)

#--#
