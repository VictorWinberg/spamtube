<template>
  <div class="">
    <v-card class="video-component mx-auto">
      <iframe
        width="360"
        height="360"
        :src="src"
        :title="title"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
      ></iframe>

      <div class="video-text">
        <v-card-title> {{ title }} </v-card-title>
        <v-card-subtitle> {{ subtitle }} </v-card-subtitle>
      </div>
    </v-card>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "VideoComponent",
  props: ["src", "title", "subtitle"],
});
</script>

<style scoped lang="scss">
.video-component {
  width: 360px;
  height: 100%;
}

.v-card {
  display: flex;
  flex-direction: column;
}
.video-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding-bottom: 1em;
}
</style>
