<style scoped>
.layout {
  border: 1px solid #d7dde4;
  background: #f5f7f9;
  position: relative;
  border-radius: 4px;
  overflow: hidden;
}
.layout-header-bar {
  background: #fff;
  box-shadow: 0 1px 1px rgba(0, 0, 0, 0.1);
}
.layout-logo {
  width: 100px;
  height: 64px;
  background-color: black;
  position: absolute;
  top: 70%;
  left: 50%;
  transform: translate(-50%, -50%);
}
</style>
<template>
  <div class="layout">
    <Sider :style="{position: 'fixed', height: '100vh', left: 0}">
      <div :style="{height:'63px',width:'100%',textAlign: 'center'}">
        <span class="logo-name">ALPHA</span>
      </div>
      <Menu active-name theme="dark" width="auto" :open-names="['1']">
        <Submenu name="1">
          <template slot="title">
            <Icon type="md-images"></Icon>图像AI
          </template>
          <MenuItem name="1-1" to="/img_identify">图像识别</MenuItem>
          <MenuItem name="1-2" to="/img_detect">图像检测</MenuItem>
          <MenuItem name="1-3" @click.native="todo_message">视频检测</MenuItem>
          <MenuItem name="1-4" @click.native="todo_message">模型训练</MenuItem>
          <MenuItem name="1-5" to="/image_api">开放API</MenuItem>
        </Submenu>
        <Submenu name="2">
          <template slot="title">
            <Icon type="md-reorder"></Icon>关于
          </template>
          <MenuItem name="2-1" to="/about/community">开发者社区</MenuItem>
          <MenuItem name="2-2" to="/about/comments">FAQ</MenuItem>
        </Submenu>
        <Submenu name="3">
          <template slot="title">
            <Icon type="md-paper" />资源
          </template>
          <MenuItem name="3-1" to="/download/resource">下载</MenuItem>
        </Submenu>
      </Menu>
      <div class="copyright-div">
        <p>Copyright © 2020 alpha.com</p>
        <p>Powered by Vue and Golang</p>
      </div>
    </Sider>
    <Layout :style="{marginLeft: '200px',height:'100vh',overflowY:'scroll'}">
      <Header :style="{background: '#fff', boxShadow: '0 2px 3px 2px rgba(0,0,0,.1)',}">
        <div :style="{display:'inline'}">
          <img
            class="context-img"
            src="./assets/logo.png"
            :style="{height:'32px',width:'32px',float:'left'}"
          />
          <h2 :style="{float:'left'}">阿尔法智能图像</h2>
        </div>
      </Header>
      <Content :style="{padding: ' 16px 16px'}">
        <Card :style="{height: '100%'}">
          <router-view></router-view>
        </Card>
      </Content>
      <!-- <Header :style="{background: '#fff', boxShadow: '0 2px 3px 2px rgba(0,0,0,.1)',height:'15%'}">
        <h2>广告位</h2>
      </Header>-->
    </Layout>
  </div>
</template>
<script>
export default {
  methods: {
            todo_message () {
                this.$Message.info('功能升级中,敬请期待！');
            }
        }
};
</script>
<style >
.context-img {
  position: relative;
  /* left: 50%; */
  top: 50%;
  transform: translate(-50%, -50%);
}
.logo-name {
  font-size: 45px;
  font-family: TimesNR;
  font-style: italic;
  text-align: center;
  /* padding-top: 0px; */
  color: #f5f5f57a;
}
.ivu-menu {
  font-size: 20px;
}
.ivu-menu-item {
  font-size: 18px;
}
.copyright-div {
  position: fixed;
  bottom: 0px;
  color: #e8e8e8;
  left: 3px;
  text-align: center;
}
</style>

