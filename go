#!/bin/bash

IMG1=~/kernel/weta_v20_us99610a/build/arch/arm64/boot/Image.lz4-dtb
IMG2=~/kernel/superr/super_us996/boot.img-zImage

make mrproper
./build.sh us996

sudo mv $IMG1 $IMG2
