package writers

import (
	"fmt"
	"github.com/Elot322/go_System_Monitor/functions"
)

func RenderCycleStart() {
	mainInfo()
	ramInfo()
	cpuInfo()
}

func mainInfo() {

	curUser, system := functions.GetMainInfo()

	fmt.Println("┌─────────────┐┌─────────────────────────────────────────────────────")
	fmt.Println("|OS:          ||", system, "                                         ")
	fmt.Println("|Uid user:    ||", curUser.Uid, "                                    ")
	fmt.Println("|Folder user: ||", curUser.HomeDir, "                                ")
	fmt.Println("|User name:   ||", curUser.Username, "                               ")
	fmt.Println("|Account name:||", curUser.Name, "                                   ")
	fmt.Println("└─────────────┘└─────────────────────────────────────────────────────")
}

func ramInfo() {
	ramData := functions.GetRamData()
	fmt.Printf("Memory total: %d bytes\n", ramData.Total)
	fmt.Printf("Memory used: %d bytes\n", ramData.Used)
	fmt.Printf("Memory free: %d bytes\n", ramData.Free)
}

func cpuInfo() {
	before, after, total := functions.GetCpuData()
	fmt.Printf("Cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("Cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("Cpu idle: %f %%\n", float64(after.Idle-before.Idle)/total*100)
}
