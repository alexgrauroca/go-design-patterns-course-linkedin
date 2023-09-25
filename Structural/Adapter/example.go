package main

import "fmt"

func main() {
	tv1 := &sammysangAdapter{
		sstv: &SammysangTV{
			currentChan:   13,
			currentVolume: 35,
			tvOn:          true,
		},
	}
	tv2 := &SohneeTV{
		vol:     20,
		channel: 9,
		isOn:    true,
	}

	tv2.turnOn()
	tv2.volumeUp()
	tv2.volumeDown()
	tv2.channelUp()
	tv2.channelDown()
	tv2.goToChannel(68)
	tv2.turnOff()

	fmt.Println("-------------------------")

	tv1.turnOn()
	tv1.volumeUp()
	tv1.volumeDown()
	tv1.channelUp()
	tv1.channelDown()
	tv1.goToChannel(82)
	tv1.turnOff()
}
