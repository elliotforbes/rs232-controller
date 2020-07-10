# Install Instructions

1. Open Terminal

```bash
$ cd Documents
$ cd rs232-controller
$ git pull origin master
$ sudo ./rs232-controller /dev/ttyUSB0 BIT_RATE "HEXCODEHALFBRIGHTNESS" "HEXCODEFULLBRIGHTNESS" &
```

## Running The Program

```bash
sudo ./rs232-controller /dev/ttyUSB0 BIT_RATE HALFBRIGHTNESS FULLBRIGHTNESS &
```

## Testing the Program

```bash
sudo ./test_rs232 /dev/ttyUSB0 BIT_RATE "HEXCODETURNON"
```
