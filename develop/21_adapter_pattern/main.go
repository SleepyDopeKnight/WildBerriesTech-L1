package main

/*
Реализовать паттерн «адаптер» на любом примере.
*/
import "fmt"

type ChargePort interface {
	adapterForMicroUSB() string
}

type MicroUSB struct{} // наш объект - micro-usb

func (*MicroUSB) connectMicroUSBPort() string { // имитируем что подключились к micro-usb
	return "connected to micro-usb port"
}

func newAdapter(mu *MicroUSB) ChargePort { // конструктор для нашего адаптера
	return &TypeC{micro: mu}
}

type TypeC struct { // наш объект - type-c
	micro *MicroUSB // создаем поле в объекте, указывающее на объект micro-usb
}

func (tc *TypeC) adapterForMicroUSB() string { // создаем сам адаптер для type-c на micro-usb
	return tc.micro.connectMicroUSBPort() // иммитруем адаптацию type-c(tc) на micro-usb(micro)
}

func main() {
	typeC := newAdapter(&MicroUSB{}) // передаем в адаптер micro-usb и получаем объект type-c, реализующий интерфейс ChargePort
	fmt.Println(typeC.adapterForMicroUSB())
}
