# WebDU

WebDU is a disk utility (similar to the well-known [ncdu](https://dev.yorhel.nl/ncdu) program). 

WebDU written in Go with a web interface in VueJS. This program performs a quick analysis of the directory tree and displays the result of the occupied space with the ability to go to the child directory.

## Install

Download executable file from [releases](hhttps://github.com/SiTiSem/webdu/releases) page.

## Usage

    ./webdu
    
or

    ./webdu [param]
    
Available params:
* --port - connection port (by default --port=7000)
* --pass - authorization password (by default automatically generated)

After starting the program, the web-panel will be available at:
http://IP:PORT
