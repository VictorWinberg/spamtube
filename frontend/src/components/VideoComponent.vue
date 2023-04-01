<template>
  <div class="video-component">
    <v-card class="mx-auto">
      <iframe
        width="280"
        height="280"
        :src="'https://www.youtube.com/embed/' + video.VideoId"
        :title="video.Title"
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
      ></iframe>

      <div class="video-stats mt-2">
        <v-card-subtitle>
          <v-icon
            v-for="n in rating(video)"
            :key="n"
            size="small"
            color="warning"
            :icon="'mdi-star'"
          />
          <p v-if="rating(video) == 0">No Rating</p>
        </v-card-subtitle>

        <div>
          <v-card-subtitle class="view-count">
            <v-icon class="mr-2" size="small" :icon="'mdi-eye'" />
            {{ video.Group.Community.Statistics.Views }}
          </v-card-subtitle>
        </div>
      </div>
      <div class="video-text">
        <v-card-title> {{ video.Title }} </v-card-title>
        <v-card-subtitle> {{ video.Group.Description }} </v-card-subtitle>
      </div>
    </v-card>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { YoutubeVideoData } from "../api/youtube";

export default defineComponent({
  name: "VideoComponent",
  props: ["video"],
  methods: {
    rating(video: YoutubeVideoData) {
      return Math.round(Number(video.Group.Community.StarRating.Average));
    },
  },
});
</script>

<style scoped lang="scss">
.video-component {
  display: flex;
  width: 280px;
}

.video-stats {
  display: flex;
  justify-content: space-between;
}

.v-card {
  display: flex;
  flex-direction: column;
  margin-bottom: 0;
}
.view-count {
  display: flex;
  align-items: center;
}
.video-text {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  padding-bottom: 1em;
}
</style>
