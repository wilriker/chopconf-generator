# chopconf-generator
Small tool to create the bitmask required for a TMC2660 CHOPCONF register.

# Usage
Run `choconf -h` to see the available options

```
Usage of ./chopconf:
  -chm
        CHM: true|false (default false)
  -hdec uint
        HDEC: 16, 32, 48, 64 (default 16)
  -hend uint
        HEND: 0..15 (default 3)
  -hex
        Output value in hexadecimal notation
  -hstrt uint
        HSTRT: 0..7 (default 3)
  -rndtf
        RNDTF: true|false (default false)
  -tbl uint
        TBL: 16, 24, 36, 54 (default 36)
  -toff uint
        TOFF: 0..15 (default 4)
```

Default values are based on values used in RepRapFirmware at https://github.com/dc42/RepRapFirmware.

# Documentation
See https://www.trinamic.com/fileadmin/assets/Support/Appnotes/AN001-spreadCycle.pdf and https://forum.duet3d.com/topic/5392/does-m906-set-rms-or-peak-current/19 for more details.
