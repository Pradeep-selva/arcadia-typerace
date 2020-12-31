import { ROUTES } from "@/configs";
import Vue from "vue";
import VueRouter, { RouteConfig } from "vue-router";
import { Home } from "../views";

Vue.use(VueRouter);

const routes: Array<RouteConfig> = [
  {
    path: ROUTES.home,
    name: "Home",
    component: Home
  },
  {
    path: ROUTES.room(),
    name: "Room",
    component: () => import("../views/Room.vue")
  },
  {
    path: ROUTES.validate,
    name: "Validate",
    component: () => import("../views/Validate.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
