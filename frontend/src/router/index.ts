import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import ConfigPage from "../views/ConfigPage.vue";
import Home from "../views/Home.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/config",
    name: "ConfigPage",
    component: ConfigPage,
  },
  {
    path: "/stats",
    name: "Stats",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/Stats.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
