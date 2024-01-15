#!/bin/bash

echo hello

AMPY=ampy
PORT=/dev/cuaU0

SUBDIRS=(
    "drivers"
    "drivers/7789"
    "gui"
    "gui/core"
    "gui/fonts"
    "gui/widgets"
    "extras"
    "extras/widgets"
)

SOURCES=(
    "gui/core/__init__.py"
    "gui/core/colors.py"
    "gui/core/fplot.py"
    "gui/core/nanogui.py"
    "gui/core/writer.py"
    "gui/fonts/arial10.py"
    "gui/fonts/arial35.py"
    "gui/fonts/arial_50.py"
    "gui/fonts/courier20.py"
    "gui/fonts/font10.py"
    "gui/fonts/font6.py"
    "gui/fonts/freesans20.py"
    "gui/widgets/__init__.py"
    "gui/widgets/dial.py"
    "gui/widgets/label.py"
    "gui/widgets/led.py"
    "gui/widgets/meter.py"
    "gui/widgets/scale.py"
    "gui/widgets/textbox.py"
    "extras/parse2d.py"
    "extras/widgets/calendar.py"
    "extras/widgets/clock.py"
    "extras/widgets/eclock.py"
    "extras/widgets/grid.py"
)

for subdir in ${SUBDIRS[@]}; do
    cmd="${AMPY} -p ${PORT} mkdir --exists-okay ${subdir}"
    echo $cmd
    $cmd
done

for source in ${SOURCES[@]}; do
    cmd="${AMPY} -p ${PORT} put ${source} /${source}"
    echo $cmd
    $cmd
done

#--#
