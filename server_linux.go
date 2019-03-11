// UDP сервер для Linux
package main

import (
	"fmt"
	"syscall"
)

func main() {
	// создаем сокет udp и переводим его в неблокирующий режим
	handle, err := syscall.Socket(syscall.AF_INET, syscall.O_NONBLOCK|syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		panic(err)
	}

	// задаем опцию широковещательной рассылки из сокета
	if err := syscall.SetsockoptInt(handle, syscall.SOL_SOCKET, syscall.SO_BROADCAST, 1); err != nil {
		panic(err)
	}

	// биндим порт к сокету
	if err := syscall.Bind(handle, &syscall.SockaddrInet4{Port: 9000, Addr: [4]byte{127, 0, 0, 1}}); err != nil {
		panic(err)
	}

	buffer := make([]byte, 256)
	for {
		n, addr, err := syscall.Recvfrom(handle, buffer, 1)
		if err != nil {
			continue
		}
		fmt.Println(n, addr, string(buffer))
	}

	if err = syscall.Close(handle); err != nil {
		panic(err)
	}
}
