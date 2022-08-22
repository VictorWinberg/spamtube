import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import ConfigPage from "../views/ConfigPage.vue";
import HomePage from "../views/HomePage.vue";
import StatsPage from "../views/StatsPage.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: HomePage,
  },
  {
    path: "/config",
    name: "ConfigPage",
    component: ConfigPage,
  },
  {
    path: "/stats",
    name: "StatsPage",
    component: StatsPage,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
