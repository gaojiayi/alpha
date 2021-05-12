<template>
  <div>
    <a-card title="alpha项目代码" style="width: 100%">

      <a slot="extra" href="http://localhost/alpha-static/download/nginx-1.19.0.tar.gz"><Icon type="md-cloud-download" :style="{fontSize: '38px'}"/></a>
      <Collapse simple>
        <Panel name="1" hide-arrow="true">
            <h4>前端alpha-ui</h4>
            <p slot="content">史蒂夫·乔布斯（Steve Jobs），1955年2月24日生于美国加利福尼亚州旧金山，美国发明家、企业家、美国苹果公司联合创办人。</p>
        </Panel>
        <Panel name="2" hide-arrow="true">
            <h4>后端alpha-core</h4>
            <p slot="content">斯蒂夫·盖瑞·沃兹尼亚克（Stephen Gary Wozniak），美国电脑工程师，曾与史蒂夫·乔布斯合伙创立苹果电脑（今之苹果公司）。斯蒂夫·盖瑞·沃兹尼亚克曾就读于美国科罗拉多大学，后转学入美国著名高等学府加州大学伯克利分校（UC Berkeley）并获得电机工程及计算机（EECS）本科学位（1987年）。</p>
        </Panel>
        <Panel name="3" hide-arrow="true">
            <h4>图像AI分析alpha-script</h4>
            <p slot="content">乔纳森·伊夫是一位工业设计师，现任Apple公司设计师兼资深副总裁，英国爵士。他曾参与设计了iPod，iMac，iPhone，iPad等众多苹果产品。除了乔布斯，他是对苹果那些著名的产品最有影响力的人。</p>
        </Panel>
    </Collapse>
    </a-card>
  </div>
</template>
<script>
export default {
  methods:{

  }
}
</script>
<style>

</style>
