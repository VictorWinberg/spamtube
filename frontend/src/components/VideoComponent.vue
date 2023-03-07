<template>
  <div class="video-component">
    <v-card class="mx-auto">
      <iframe
        width="280"
        height="280"
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
  display: flex;
  width: 280px;
}

.v-card {
  display: flex;
  flex-direction: column;
  margin-bottom: 0;
}
.video-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding-bottom: 1em;
}
</style>
