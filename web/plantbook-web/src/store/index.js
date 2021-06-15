const appUrl = "http://localhost:8085";

import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import VueAxios from "vue-axios";
// import auth from "./modules/auth";

Vue.use(VueAxios, axios);

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: null,
    status: null,
  },
  mutations: {
    POST_LOGIN(state, result) {
      state.token = result;
      state.status = "success";
    },
    AUTH_ERROR(state) {
      state.status = "error";
    },
    USER_OUT(state) {
      state.status = null;
      state.token = null;
    },
  },

  actions: {
    async LOGIN({ commit }, user) {
      try {
        const data = JSON.stringify(user);
        let res = await axios.post(appUrl + "/api/v1/user/login", data, {
          headers: {
            "Content-Type": "application/json",
          },
        });
        let result = res.statusText;
        console.log(res);
        commit("POST_LOGIN", result);
      } catch (e) {
        console.log("Ошибка", e);
        commit("AUTH_ERROR");
      }
    },
    async LOG_OUT({ commit }) {
      let user = null;
      commit("USER_OUT", user);
    },
  },
  register() {},

  getters: {
    IS_LOGGED_IN: (state) => !!state.token,
  },

  // modules: { auth },
});
