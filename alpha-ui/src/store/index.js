import Vue from 'vue';
import Vuex from 'vuex';
import identify from './modules/identify'
import detect from './modules/detect'
import message from './modules/message'
Vue.use(Vuex);

const store = new Vuex.Store({
  modules: {
    identify,
    detect,
    message,
  }
})

export default store;
