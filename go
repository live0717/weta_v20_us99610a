#!/bin/bash
############################################
### mentalmuso starter and output script ###
############################################

VERS=1.23
FIR=us99610a

BUILD=~/kernel/weta_v20_us99610a/build/arch/arm64/boot/Image.lz4-dtb
IMG=~/kernel/weta_v20_us99610a/imgout/boot
IMGOUT=~/kernel/weta_v20_us99610a/imgout
ZIP=~/kernel/weta_v20_us99610a/imgout/zip
ZIPOUT=~/kernel/weta_v20_us99610a/zipout
WETAIMG=~/kernel/weta_v20_us99610a/imgout/zip/kernel/weta_boot.img

export VER=$VERS
export FIRM=$FIR

############################################
### ---> clean up and build

make mrproper
./build.sh us996

############################################
### ---> make flashable zip 

mv $BUILD $IMG/boot.img-kernel
cd $IMG
./mkbootimg --kernel boot.img-kernel --ramdisk boot.img-ramdisk.cpio.gz -o $WETAIMG
cd $ZIP
rm -f $ZIPOUT/*
zip -r $ZIPOUT/WETA_Kernel_$FIR_$VERS.zip *




