<template>
  <div class="home">
    <div class="videos mt-8">
      <VideoComponent
        v-for="video in videos"
        :key="video.VideoId"
        :video="video"
      />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import VideoComponent from "@/components/VideoComponent.vue";
import { getUploadedYoutubeVideos, YoutubeVideoData } from "../api/youtube";

interface DataProps {
  videos: YoutubeVideoData[];
}

export default defineComponent({
  name: "HomePage",
  components: {
    VideoComponent,
  },
  data(): DataProps {
    return {
      videos: [],
    };
  },
  async created() {
    try {
      const response = await getUploadedYoutubeVideos();
      this.videos = response.parsedBody || [];
    } catch (error) {
      console.error(error);
    }
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
