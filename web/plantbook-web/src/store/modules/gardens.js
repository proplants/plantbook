import axios from "axios";

export default {
  namespaced: true,
  state: {
    dataGardens: null,
  },
  mutations: {
    DATA_GARDENS(state, result) {
      state.dataGardens = result;
    },
  },
  actions: {
    async GARDENS({ commit }) {
      try {
        const response = await axios.get("http://localhost:3000/data");
        const result = response.data;
        commit("DATA_GARDENS", result);
      } catch (error) {
        console.error(error);
      }
    },
  },
  getters: {
    GET_GARDENS: (state) => state.dataGardens,
  },
};
