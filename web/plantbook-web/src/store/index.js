const appUrl = "http://localhost:8085";
import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import VueAxios from "vue-axios";

Vue.use(VueAxios, axios);

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: null,
  },
  mutations: {
    POST_LOGIN(state, result) {
      state.token = result;
    },
  },

  actions: {
    async LOGIN({ commit }, user) {
      const data = JSON.stringify(user);
      console.log(data);
      await axios
        .post(appUrl + "/api/v1/user/login", data, {
          headers: { "Content-Type": "application/json" },
        })
        .then((response) => {
          let result = response.json();
          console.log(result);
          commit("POST_LOGIN", result);
        })
        .catch((e) => {
          console.log("Ошибка", e);
        });
    },
    logOut() {},
    register() {},
  },
  getters: {
    isLoggedIn: (state) => state.token,
  },
  modules: {},
});
