const appUrl = "http://localhost:8085";
export default {
  state: {
    token: null,
  },
  mutations: {},
  actions: {
    login(context, { login, password }) {
      let response = await fetch(appUrl + "/api/v1/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
        },
        body: JSON.stringify({ login, password }),
      });
    },
    logOut() {},
    register() {},
  },
  getters: {
    isLoggedIn: (state) => state.token,
  },
};
