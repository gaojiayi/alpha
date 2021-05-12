import API from '../../api/api'
import request from '../../utils/alphaRequest'
/**
 * 图像识别
 */
const message = {
  namespaced: true,
  state: {
    messageSpinShow: false,
    messages: [],
    total: 100
  },
  mutations: {
    SET_FOLD_STATUS: (state, id) => {
      state.messages.forEach(item => {
        if (item.id == id) {
          item.fold = !item.fold
        }
      })
    },
    SET_MESSAGE: (state, messages) => {
      state.messages = messages
    },
    TRIGGER_SPIN: (state) => {
      state.messageSpinShow = !state.messageSpinShow
    },
    SET_TOTAL_PAGE: (state,count) => {
      state.total = count
    }

  },
  getters: {
  },
  actions: {
    // 新增留言或者回复
    ADD_MESSAGE({ dispatch,commit }, params) {
      request.post(API.addMessage, params).then(res => {
        // 再次执行一次查询
        var params = {
          "page_index": 0,
          "page_size": 10
        }
        dispatch('QUERY_MESSAGE',{ commit },params)
      }).catch(err => {
        console.log(err)
      })
    },
    // 查询留言
    QUERY_MESSAGE({ commit }, params) {

      console.log(JSON.stringify(params))
      request.get(API.queryMessage, params).then(res => {
        // 设置loading
        commit('TRIGGER_SPIN');
        // 再次执行一次查询

        var message  = res.data.MessageResp;
        message.forEach(item => {
          item['fold'] = false;
        });
        commit('SET_MESSAGE', message);
        commit('SET_TOTAL_PAGE',res.data.Total)
        commit('TRIGGER_SPIN');
      }).catch(err => {
        console.log(err)
      })
    }
  }

};
export default message;
