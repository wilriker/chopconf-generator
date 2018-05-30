package main

import (
	"flag"
	"fmt"
)

func toff(n uint64) uint64 {
	if n > 15 {
		panic("toff must be <= 15")
	}
	return (uint64(n) & 15) << 0
}

func hstrt(n uint64) uint64 {
	if n > 7 {
		panic("hstrt <= 7")
	}
	return (uint64(n) & 7) << 4
}

func hend(n uint64) uint64 {
	if n > 15 {
		panic("hend <=15")
	}
	return (uint64(n) & 15) << 7
}

func hdec(n uint64) uint64 {
	ret := 0
	switch n {
	case 16:
		ret = 0
	case 32:
		ret = 1
	case 48:
		ret = 2
	case 64:
		ret = 3
	default:
		panic("hdec must be 16, 32, 48 or 64")
	}
	return (uint64(ret) & 3) << 11
}

func rndtf(enable bool) uint64 {
	if !enable {
		return 0
	}
	return 1 << 13
}

func chm(enable bool) uint64 {
	if !enable {
		return 0
	}
	return 1 << 14
}

func tbl(n uint64) uint64 {
	ret := 0
	switch n {
	case 16:
		ret = 0
	case 24:
		ret = 1
	case 36:
		ret = 2
	case 54:
		ret = 3
	default:
		panic("tbl must be 16, 24, 36 or 54")
	}
	return (uint64(ret) & 3) << 15
}

func main() {
	var toffVal, hstrtVal, hendVal, hdecVal, tblVal uint64
	var rndtfVal, chmVal bool

	flag.Uint64Var(&toffVal, "toff", 4, "TOFF: 0..15")
	flag.Uint64Var(&hstrtVal, "hstrt", 3, "HSTRT: 0..7")
	flag.Uint64Var(&hendVal, "hend", 3, "HEND: 0..15")
	flag.Uint64Var(&hdecVal, "hdec", 16, "HDEC: 16, 32, 48, 64")
	flag.BoolVar(&rndtfVal, "rndtf", false, "RNDTF: true|false")
	flag.BoolVar(&chmVal, "chm", false, "CHM: true|false")
	flag.Uint64Var(&tblVal, "tbl", 36, "TBL: 16, 24, 36, 54")
	flag.Parse()

	if hendVal+hstrtVal > 15 {
		panic("hstrt + hend must be <= 15")
	}

	chopregister := uint32(tbl(tblVal) | chm(chmVal) | rndtf(rndtfVal) | hdec(hdecVal) | hend(hendVal) | hstrt(hstrtVal) | toff(toffVal))

	fmt.Println(chopregister)
}
