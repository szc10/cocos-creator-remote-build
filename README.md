# cocos-creator-remote-build

## 项目介绍

当前项目为cocos creator 游戏项目快速制作热更新包的工具

使用golang 编写


## 使用说明

### 如何编译

```shell
go build version_generator.go
```

### 如何使用

1. 将编译好的文件`version_generator`,`config.json`,`version_generator.js`放置到cocos creator项目的根目录下
2. 根据实际情况修改config.json的内容 url:表示热更新的地址,asset表示本地jsb文件的位置
3. 点击version_generator,按对话输入相应指令,生成热更新文件 remote_assets.tar.gz

## 其他

 Platform | Support Status 
 -------- | ------------
 Mac | 10.10.0及以上 
 Window | 7及以上 