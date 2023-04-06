<template>
  <v-navigation-drawer
    class="menu"
    v-model="drawer"
    :rail="rail"
    floating
    permanent
  >
    <v-list-item prepend-icon="mdi-menu" title="Menu" @click="rail = !rail">
    </v-list-item>

    <v-divider></v-divider>

    <v-list density="compact" nav>
      <router-link to="/">
        <v-list-item
          class="menu-item"
          prepend-icon="mdi-home"
          title="Home"
          value="home"
        />
      </router-link>
      <router-link to="/upload">
        <v-list-item
          class="menu-item"
          prepend-icon="mdi-upload"
          title="Upload"
          value="upload"
        />
      </router-link>
      <router-link to="/config">
        <v-list-item
          class="menu-item"
          prepend-icon="mdi-cog"
          title="Config"
          value="config"
        />
      </router-link>
    </v-list>
  </v-navigation-drawer>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "MenuComponent",
  data: () => ({
    drawer: true,
    rail: true,
  }),
  computed: {
    // currentComponent(): string {
    //   return this.$vuetify.breakpoint.xsOnly ? "MobileLayout" : "DesktopLayout";
    // },
  },
});
</script>

<style lang="scss">
.menu {
  background-color: rgba(var(--v-theme-black), 0.5);
  color: white;
}
.menu-item {
  color: white;
}
</style>
