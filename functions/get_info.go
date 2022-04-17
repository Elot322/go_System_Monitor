package functions

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"os/user"
	"runtime"
	"time"
)

func GetMainInfo() {
	//Get data current user and system
	system := runtime.GOOS
	curUser, err := user.Current()
	if err != nil {
		fmt.Println("Current user: NO DATA!")
	}
	fmt.Println("┌─────────────┐┌─────────────────────────────────────────────────────┐")
	fmt.Println("|OS:          ||", system, "                                            |")
	fmt.Println("|Uid user:    ||", curUser.Uid, "      |")
	fmt.Println("|Folder user: ||", curUser.HomeDir, "                                     |")
	fmt.Println("|User name:   ||", curUser.Username, "                              |")
	fmt.Println("|Account name:||", curUser.Name, "                                     |")
	fmt.Println("└─────────────┘└─────────────────────────────────────────────────────┘")
}

func GetCpuData() {
	fmt.Println(" ")
	before, err := cpu.Get()
	if err != nil {
		fmt.Println("CPU DATA NOT WORK FOR THIS SYSTEM!")
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Println("CPU DATA NOT WORK FOR THIS SYSTEM!")
		return
	}
	total := float64(after.Total - before.Total)
	fmt.Printf("Cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("Cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("Cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}

func GetRamData() {
	ramData, err := memory.Get()
	if err != nil {
		fmt.Println("RAM DATA ERROR!")
	}

	fmt.Printf("Memory total: %d bytes\n", ramData.Total)
	fmt.Printf("Memory used: %d bytes\n", ramData.Used)
	fmt.Printf("Memory free: %d bytes\n", ramData.Free)
}

func GetMemoryData() {

}

func GetInternetData() {

}
