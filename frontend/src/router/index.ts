import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import ConfigPage from "../views/ConfigPage.vue";
import HomePage from "../views/HomePage.vue";
import StatsPage from "../views/StatsPage.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: HomePage,
    meta: {
      order: 0,
    },
  },
  {
    path: "/config",
    name: "ConfigPage",
    component: ConfigPage,
    meta: {
      order: 1,
    },
  },
  {
    path: "/stats",
    name: "StatsPage",
    component: StatsPage,
    meta: {
      order: 2,
    },
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

router.beforeEach((to, from) => {
  const toPos = to.meta.order || 0;
  const fromPos = from.meta.order || 0;
  if (from.name) {
    to.meta.transitionName = toPos > fromPos ? "slide-left" : "slide-right";
  }
});

export default router;
