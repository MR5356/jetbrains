# 工具使用说明
## 作为命令行工具使用
### jetbrains ls
获取jetbrains产品最新版本的下载链接
```bash
jetbrains ls  # 获取所有平台的安装包链接
jetbrains ls windows  # 获取Windows系统安装包
jetbrains ls linux  # 获取linux系统安装包
jetbrains ls mac  # 获取mac系统安装包
jetbrains ls mac-m1  # 获取m1芯片的mac系统安装包
```

## 作为web API服务端使用
### jetbrains start
```bash
jetbrains start  # 在8999端口启动web服务
jetbrains start 2345  # 在自定义端口（2345）启动web服务
```
### API
接口地址：/{os}/links  # 其中os取值范围是"windows"、"linux"、"mac"、"mac-m1"

请求方式：GET

返回格式：JSON

返回示例：
```json
[
    {
        "name": "CLion", 
        "version": "2022.1.2", 
        "size": "601.72MB", 
        "build": "221.5787.29", 
        "date": "2022-06-01", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/cpp/CLion-2022.1.2.exe"
    }, 
    {
        "name": "DataGrip", 
        "version": "2022.1.5", 
        "size": "376.21MB", 
        "build": "221.5787.39", 
        "date": "2022-06-07", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/datagrip/datagrip-2022.1.5.exe"
    }, 
    {
        "name": "Goland", 
        "version": "2022.1.2", 
        "size": "430.07MB", 
        "build": "221.5787.30", 
        "date": "2022-06-02", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/go/goland-2022.1.2.exe"
    }, 
    {
        "name": "IntelliJ IDEA", 
        "version": "2022.1.2", 
        "size": "697.46MB", 
        "build": "221.5787.30", 
        "date": "2022-06-01", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/idea/ideaIU-2022.1.2.exe"
    }, 
    {
        "name": "PhpStorm", 
        "version": "2022.1.2", 
        "size": "416.24MB", 
        "build": "221.5787.33", 
        "date": "2022-06-02", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/webide/PhpStorm-2022.1.2.exe"
    }, 
    {
        "name": "PyCharm", 
        "version": "2022.1.2", 
        "size": "465.56MB", 
        "build": "221.5787.24", 
        "date": "2022-06-01", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/python/pycharm-professional-2022.1.2.exe"
    }, 
    {
        "name": "ReSharper Ultimate", 
        "version": "2022.1.2", 
        "size": "1.06GB", 
        "build": "2022.1.2.65536", 
        "date": "2022-06-03", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/resharper/dotUltimate.2022.1.2/JetBrains.dotUltimate.2022.1.2.exe"
    }, 
    {
        "name": "Rider", 
        "version": "2022.1.2", 
        "size": "711.65MB", 
        "build": "221.5787.36", 
        "date": "2022-06-03", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/rider/JetBrains.Rider-2022.1.2.exe"
    }, 
    {
        "name": "RubyMine", 
        "version": "2022.1.2", 
        "size": "396.25MB", 
        "build": "221.5787.34", 
        "date": "2022-06-03", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/ruby/RubyMine-2022.1.2.exe"
    }, 
    {
        "name": "WebStorm", 
        "version": "2022.1.2", 
        "size": "379.08MB", 
        "build": "221.5787.30", 
        "date": "2022-06-02", 
        "platFrom": "windows", 
        "link": "https://download.jetbrains.com/webstorm/WebStorm-2022.1.2.exe"
    }
]
```