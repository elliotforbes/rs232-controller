# Install Instructions

1. Open Terminal

```bash
$ cd Documents
$ cd rs232-controller
$ git pull origin master
$ ./rs232-controller /dev/ttyUSB0 "HEXCODEHALFBRIGHTNESS" "HEXCODEFULLBRIGHTNESS" &
```

## Running The Program

```bash
sudo ./rs232-controller /dev/ttyUSB0 HALFBRIGHTNESS FULLBRIGHTNESS &
```

## Testing the Program

```bash
sudo ./test_rs232 /dev/ttyUSB0 HEXSTRING
```
