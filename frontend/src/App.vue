<template>
  <v-app>
    <nav>
      <router-link to="/" class="item">
        <img class="item__image" :src="logo" />
      </router-link>
      <div class="menu">
        <router-link to="/" class="item">
          <v-icon small> mdi-home </v-icon>
          <p>Home</p>
        </router-link>
        <router-link to="/upload" class="item">
          <v-icon small> mdi-upload </v-icon>
          <p>Upload</p>
        </router-link>
        <router-link to="/config" class="item">
          <v-icon x-large> mdi-cog </v-icon>
          <p>Config</p>
        </router-link>
      </div>
    </nav>
    <v-main justify="center">
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
import logo from "./assets/images/logo.svg";

export default defineComponent({
  name: "App",
  data() {
    return {
      logo,
    };
  },
});
</script>

<style lang="scss">
body {
  margin: 0;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  display: flex;
  flex-direction: column;
  text-align: center;
  height: 100vh;
}

nav {
  display: flex;
  justify-content: space-between;
  position: sticky;
  top: 0;
  background: rgb(var(--v-theme-background));
  font-family: "JonzeJonzing";
  font-size: 2em;
  height: 6em;
  transition: all 500ms;
  border-bottom: solid 3px rgb(var(--v-theme-primary));
  z-index: 2;

  @media screen and (max-width: 600px) {
    font-size: 1.5em;
    height: 3em;
  }

  > * {
    display: flex;
    align-items: center;
  }

  .menu {
    background: -webkit-linear-gradient(
      left,
      rgb(var(--v-theme-primary)),
      rgb(var(--v-theme-tetriary))
    );
    background-clip: text;
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    gap: 1em;
    font-size: 1em;
    margin-right: 1em;
    @media screen and (max-width: 600px) {
      -webkit-text-fill-color: unset;
    }

    p {
      display: block;
      &:hover {
        border-bottom: solid rgb(var(--v-theme-primary)) 3px;
      }
      @media screen and (max-width: 600px) {
        display: none;
      }
    }

    .v-icon {
      display: none;
      @media screen and (max-width: 600px) {
        color: rgb(var(--v-theme-primary));
        display: block;
        &:active {
          transform: scale(1.25);
        }
      }
    }
  }

  .item {
    cursor: pointer;

    &__image {
      max-width: 100%;
      max-height: 100%;
    }
  }
}

.v-main {
  flex: 1;
  background: linear-gradient(
    to bottom,
    rgb(var(--v-theme-background)),
    rgb(var(--v-theme-backgroundFade))
  );
  background-attachment: fixed;
  color: rgb(var(--v-theme-lightText));
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

@font-face {
  font-family: "JonzeJonzing";
  src: local("JonzeJonzing"),
    url(./assets/fonts/JonzeJonzing.ttf) format("truetype");
}
</style>
