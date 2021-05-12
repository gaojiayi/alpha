<template>
  <a-comment>
    <span
      slot="actions"
      key="comment-basic-reply-to"
      @click="handlReply(comment)"
      class="comment-basic-reply-to"
    >
      <a href="#my-textarea">回复</a>
      <!-- 此处设置展开 折叠按钮 -->
      <slot name="icon_fold"></slot>
    </span>

    <a slot="author" style="font-size: 15px">{{comment.username}}{{relyer}}</a>
    <a
      v-if="comment.parent_msg_username"
      slot="author"
      class="reply-to"
    >@{{comment.parent_msg_username}}</a>
    <a-avatar
      slot="avatar"
      :src="'http://localhost/alpha-static/headphotos/'+comment.head_photo+'.jpg'"
      alt
    />
    <p
      slot="content"
      :class="isEmpty(comment.parent_msg_id)? 'comment-parent-context comment-context': 'comment-child-context comment-context'"
    >{{comment.context}}</p>
    <a-tooltip slot="datetime">
      <span>{{comment.date}}</span>
    </a-tooltip>
    <slot name="childComment"></slot>
  </a-comment>
</template>

<script>
export default {
  name: "Comment",
  props: {
    comment: ""
  },
  computed: {
    relyer: function() {
      if (this.comment.parent_msg_id != "") {
        return "@" + this.comment.parent_username;
      }
      return "";
    }
  },
  methods: {
    handlReply(comment) {
      // 暂时保持2层消息，优先获取root message id
      var msgId = comment.id;
      if (!this.isEmpty(comment.parent_msg_id)) {
        msgId = comment.parent_msg_id;
      }
      this.$emit("handleReply", { msgId, msgUsername: comment.username });
    },
    isEmpty(ls) {
      return ls.length === 0;
    }
  }
};
</script>

<style scoped>
.reply-to {
  padding-left: 5px;
  color: #409eff;
  font-weight: 500;
  font-size: 15px;
}
.comment-basic-reply-to {
  width: 100%;
}
.comment-context {
  width: 40%;
}
.comment-parent-context {
  font-family: "Courier New", Courier, monospace;
  font-weight: 600;
  font-style: normal;
  font-size: 18px;
}
.comment-child-context {
  font-family: Impact, Haettenschweiler, "Arial Narrow Bold", sans-serif;
  font-weight: normal;
  font-style: italic;
}
</style>
