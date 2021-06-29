import axios from "axios";

export default {
  namespaced: true,
  state: {
    token: null,
    status: null,
    login: "",
  },

  mutations: {
    POST_LOGIN(state, result) {
      state.token = result.status;
      state.status = "success";
      state.login = result.login.login;
    },
    AUTH_ERROR(state) {
      state.status = "error";
    },
    USER_OUT(state) {
      state.status = null;
      state.token = null;
      state.login = "";
    },
  },

  actions: {
    async LOGIN({ commit }, user) {
      try {
        const data = JSON.stringify(user);
        let res = await axios.post(
          "/api/v1/user/login",
          data,

          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );
        let login = JSON.parse(res.config.data);
        let result = { status: res.statusText, login };
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
    register() {},
  },

  getters: {
    IS_LOGGED_IN: (state) => !!state.token,
    LOGIN: (state) => state.login,
  },
};
