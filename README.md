# cloudgo-io
处理 web 程序的输入与输出

## 1、概述
设计一个 web 小应用，展示静态文件服务、js 请求支持、模板输出、表单处理、Filter 中间件设计等方面的能力。（不需要数据库支持）

## 2、任务要求
编程 web 应用程序 cloudgo-io。 
链接：

![](https://github.com/FlyingFeather/cloudgo-io/blob/master/screenshot/1.png)

### 基本要求
1. 支持静态文件服务
static为静态文件夹，有css，js，images文件夹

2. 支持简单 js 访问
这一项按照课堂文档中完成，通过http://localhost:8080/jsontest访问

![](https://github.com/FlyingFeather/cloudgo-io/blob/master/screenshot/4.png)

3. 提交表单，并输出一个表格
通过http://localhost:8080访问，提交后跳转到通过http://localhost:8080/index

![](https://github.com/FlyingFeather/cloudgo-io/blob/master/screenshot/2.png)

提交后跳转：
![](https://github.com/FlyingFeather/cloudgo-io/blob/master/screenshot/3.png)

4. 对 /unknown 给出开发中的提示，返回码 505
通过http://localhost:8080/unknown访问

![](https://github.com/FlyingFeather/cloudgo-io/blob/master/screenshot/5.png)


提示：

需要关注的 headers 与 格式

REQUEST POST：
Content-type: application/x-www-form-urlencoded;charset:gb2312

RESPONSE：
Content-Type: text/html;charset=gb2312
