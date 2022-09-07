<template>
  <v-app>
    <v-main justify="center">
      <v-toolbar app>
        <v-toolbar-title>
          <router-link to="/" style="cursor: pointer">SPAMTUBE</router-link>
        </v-toolbar-title>
        <v-spacer></v-spacer>
        <v-toolbar-items class="hidden-xs-only">
          <v-btn flat> <router-link to="/">Home</router-link> </v-btn>
          <v-btn flat>
            <router-link to="/config">Configuration</router-link>
          </v-btn>
          <v-btn flat>
            <router-link to="/stats">Stats</router-link>
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <router-view v-slot="{ Component }">
        <transition :name="$route.meta.transitionName">
          <component :is="Component" />
        </transition>
      </router-view>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",
  data() {
    return {};
  },
});
</script>

<style lang="scss">
body {
  margin: 0;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  position: absolute;
  width: 100%;
  transition: transform 0.5s ease;
}

.slide-left-leave-to,
.slide-right-enter-from {
  transform: translateX(-100vw);
}

.slide-left-enter-from,
.slide-right-leave-to {
  transform: translateX(100vw);
}
</style>
