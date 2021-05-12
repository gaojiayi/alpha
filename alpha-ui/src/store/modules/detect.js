import API from '../../api/api'
import request from '../../utils/alphaRequest'
/**
 * 图像识别
 */
const detect = {
  namespaced: true,
  state: {
    spinning: false,
    fileList: [],
    detectImages: []
  },
  mutations: {
    CLEAR_DETECT_IMAGES: (state) => {
      state.detectImages = []
    },
    SET_FILE_NAME: (state, filename) => {
      state.filename = filename
    },
    SET_DETECT_IMAGES: (state, data) => {
      state.detectImages = data
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
    GET_DETECT_REQUEST: (state) => {
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
    DETECT_IMAGE({ commit }, params) {
      // console.log(API.execIdentify(''))
      commit('SET_SPINNING', true)
      request.post(API.detectImage, params).then(res => {
        commit('SET_DETECT_IMAGES', res.data['image_detect_items'])
        commit('SET_SPINNING', false)
      })
        .catch(err => {
          console.log(err)
        })

    }
  }

};
export default detect;
