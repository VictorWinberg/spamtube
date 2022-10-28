import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import UploadPage from "../views/UploadPage.vue";
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
    path: "/upload",
    name: "UploadPage",
    component: UploadPage,
    meta: {
      order: 2,
    },
  },
  {
    path: "/stats",
    name: "StatsPage",
    component: StatsPage,
    meta: {
      order: 3,
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
