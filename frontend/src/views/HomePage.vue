<template>
  <div class="home">
    <img :src="logo" alt="logo" class="logo" />
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
import logo from "../assets/images/logo.svg";

interface DataProps {
  logo: string;
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
      logo,
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

.logo {
  max-width: 500px;
}

.videos {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  gap: 2em;
}
</style>
