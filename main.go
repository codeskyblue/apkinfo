package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shogo82148/androidbinary/apk"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func init() {
	flag.Usage = func() {
		fmt.Printf(`Apkinfo is a tool to see apk info.

Usage:

    apkinfo <apk-file-path>

About:

    version %s, commit %s date %s

    If there is any problem, please raise a issue in
    https://github.com/codeskyblue/apkinfo
`, version, commit, date)
	}
}

func main() {

	flag.Parse()
	filename := flag.Arg(0)
	pkg, err := apk.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	pkgName := pkg.PackageName()
	mainActivity, _ := pkg.MainAcitivty()
	// log.Println(os.Args)
	fmt.Printf("Package >>\n")
	fmt.Printf("\tPackageName:  %s\n", pkgName)
	fmt.Printf("\tMainActivity: %s\n", mainActivity)
	fmt.Print("Shell >>\n")
	fmt.Printf("\tadb shell am start -a %s/%s\n", pkgName, mainActivity)
	fmt.Println("AppCrawler >>")
	shortMainActivity := mainActivity
	if strings.HasPrefix(mainActivity, pkgName) {
		shortMainActivity = mainActivity[len(pkgName):]
	}
	fmt.Printf("\t--capability appPackage=%s,appActivity=%s\n", pkgName, shortMainActivity)

	fmt.Print("Press Enter to exit. ")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
