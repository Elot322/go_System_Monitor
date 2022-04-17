package functions

import (
	"fmt"
	"github.com/mackerelio/go-osstat/memory"
	"os/user"
)

func GetMainInfo() {
	//Get data current user
	curUser, err := user.Current()
	if err != nil {
		fmt.Println("Current user: NO DATA!")
	}
	fmt.Println("┌─────────────┐┌─────────────────────────────────────────────────────┐")
	fmt.Println("|Uid user:    ||", curUser.Uid, "      |")
	fmt.Println("|Folder user: ||", curUser.HomeDir, "                                     |")
	fmt.Println("|User name:   ||", curUser.Username, "                              |")
	fmt.Println("|Account name:||", curUser.Name, "                                     |")
	fmt.Println("└─────────────┘└─────────────────────────────────────────────────────┘")
}

func GetCpuData() {

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
