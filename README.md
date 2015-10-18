Simple Lottery
==============

Simple command line lottery.

Generates random number among given minimum and maximum number.
Max number is required. Min number is by default 1.


Usage
-----

The command line tool support 2 lottery mode.

You can use it to generate random number from `min` to `max`:

    lottery -max <NUMBER> [-min <NUMBER>]

or you may read random row from a provided excel file (.xlsx).
In which case, the random number generated will be specifying the
row number in the file. The first 2 cell in the row will be displayed:

    lottery -file <.XLSX FILENAME> [-max <NUMBER>] [-min <NUMBER>]


Installation
------------
With `$GOPATH` setup correctly and `$GOPATH/bin` as part of
your `$PATH`, you may just install with:

    go get github.com/yookoala/lottery

Or you may clone the source then build it manually:

    git clone https://github.com/yookoala/lottery.git
    cd lottery
    go build


Licence
-------

This software is licenced with MIT Licence.

You may find a copy of the LICENCE in this repository.
