import API from '../../api/api'
import request from '../../utils/alphaRequest'
/**
 * 图像识别
 */
const identify = {
  namespaced: true,
  state: {
    filename: 'hdsjhdjs.txt',
    spinning: false,
    fileList: [],
    identifyImages: []
  },
  mutations: {
    CLEAR_IDENTIFY_IMAGES: (state) => {
      state.identifyImages = []
    },
    SET_FILE_NAME: (state, filename) => {
      state.filename = filename
    },
    SET_IDENTIFY_IMAGES: (state, data) => {
      state.identifyImages = data
    },
    SET_FILE_LIST: (state, fileList) => {
      fileList.forEach(element => {
        delete element['lastModified'];
        delete element['lastModifiedDate'];
        delete element['lastModified'];
        delete element['thumbUrl'];
        delete element['xhr'];

      });
      state.fileList = fileList
    },
    SET_SPINNING: (state, spin) => state.spinning = spin

  },
  getters: {
    GET_IDENTIFY_REQUEST: (state) => {
      var request = new Array()
      state.fileList.forEach(
        element => {
          if (element['status'] == 'done') {
            var item = {};
            item['file_name'] = element['name'];
            item['file_save_name'] = element['response']['save_name'];
            console.log(item);
            request.push(item);
          }

        }
      )
      return JSON.stringify(request);
    }
  },
  actions: {
    IDENTIFY_IMAGE({ commit }, params) {
      // console.log(API.execIdentify(''))
      commit('SET_SPINNING', true)
      request.post(API.identifyImage, params).then(res => {
        commit('SET_IDENTIFY_IMAGES', res.data)
        commit('SET_SPINNING', false)
      })
        .catch(err => {
          console.log(err)
        })

    }
  }

};
export default identify;
