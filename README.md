# <img src="assets/logo.png" width="30px" height="30px" />Alpha智能图像

## _Image Detection And Image Recognition Tools_

**首页展示**  
![home](assets/home.png)  

### 项目介绍

alpha智能图像项目是全栈项目，主要对图片的识别，探测，相似度对比，留言交流，以及视频检测的工具平台。前端vue.js+vueX,后端golang+gin,以及模型训练python脚本。  

项目分3个部分:

- alpha-ui  
 前端部分，提供图像探测，图像检测，视频检测，模型训练以及开放API,提供图片的上传。  
- alpha-core  
 使用gin作为http服务框架，提供resetful服务。
 启动成功后展示所有相关API：  
 ![alpha-core-log](assets/alpha-core-log.png)  
- alpha-script  
  以python作为数据执行脚本。通过训练的模型数据，检测图片。参考[ImageAI](https://imageai-cn.readthedocs.io/zh_CN/latest/)  

### 安装部署

- 安装mongodb  
 安装mongodb并配置在alpha-core/src/config中

 ```golang
 const DB_URL = "mongodb://localhost:27017"
 ```

- 搭建文件服务器  
 1 这里搭建nginx作为文件服务器,有条件的话最好使用CDN服务器。搭建nginx自行网上查阅.配置完后将上传目录配置在alpha-core的config下

 ```golang
 const FILE_DOWNLOAD_PATH = "/Users/gaojiayi/alpha-static/download/"
const FILE_UPLOAD_PATH = "/Users/gaojiayi/alpha-static/upload/"

 ```
 
- alpha-script安装  
  将alpha-script拷贝到任意路径，同时在ImageAI的官网上下载训练的模型数据放在models目录下，具体的执行入口在scripts目录下。接下来将alpha-script的安装目录配置早alpha-core的config下

 ```golang
const Python_home = "/Users/gaojiayi/PycharmProjects/alpha-script/venv/bin/python"
const Alpha_script = "/Users/gaojiayi/PycharmProjects/alpha-script/"
const Detect_script_path = Alpha_script +"scripts/image_detect_executor.py"
const Identify_script_path = Alpha_script +"scripts/image_identify_executor.py"
const Identify_model_type  = "ResNet"
//Identify_model_type提供用于图像预测的4种算法包括 SqueezeNet，ResNet，InceptionV3 和 DenseNet
 ```

- alpha-core安装 

```golang
//编译
go build -o ./build/alpha-core  ./src/
// 启动
./build/alpha-core

```

接下来打印如下日志:

```
gaojiayideMacBook-Pro:alpha-core gaojiayi$ ./build/alpha-core 
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /alpha/v1/getb            --> alpha-core/src/controller.GetDataB (2 handlers)
[GIN-debug] POST   /alpha/image/identify     --> alpha-core/src/controller.ExecmagesIdentify (2 handlers)
[GIN-debug] POST   /alpha/image/detect       --> alpha-core/src/controller.ExecImagesDetect (2 handlers)
[GIN-debug] POST   /alpha/image/upload       --> alpha-core/src/controller.UploadImages (2 handlers)
[GIN-debug] POST   /alpha/about/message/add  --> alpha-core/src/controller.AddMessage (2 handlers)
[GIN-debug] GET    /alpha/about/message/query --> alpha-core/src/controller.QueryAllMessage (2 handlers)
[GIN-debug] GET    /alpha/resource/download  --> alpha-core/src/controller.FileDownload (2 handlers)
```

- alpha-ui安装  

```sh
npm run build 
然后将dist目录下的文件 部署至nginx 
```

### 功能特性  

- 图像识别
![recognition](assets/image-recognition.png)
- 图像检测
![detection](assets/image-detection.png)
当上传图片后，后端需要一定的时间进行分析，最后结果呈现如下：  
![detection-report](assets/image-detection-report.png)
- 留言板功能
![message-board](assets/message-board.png)
