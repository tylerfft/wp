package mylib

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
)

func CheckError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {
			fmt.Println("用户退出了")
			return 0
		}
		log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}

func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}
