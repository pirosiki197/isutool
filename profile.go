package isutool

import (
	"os"
	"runtime/pprof"
	"time"

	"github.com/felixge/fgprof"
)

const rootDir = "/home/isucon/"

func Profile() {
	fgFile, err := os.Create(rootDir + "fgprof.out")
	if err != nil {
		panic(err)
	}
	stop := fgprof.Start(fgFile, fgprof.FormatPprof)

	pFile, err := os.Create(rootDir + "pprof.out")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(pFile)

	go func() {
		time.Sleep(60 * time.Second)
		stop()
		fgFile.Close()
		pprof.StopCPUProfile()
		pFile.Close()
	}()
}
