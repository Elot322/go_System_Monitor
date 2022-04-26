package functions

import (
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
	"os/user"
	"runtime"
	"time"
)

func GetMainInfo() (*user.User, string) {
	//Get data current user and system
	system := runtime.GOOS
	curUser, err := user.Current()
	if err != nil {
		GetError("ERROR USER DATA!")
	}
	return curUser, system
}

func GetCpuData() (*cpu.Stats, *cpu.Stats, float64) {

	before, err := cpu.Get()
	if err != nil {
		GetError("CPU DATA NOT WORK FOR THIS SYSTEM!")
	}

	time.Sleep(time.Duration(1) * time.Second)

	after, _ := cpu.Get()
	total := float64(after.Total - before.Total)

	return before, after, total
}

func GetRamData() *memory.Stats {
	ramData, err := memory.Get()
	if err != nil {
		GetError("RAM DATA ERROR")
	}

	return ramData
}

func GetMemoryData() {

}

func GetInternetData() {

}
