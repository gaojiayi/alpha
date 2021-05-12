<template>
  <div class="dev-community-div">
    <Divider><h2>你的支持是我前进的动力，感谢打赏</h2></Divider>
    <div class="img-qrcode">
      <img src="../../assets/wechat.jpg" />
      <img src="../../assets/alipay.jpg" />
    </div>

    <Divider><h2>关注社区</h2></Divider>
    <p>Alpha 智能图像技术交流群，提供项目资源文件下载，模型训练最新进展，以及接入开放API.</p>
    <div class="group-qrcode">
      <img src="../../assets/group.jpeg" />
    </div>
    <Divider dashed><h2>关于作者</h2></Divider>
    <p>
      Blog:
      <a href="https://www.cnblogs.com/gaojy/" target="_blank" >https://www.cnblogs.com/gaojy/</a>
    </p>
    <p>
      Email:
      <a href="#">775670823@qq.com</a>
    </p>
  </div>
</template>
<script>
export default {};
</script>
<style>
.img-qrcode.group-qrcode {
  position: relative;
  display: flex;
  justify-content: flex-start;
}
.img-qrcode img {
  width: 400px;
  height: 400px;
  margin-right: 50px;
}
.group-qrcode img {
  width: 600px;
  margin-right: 50px;
}
 .dev-community-div p{
  font-family: 'Times New Roman', Times, serif;
  font-size: 20px
}
</style>
