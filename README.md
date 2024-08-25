# Claymore    :tada::tada::tada: :tada::tada::tada:

> 当前包使用的go版本 `go 1.22.6`

go开发中经常使用一些工具函数，每次新项目或者到了一个新坑位都要重新去写，很是麻烦
所以，这个项目就是封装一些常用的工具函数，方便 Gopher 开发，希望能成为 Gopher 开发中经常使用的 **claymore**。

**该包目前已支持的工具函数**

### 文件(fileutil) ###

| 编号  | 函数             | 功能            |   
|-----|----------------|---------------|
| 001 | Download       | 下载文件到本地       |
| 002 | GetExtension   | 获取文件后缀        |
| 003 | GetFullName    | 获取文件名称        |
| 004 | GetBaseName    | 获取文件名称(不带后缀名) |
| 005 | GetExtNoDot    | 获取文件后缀(不带点)   |
| 006 | GetDirFileList | 获取某个目录下的所有文件  |

### 字符串(stringutil) ###
| 编号  | 函数            | 功能            |   
|-----|---------------|---------------|
| 001 | Substr        | 截取字符串的子串      |

### 解压缩(ziputil) ###
| 编号  | 函数    | 功能              |   
|-----|-------|-----------------|
| 001 | Zip   | 压缩某个目录下的文件为zip包 |
| 002 | Unzip | 解压zip文件         |

:shipit: :shipit: :shipit: 其他函数持续增加中... :heart: :heart: :heart:

为了不与一些其他的go工具包中的函数重复，这里列了几个经常使用的其他工具包函数，有兴趣的可以看下
- [类型转换](https://github.com/spf13/cast) 
- [泛型函数工具包](https://github.com/samber/lo) 
- [强大工具函数库](https://github.com/duke-git/lancet)
