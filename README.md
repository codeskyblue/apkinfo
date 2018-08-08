# apkinfo
get apkinfo use pure go code. (packageName, mainActivty etc...)

## Install
Download binaries from [releases](https://github.com/codeskyblue/apkinfo/releases)

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

获取格式化的json信息

```bash
$ apkinfo -json cloudmusic.apk
{
  "label": "网易云音乐",
  "mainActivity": "com.netease.cloudmusic.activity.LoadingActivity",
  "packageName": "com.netease.cloudmusic",
  "versionCode": 118,
  "versionName": "5.2.0"
}
```

获取icon

```bash
$ apkinfo -icon icon.jpg
```

## Not supported
Can not extract icon which is using [Adaptive icon](https://developer.android.com/guide/practices/ui_guidelines/icon_design_adaptive)

## References
- [Appium Capabilities](https://github.com/appium/appium/blob/master/docs/en/writing-running-appium/caps.md)
- [AppCrawler](https://github.com/seveniruby/AppCrawler)
- [Awesome-ADB](https://github.com/mzlogin/awesome-adb)

# LICENSE
[MIT](LICENSE)
