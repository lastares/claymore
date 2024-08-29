# Claymore    :tada::tada::tada: :tada::tada::tada:

> 本包开发使用的go版本 `go 1.22.6`

go开发中经常使用一些工具函数，每次新项目或者到了一个新坑位都要重新去写，很是麻烦
所以，这个项目就是封装一些常用的工具函数，方便 Gopher 开发，希望能成为 Gopher 开发中经常使用的 **claymore**。

**该包目前已支持的工具函数**

### 文件(fileutil) ###

| 编号  | 函数               | 功能                     |   
|-----|------------------|------------------------|
| 001 | Download         | 下载文件到本地                |
| 002 | GetExtension     | 获取文件后缀                 |
| 003 | GetFullName      | 获取文件名称                 |
| 004 | GetBaseName      | 获取文件名称(不带后缀名)          |
| 005 | GetExtNoDot      | 获取文件后缀(不带点)            |
| 006 | GetDirFileList   | 获取某个目录下的所有文件           |
| 007 | GetDirFileListV2 | 获取某个目录下的所有文件(不包含多级子目录) |
| 007 | DirOrFileExists  | 判断本地文件或目录是否存在          |                   |

### 字符串(stringutil) ###

| 编号  | 函数        | 功能            |   
|-----|-----------|---------------|
| 001 | Substr    | 截取字符串的子串      |
| 002 | Md5       | 生成 md5 hash 值 |
| 003 | StrToByte | 字符串转byte      |
| 004 | ByteToStr | byte转字符串      |

### 解压缩(ziputil) ###

| 编号  | 函数    | 功能              |   
|-----|-------|-----------------|
| 001 | Zip   | 压缩某个目录下的文件为zip包 |
| 002 | Unzip | 解压zip文件         |

### 切片(sliceutil) ###

| 编号  | 函数         | 功能    |   
|-----|------------|-------|
| 001 | MakeSorter | 泛型排序器 |

### Json(jsonutil) ###

| 编号  | 函数           | 功能        |
|-----|--------------|-----------|
| 001 | JsonEncode() | json 序列化  |
| 002 | JsonDecode() | json 反序列化 |

### Gorm(dbutil) ###

| 编号  | 函数           | 功能             |
|-----|--------------|----------------|
| 001 | New()        | 连接数据库，获取Gorm实例 |

### 其他(generalutil) ###

| 编号  | 函数                  | 功能           |
|-----|---------------------|--------------|
| 001 | NewPaginator()      | 基于泛型的通用分页构造器 |
| 001 | PrettyPrintStruct() | 优雅的打印结构体     |

:shipit: :shipit: :shipit: 其他函数持续增加中... :heart: :heart: :heart:

这里列了几个经常使用的其他工具包函数，有兴趣的可以看下

- [类型转换](https://github.com/spf13/cast)
- [泛型函数工具包](https://github.com/samber/lo)
- [强大工具函数库](https://github.com/duke-git/lancet)

本包是对其他包缺少函数的一些补充
