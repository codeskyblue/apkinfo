package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
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
	mainActivity, _ := pkg.MainActivity()
	if !strings.Contains(mainActivity, ".") {
		mainActivity = "." + mainActivity
	}

	shortMainActivity := mainActivity
	if strings.HasPrefix(mainActivity, pkgName) {
		shortMainActivity = mainActivity[len(pkgName):]
	}

	// log.Println(os.Args)
	fmt.Printf("## Package\n")
	fmt.Printf("PackageName:  %s\n", pkgName)
	fmt.Printf("MainActivity: %s\n", mainActivity)

	fmt.Print("\n## ADB\n")
	fmt.Printf("adb shell am start -a %s/%s\n", pkgName, shortMainActivity)

	fmt.Println("\n## AppCrawler")
	fmt.Printf("appcrawler --capability appPackage=%s,appActivity=%s\n", pkgName, shortMainActivity)

	fmt.Println("\n## Appium")
	data, _ := json.MarshalIndent(map[string]interface{}{
		"platformName":    "Android",
		"deviceName":      "whatever",
		"appPackage":      pkgName,
		"appActivity":     shortMainActivity,
		"unicodeKeyboard": true,
		"resetKeyboard":   true,
	}, "", "   ")
	fmt.Println(string(data))
	// fmt.Print("Press Enter to exit. ")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
}
