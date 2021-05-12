<template>
  <div class="container">
    <Form>
      <FormItem>
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
        <Button type="primary" @click="trigerDetect()">探测</Button>
      </FormItem>
    </Form>
    <a-spin tip="GPU正在计算..." :spinning="spinning" :delayTime="delayTime">
      <div :class="spinning?'spin-content':'spin-content-hidden'"></div>
    </a-spin>
    <div v-for="item in detectImages" :key="index">
      <Divider size="small" orientation="left">
        <h2>{{ item['file_name'] }}</h2>
      </Divider>
      <div class="detect-result-card">
        <a-card
          v-for="(image, key, index) in item['image_detect_objects']"
          :key="index"
          hoverable
          style="width: 240px;height: 300px;"
        >
          <img slot="cover" alt="example" :src="image['url']" />
          <a-card-meta :title="image['name']">
            <template slot="description">相似度:{{image['percet']}}</template>
          </a-card-meta>
        </a-card>
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
} = createNamespacedHelpers("detect");
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
      previewVisible: false,
      previewImage: ""
    };
  },
  computed: {
    ...mapState({
      fileList: state => state.fileList,
      detectImages: state => state.detectImages,
      spinning: state => state.spinning
    }),
    ...mapGetters(["GET_DETECT_REQUEST"])
  },
  methods: {
    ...mapActions(["DETECT_IMAGE"]),
    ...mapMutations(["SET_FILE_LIST", "CLEAR_DETECT_IMAGES"]),
    trigerDetect() {
      if (this.fileList.length <= 0) {
        this.$Notice.error({
          title: "请至少上传一张图片!",
          desc: ""
        });
      } else {
        this.CLEAR_DETECT_IMAGES();
        var param = this.GET_DETECT_REQUEST;
        this.DETECT_IMAGE(param);
      }
    },

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
    handleChange({ fileList }) {
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
.detect-result-card {
  display: inline-flex;
  justify-content: flex-start;
  flex-wrap: wrap;
  align-items: flex-start;
  align-content: flex-start;
}
.detect-result-card img {
  height: 200px;
}
.ant-card.ant-card-bordered {
  margin-right: 50px;
  margin-bottom: 20px;
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
