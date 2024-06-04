#!/bin/bash

set -e

tinygo flash -target=nucleo-f103rb main.go
