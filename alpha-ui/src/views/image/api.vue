<template>
  <div>
    <Divider>Alpha智能图像识别开发文档</Divider>
    <p>Alpha智能图像依托互联网大数据24小时训练自己的分析模型，并提供api开放给个人或者企业开发者,为更多的用户提供精准推荐.</p>
    <p>目前暂未支持平台自主接入，可线下联系，联系方式查看导航 关于->开发者社区</p>
    <p>收费标准: ￥1/1000次/月</p>
    <Divider orientation="left">接口描述</Divider>
    <p>地址: https://localhost:8080/alpha/api/image/identify</p>
    <p>描述: 识别图片特征,按照相识度排名降序返回识别结果</p>
    <p>方法: POST</p>
    <p>格式: application/json</p>
    <p>Basic Auth: 例 admin / admin</p>
    <Divider orientation="left">请求参数</Divider>
    <Table :columns="columns1" :data="data1"></Table>
    <div>
      <h4>请求示例:</h4>
      <pre id="jsonShow">
  {
      name: "John Brown",
      age: 18,
      address: "New York No. 1 Lake Park",
      date: "2016-10-03"
  }
    </pre>
    </div>
    <Divider orientation="left">响应参数</Divider>
    <Table :columns="columns1" :data="data1"></Table>
    <div>
      <h4>响应示例:</h4>
      <pre id="jsonShow">
  {
      name: "John Brown",
      age: 18,
      address: "New York No. 1 Lake Park",
      date: "2016-10-03"
  }
    </pre>
    </div>
  </div>
</template>
<script>
export default {
  // 目前只有一个接口  写死
  data() {
    return {
      columns1: [
        {
          title: "字段",
          key: "name"
        },
        {
          title: "名称",
          key: "age"
        },
        {
          title: "类型",
          key: "address"
        },
        {
          title: "必填",
          key: "address"
        },
        {
          title: "示例",
          key: "address"
        }
      ],
      data1: [
        {
          name: "John Brown",
          age: 18,
          address: "New York No. 1 Lake Park",
          date: "2016-10-03"
        },
        {
          name: "Jim Green",
          age: 24,
          address: "London No. 1 Lake Park",
          date: "2016-10-01"
        },
        {
          name: "Joe Black",
          age: 30,
          address: "Sydney No. 1 Lake Park",
          date: "2016-10-02"
        },
        {
          name: "Jon Snow",
          age: 26,
          address: "Ottawa No. 2 Lake Park",
          date: "2016-10-04"
        }
      ]
    };
  }
};
</script>
