package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
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
--json          Show as json output
--icon <path>   Save icon as file, only jpg support
`, version)
)

func init() {
	flag.Usage = func() {
		fmt.Print(help)
	}
}

func printDefault(pkgName, mainActivity, label string) {
	if !strings.Contains(mainActivity, ".") {
		mainActivity = "." + mainActivity
	}

	shortMainActivity := mainActivity
	if strings.HasPrefix(mainActivity, pkgName) {
		shortMainActivity = mainActivity[len(pkgName):]
	}

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
}

func saveAsJpeg(pkg *apk.Apk, path string) error {
	img, err := pkg.Icon(nil)
	if err != nil {
		return errors.Wrap(err, "get icon")
	}
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "create file")
	}
	defer f.Close()
	// log.Println(img)
	return jpeg.Encode(f, img, &jpeg.Options{Quality: 80})
}

func main() {
	iconPath := flag.String("icon", "", "save apk icon")
	bjson := flag.Bool("json", false, "show as json format")
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
	label, _ := pkg.Label(nil)
	if *iconPath != "" {
		if err := saveAsJpeg(pkg, *iconPath); err != nil {
			log.Fatal(err)
		}
		return
	}

	if *bjson {
		data, _ := json.MarshalIndent(map[string]interface{}{
			"label":        label,
			"packageName":  pkgName,
			"mainActivity": mainActivity,
			"versionName":  pkg.Manifest().VersionName,
			"versionCode":  pkg.Manifest().VersionCode,
		}, "", "  ")
		fmt.Println(string(data))
	} else {
		printDefault(pkgName, mainActivity, label)
	}

	if filepath.IsAbs(filename) {
		fmt.Print("Press Enter to exit. ")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
