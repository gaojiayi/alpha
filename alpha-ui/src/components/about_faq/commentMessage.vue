<template>
  <div id="commentMsg">
    <div v-if="isEmpty(commentList)" class="head-message">暂无留言内容</div>
    <div v-else class="head-message">
      <h1>
        <span>"口吐芬芳"</span>
        <span>实地现场</span>
      </h1>
    </div>
    <comment
      @handleReply="handleReply"
      v-for="(item1, index) in commentList"
      :key="'parent-' + index"
      :comment="item1"
    >
      <template #icon_fold v-if="!isEmpty(item1.children)">
        <Icon
          type="md-arrow-dropdown"
          class="icon-drop-down"
          v-if="item1.fold"
          @click="foldMessage(item1.id)"
        />
        <Icon type="md-arrow-dropup" class="icon-drop-up" v-else @click="foldMessage(item1.id)" />
      </template>

      <!-- 二层留言 -->
      <template #childComment v-if="!isEmpty(item1.children) && !item1.fold">
        <comment
          v-for="(item2, index) in item1.children"
          :key="'children-' + index"
          :comment="item2"
          @handleReply="handleReply"
        ></comment>
      </template>
    </comment>
  </div>
</template>

<script>
import Comment from "./coment";
import Vue from "vue";
import { createNamespacedHelpers } from "vuex";
const {
  mapState,
  mapMutations,
  mapActions,
  mapGetters
} = createNamespacedHelpers("message");
export default {
  name: "CommentMessage",
  components: {
    Comment
  },
  props: {
    commentList: {
      type: Array,
      default: []
    }
  },
  computed: {
    
  },
  methods: {
    ...mapMutations(["SET_FOLD_STATUS"]),
    foldMessage(id) {
      this.SET_FOLD_STATUS(id);
    },
    isEmpty(ls) {
      return ls.length === 0;
    },
    handleReply(data) {
      this.$emit("handleReply", {
        msgId: data.msgId,
        msgUsername: data.msgUsername
      });
    }
  }
};
</script>

<style scoped>
.head-message {
  font-size: 20px;
  text-align: center;
  display: flex;
  justify-content: center;
  align-items: center;
  font-family: "Luckiest Guy", cursive;
}
.icon-drop-down,
.icon-drop-up {
  font-size: 20px;
  margin-left: 10px;
  margin-bottom: 0;
  position: absolute;
  height: 15px;
}
h1 {
  margin: 0;
  font-size: 4em;
  padding: 0;
  color: #5cff6a;
  text-shadow: 0 0.1em 20px black, 0.05em -0.03em 0 black,
    0.05em 0.005em 0 black, 0em 0.08em 0 black, 0.05em 0.08em 0 black,
    0px -0.03em 0 black, -0.03em -0.03em 0 black, -0.03em 0.08em 0 black,
    -0.03em 0 0 black;
}
h1 span {
  -webkit-transform: scale(0.9);
  transform: scale(0.9);
  display: inline-block;
}
h1 span:first-child {
  -webkit-animation: bop 1s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards
    infinite alternate;
  animation: bop 1s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards infinite
    alternate;
}
h1 span:last-child {
  -webkit-animation: bopB 1s 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275)
    forwards infinite alternate;
  animation: bopB 1s 0.2s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards
    infinite alternate;
}

@-webkit-keyframes bop {
  0% {
    -webkit-transform: scale(0.9);
    transform: scale(0.9);
  }
  50%,
  100% {
    -webkit-transform: scale(1);
    transform: scale(1);
  }
}

@keyframes bop {
  0% {
    -webkit-transform: scale(0.9);
    transform: scale(0.9);
  }
  50%,
  100% {
    -webkit-transform: scale(1);
    transform: scale(1);
  }
}
@-webkit-keyframes bopB {
  0% {
    -webkit-transform: scale(0.9);
    transform: scale(0.9);
  }
  80%,
  100% {
    -webkit-transform: scale(1) rotateZ(-3deg);
    transform: scale(1) rotateZ(-3deg);
  }
}
@keyframes bopB {
  0% {
    -webkit-transform: scale(0.9);
    transform: scale(0.9);
  }
  80%,
  100% {
    -webkit-transform: scale(1) rotateZ(-3deg);
    transform: scale(1) rotateZ(-3deg);
  }
}
</style>
