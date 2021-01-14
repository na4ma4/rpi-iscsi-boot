# rpi-iscsi-boot

Booting a Raspberry Pi over iSCSI using a sdcard stub.

## Building kernels

Based on [this guide](https://troy.dack.com.au/raspberry-pi-iscsi-rootfs/)

```shell
docker build -t raspi-build:latest .
docker run -ti --rm -v "$(pwd):/data" --workdir /data raspi-build:latest
```

```shell
mkdir /data/{ext4,fat32{,/overlays}}
git clone --depth=1 https://github.com/raspberrypi/linux
cd linux
KERNEL=kernel7l
make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- bcm2711_defconfig
make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- menuconfig
```

change:

- `CONFIG_ISCSI_BOOT_SYSFS=y`
- `CONFIG_ISCSI_TCP=y`
- `CONFIG_LOCALVERSION="-iscsi"`

```shell
time make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- zImage modules dtbs -j5
time env PATH=$PATH make ARCH=arm CROSS_COMPILE=arm-linux-gnueabihf- INSTALL_MOD_PATH=/data/ext4 modules_install
cp arch/arm/boot/zImage /data/fat32/$KERNEL.img
cp arch/arm/boot/dts/*.dtb /data/fat32/
cp arch/arm/boot/dts/overlays/*.dtb* /data/fat32/overlays/
cp arch/arm/boot/dts/overlays/README /data/fat32/overlays/
```

## Initramfs

`/boot/cmdline.txt`

```text
console=serial0,115200 console=tty1 root=/dev/sda2 iscsi_auto elevator=deadline fsck.repair=yes rootwait quiet splash plymouth.ignore-serial-consoles
```

- Add [device-iscsi](device-iscsi) to `/etc/initramfs-tools/scripts/local-top/`.
- Add `device-iscsi` to `PREREQ` in `/usr/share/initramfs-tools/scripts/local-top/iscsi`.
- Add `iscsi` to `/etc/initramfs-tools/modules`.
- `sudo update-initramfs -v -c -k $(uname -r)` (may need to replace `uname -r` with specific version including localver from build step).

## References

- [How To Connect And Mount iSCSI Onto Linux Servers](https://www.unixmen.com/attach-iscsi-target-disks-linux-servers/)
- [Raspberry Pi - iSCSI RootFS](https://troy.dack.com.au/raspberry-pi-iscsi-rootfs/)
- [Raspberry Pi - Kernel building](https://www.raspberrypi.org/documentation/linux/kernel/building.md)
