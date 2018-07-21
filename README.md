# apkinfo
get apkinfo use pure go code. (packageName, mainActivty etc...)

## Usage
```bash
$ go get -v github.com/codeskyblue/apkinfo
$ ./apkinfo example.apk
Package >>
	PackageName:  com.netease.cloudmusic
	MainActivity: com.netease.cloudmusic.activity.LoadingActivity
Shell >>
	adb shell am start -a com.netease.cloudmusic/com.netease.cloudmusic.activity.LoadingActivity
AppCrawler >>
	--capability appPackage=com.netease.cloudmusic,appActivity=.activity.LoadingActivity
Press Enter to exit.
```

# LICENSE
[MIT](LICENSE)
