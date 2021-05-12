<template>
  <div>
    <a-comment class="area-comment">
      <!-- <a-avatar
        slot="avatar"
        src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
        alt="Han Solo"
      />-->
      <div slot="content">
        <a-form-item>
          <a-textarea :rows="4" :value="content" @change="handleChange" id="my-textarea" />
        </a-form-item>
        <a-form-item>
          <a-button html-type="submit" :loading="submitting" type="primary" @click="handleSubmit">留言</a-button>
          <a-button html-type="submit" @click="handleCancel">取消</a-button>
        </a-form-item>
      </div>
    </a-comment>
  </div>
</template>
<script>
import moment from "moment";
import mock_message from "../../mock/comments";
import { createNamespacedHelpers } from "vuex";
const {
  mapState,
  mapMutations,
  mapActions,
  mapGetters
} = createNamespacedHelpers("message");
export default {
  data() {
    return {
      content: "",
      submitting: false,
      parent_msg_id: this.parentMsgId,
      reply_msg_username: this.replyMsgUsername
    };
  },
  props: ['parentMsgId', 'replyMsgUsername'],

  watch: {
    replyMsgUsername() {
      document
        .querySelector("#my-textarea")
        .setAttribute("placeholder", "回复: " + "@" + this.replyMsgUsername);
        this.parent_username = this.replyMsgUsername
    },
    parentMsgId(){
      this.parent_msg_id = this.parentMsgId
    }
  },
  methods: {
    ...mapActions(["ADD_MESSAGE"]),
    handleSubmit() {
      if (!this.content) {
        return;
      }

      this.submitting = true;
      let attr = {
        context: this.content,
        parent_msg_id: this.parent_msg_id,
        parent_username: this.parent_username,
        username:
          "游客" +
          (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1),
      };
      this.ADD_MESSAGE(attr)
      this.submitting = false;
      this.content = "";
      document.querySelector("#my-textarea").setAttribute("placeholder", "");
    },
    handleChange(e) {
      console.log()
      this.content = e.target.value;
    },
    handleCancel() {
      this.content = "";
      this.parent_msg_id= ""
      this.reply_msg_username = ""
      document.querySelector("#my-textarea").setAttribute("placeholder", "");

    }
  }
};
</script>
<style scope>
.area-comment .ant-comment-avatar {
  display: none;
}
button {
  margin-right: 40px;
}
</style>
