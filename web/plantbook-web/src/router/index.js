import Vue from "vue";
import VueRouter from "vue-router";
import Main from "../views/Main.vue";
import Login from "../views/Login.vue";
import UserGallery from "../views/UserGallery.vue";
import UserPage from "../views/UserPage.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Main",
    component: Main,
  },
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/UserGallery",
    name: "UserGallery",
    component: UserGallery,
  },
  {
    path: "/UserPage",
    name: "UserPage",
    component: UserPage,
  },
];
const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
