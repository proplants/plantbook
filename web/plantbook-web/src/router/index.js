import Vue from "vue";
import VueRouter from "vue-router";
import Main from "../views/Main.vue";
import Login from "../views/Login.vue";
import UserGallery from "../views/UserGallery.vue";
import UserPage from "../views/UserPage.vue";
import NotFound from "../views/NotFound.vue";
import Plant from "../views/Plant.vue";

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
  {
    path: "/Plant/:title",
    name: "Plant",
    component: Plant,
  },
  {
    path: "/404",
    name: "NotFound",
    component: NotFound,
  },
  { path: "*", redirect: "/404" },
];
const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
