import Vue from "vue";
import Vuex from "vuex";
import auth from "./modules/auth";
import gardens from "./modules/gardens";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: { auth, gardens },
});
