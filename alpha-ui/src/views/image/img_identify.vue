<template>
  <div class="container">
    <Form :label-width="0">
      <FormItem label>
        <div class="clearfix">
          <a-upload
            :action="upload_image_url"
            list-type="picture-card"
            :file-list="fileList"
            @preview="handlePreview"
            @change="handleChange"
            accept=".jpg, .png"
          >
            <div v-if="fileList.length < 4">
              <a-icon type="plus" />
              <div class="ant-upload-text">上传图片</div>
            </div>
          </a-upload>
          <a-modal :visible="previewVisible" :footer="null" @cancel="handleCancel">
            <img alt="example" style="width: 100%" :src="previewImage" />
          </a-modal>
        </div>
      </FormItem>
      <FormItem>
        <Button type="primary" @click="trigerIdentify()">识别</Button>
      </FormItem>
    </Form>
    <div>
      <a-spin tip="GPU正在计算..." :spinning="spinning" :delayTime="delayTime">
        <div :class="spinning?'spin-content':'spin-content-hidden'"></div>
      </a-spin>
      <div v-for="(item ,index) in identifyImages" :key="index">
        <Divider size="small" orientation="left">
          <h2>{{ item['file_name'] }}</h2>
        </Divider>
        <Table
          class="identify-table"
          :columns="columns"
          :data="item['identify_rets']"
          :showHeader="true"
          no-data-text="无法识别，请重新上传"
        ></Table>
      </div>
    </div>
  </div>
</template>
<script>
//import { mapState, mapMutations, mapActions } from "vuex";
import { createNamespacedHelpers } from "vuex";
import API from "../../api/api";
const {
  mapState,
  mapMutations,
  mapActions,
  mapGetters
} = createNamespacedHelpers("identify");
function getBase64(file) {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });
}
export default {
  data() {
    return {
      upload_image_url: API.uploadImage,
      delayTime: 60 * 1000,
      columns: [
        {
          title: "识别结果",
          key: "name",
          resizable: true
        },
        {
          title: "相似度",
          key: "percet",
          resizable: true
        }
      ],
      previewVisible: false,
      previewImage: ""
    };
  },
  computed: {
    ...mapState({
      filename: state => state.filename,
      fileList: state => state.fileList,
      identifyImages: state => state.identifyImages,
      spinning: state => state.spinning
    }),
    ...mapGetters(["GET_IDENTIFY_REQUEST"])
  },
  methods: {
    ...mapActions(["IDENTIFY_IMAGE"]),
    ...mapMutations(["SET_FILE_LIST", "CLEAR_IDENTIFY_IMAGES"]),
    trigerIdentify() {
      if (this.fileList.length <= 0) {
        this.$Notice.error({
          title: "请至少上传一张图片!",
          desc: ""
        });
      } else {
        this.CLEAR_IDENTIFY_IMAGES();
        var param = this.GET_IDENTIFY_REQUEST;
        this.IDENTIFY_IMAGE(param);
      }
    },
    // UPLOAD_IMAGE(){
    //因为已经进行局部命名，所以在全局store里面没法找到局部的action
    //   return this.$store.dispatch('UPLOAD_IMAGE ','woaini')
    // }

    handleCancel() {
      this.previewVisible = false;
    },
    async handlePreview(file) {
      if (!file.url && !file.preview) {
        file.preview = await getBase64(file.originFileObj);
      }
      this.previewImage = file.url || file.preview;
      this.previewVisible = true;
    },
    handleChange({ file, fileList, event }) {
      this.SET_FILE_LIST(fileList);
    }
  }
};
</script>
<style>
/* you can make up upload button and sample style by using stylesheets */
.ant-upload-select-picture-card i {
  font-size: 32px;
  color: #999;
}

.ant-upload-select-picture-card .ant-upload-text {
  margin-top: 8px;
  color: #666;
}
.identify-table {
  width: 100%;
}
.spin-content {
  text-align: center;
  padding: 30px;
  background: rgba(0, 0, 0, 0.05);
  border-radius: 4px;
}
.spin-content-hidden {
  display: none;
}
</style>
