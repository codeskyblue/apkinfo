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

	help = fmt.Sprintf(`apkinfo %s - (C) github.com/codeskyblue/apkinfo
Released under the MIT License

Usage: apkinfo [OPTIONS] apk-file

OPTIONS:
-h --help		Print this help screen
`, version)
)

func init() {
	flag.Usage = func() {
		fmt.Print(help)
	}
}

func main() {
	flag.String("icon", "", "save apk icon")
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

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
	label, _ := pkg.Label(nil)
	fmt.Printf("## Package\n")
	fmt.Printf("LabelName: %s\n", label)
	fmt.Printf("PackageName:  %s\n", pkgName)
	fmt.Printf("MainActivity: %s\n", mainActivity)

	fmt.Print("\n## ADB\n")
	fmt.Printf("adb shell am force-stop %s\n", pkgName)
	fmt.Printf("adb shell pm clear %s\n", pkgName)
	fmt.Printf("adb shell am start -n %s/%s\n", pkgName, shortMainActivity)

	fmt.Println("\n## AppCrawler")
	fmt.Printf("appcrawler --capability appPackage=%s,appActivity=%s\n", pkgName, shortMainActivity)

	fmt.Println("\n## Appium")
	data, _ := json.MarshalIndent(map[string]interface{}{
		"platformName":       "Android",
		"deviceName":         "whatever",
		"appPackage":         pkgName,
		"appActivity":        shortMainActivity,
		"automationName":     "UiAutomator2",
		"newCommandTimeout":  300,
		"noReset":            false,
		"dontStopAppOnReset": false,
		"unicodeKeyboard":    true,
		"resetKeyboard":      true,
	}, "", "   ")
	fmt.Println(string(data))
	// fmt.Print("Press Enter to exit. ")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')
}
