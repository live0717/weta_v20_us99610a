#!/bin/bash

DEVICE=us996
BUILD=./build.sh
CLEAN=make mrproper
OUT=~/home/mentalmuso/kernel/out/us996
IMG=~/home/mentalmuso/kernel/weta_v20_us99610a/build/arch/arm64/boot/Image.lz4-dtb
IMG2=~/home/mentalmuso/superr/boot.img-zImage



### clean good and build

$CLEAN
$BUILD $DEVICE


### build finished, copy img to out

mv $IMG $IMG2
