<template>
  <v-app>
    <v-main justify="center">
      <v-list>
        <router-link to="/">Home</router-link> |
        <router-link to="/config">Configuration</router-link> |
        <router-link to="/stats">Stats</router-link>
      </v-list>
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App",

  data() {
    return {
      //
    };
  },
});
</script>
