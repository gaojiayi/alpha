<template>
  <div class="message-whole">
    <div class="demo-spin-container">
      <comment-message @handleReply="handleReply" :commentList="messages"></comment-message>
      <!-- 分页 -->
      <Page :total="totalPage" class="comment-page" @on-change="changePage" />
      <comment-area @reload="reload" :parentMsgId="replyMsgId" :replyMsgUsername="replyMsgUsername"></comment-area>
      <Spin size="large" fix v-if="messageSpinShow"></Spin>
    </div>
  </div>
</template>

<script>
import commentArea from "../../components/about_faq/commentArea";
import commentMessage from "../../components/about_faq/commentMessage";
import mock_message from "../../mock/comments";
import { createNamespacedHelpers } from "vuex";
const {
  mapState,
  mapMutations,
  mapActions,
  mapGetters
} = createNamespacedHelpers("message");
export default {
  components: {
    commentArea,
    commentMessage
  },
  data() {
    return {
      comments: mock_message.messages,
      replyMsgId: "",
      replyMsgUsername: "",
      pageIndex: 1,
      pageSize: 10
    };
  },
  computed: {
    ...mapState({
      messages: state => state.messages,
      messageSpinShow: state => state.messageSpinShow,
      totalPage: state => state.total
    })
  },
  created: function() {
    // 查询
    var params = {
      page_index: this.pageIndex,
      page_size: this.pageSize
    };
    this.QUERY_MESSAGE(params);
    //this.comments = mock_message.MessageResp;
  },
  methods: {
    ...mapActions(["QUERY_MESSAGE"]),
    handleReply(data) {
      this.replyMsgId = data.msgId;
      this.replyMsgUsername = data.msgUsername;
    },
    reload() {
      // 在上层index.js中添加reload方法 用于刷新留言板
      //this.$emit("reload")
    },
    changePage(newPage) {
      console.log(newPage);
      var params = {
        page_index: newPage,
        page_size: this.pageSize
      };
      this.pageIndex = newPage;
      this.QUERY_MESSAGE(params);
    }
  }
};
</script>
<style>
.demo-spin-container {
  display: inline-block;
  width: 100%;
  position: relative;
}
.comment-page{
  display: flex;
 justify-content:flex-end
}
</style>
