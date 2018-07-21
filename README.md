# apkinfo
get apkinfo use pure go code. (packageName, mainActivty etc...)

## Install
Download binaries from [releases](/releases)

Or build from source

```bash
$ go get -v github.com/codeskyblue/apkinfo
```

## Usage
This tool is very simple to use.

```
$ apkinfo cloudmusic.apk
## Package
PackageName:  com.netease.cloudmusic
MainActivity: com.netease.cloudmusic.activity.LoadingActivity

## ADB
adb shell am start -a com.netease.cloudmusic/.activity.LoadingActivity

## AppCrawler
appcrawler --capability appPackage=com.netease.cloudmusic,appActivity=.activity.LoadingActivity

## Appium
{
   "appActivity": ".activity.LoadingActivity",
   "appPackage": "com.netease.cloudmusic",
   "deviceName": "whatever",
   "platformName": "Android",
   "resetKeyboard": true,
   "unicodeKeyboard": true
}
```

# LICENSE
[MIT](LICENSE)
