package main

/*
#cgo LDFLAGS: -lz[LJ49]

#include <stdio.h>
#include <string.h>
#include <zlib.h>

void CExample() {
	char *text = "Hello, world!";
	unsigned long textSize = strlen(text);
	char buffer[1024] = {0, };
	unsigned long bufferSize = 1024;

	int err  = compress(buffer, &bufferSize, text, textSize); // zlib 함수로 압축
	printf("Compressed size: %lu\n", bufferSize);             // 압축한 크기 출력
	printf("zlib Version: %s\n", zlibVersion());              // zlib 버전 출력
}
*/
import "C"

func main() {
	C.CExample()
}
