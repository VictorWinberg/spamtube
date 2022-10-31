<template>
  <div class="home">
    <h1>Your videos</h1>
    <div class="videos mt-8">
      <VideoComponent
        v-for="video in videos"
        :key="video.src"
        :title="video.title"
        :subtitle="video.description"
        :src="video.src"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import VideoComponent from "@/components/VideoComponent.vue";
import json from "../assets/videos.json";

interface DataProps {
  videos: {
    title: string;
    description: string;
    src: string;
  }[];
}

export default defineComponent({
  name: "HomePage",
  components: {
    VideoComponent,
  },
  data(): DataProps {
    return {
      videos: json,
    };
  },
});
</script>

<style scoped>
.home {
  padding: 2em;
}

.videos {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 2em;
}
</style>
