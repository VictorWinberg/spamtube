<template>
  <div class="upload-component mx-auto">
    <h1>Finalizing your amazing creation</h1>
    You are almost done! <br />
    Review your video details and press <strong>UPLOAD</strong> to upload your
    video to YouTube.
    <div class="mt-8" :class="{ 'color-animation': uploaded }">
      <v-card class="mx-auto mb-16" max-width="500">
        <v-img :src="uploaded ? neonCatGif : neonCatPng"></v-img>
        <v-card-title>{{ title }}</v-card-title>
        <v-card-subtitle> {{ description }} </v-card-subtitle>
        <v-card-text>
          <div>{{ voice }}</div>
        </v-card-text>
        <v-divider class="mx-4"></v-divider>
        <v-card-title>Image generating words</v-card-title>
        <v-card-text>
          <v-chip-group class="justify-center">
            <v-chip v-for="(word, index) in imageWords" :key="index">
              {{ word }}
            </v-chip>
          </v-chip-group>
        </v-card-text>
        <v-btn
          v-if="!uploaded"
          class="upload-button"
          color="tetriary"
          size="x-large"
          prepend-icon="mdi-upload"
          @click="callUploadFlow"
        >
          Upload
        </v-btn>
        <v-btn
          v-else
          class="success-button"
          color="success"
          size="x-large"
          prepend-icon="mdi-check"
        >
          Done
        </v-btn>
      </v-card>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { startUploadFlow } from "@/api/upload";
import neonCatGif from "../assets/images/neon-cat.gif";
import neonCatPng from "../assets/images/neon-cat.png";

interface DataProps {
  title: string;
  description: string;
  image: string;
  voice: string;
  status?: number;
  neonCatGif: string;
  neonCatPng: string;
}

export default defineComponent({
  name: "UploadComponent",
  props: ["data"],
  data(): DataProps {
    return {
      title: this.data?.title || "",
      description: this.data?.description || "",
      image: this.data?.image || "",
      voice: this.data?.voice || "",
      status: undefined,
      neonCatGif,
      neonCatPng,
    };
  },
  computed: {
    imageWords(): string[] {
      return this.image.replace(/\s+$/, "").split(" ");
    },
    uploaded(): boolean {
      return this.status === 200;
    },
  },
  watch: {
    data(newData: DataProps) {
      this.title = newData.title;
      this.description = newData.description;
      this.image = newData.image;
      this.voice = newData.voice;
    },
  },
  methods: {
    async callUploadFlow() {
      try {
        const res = await startUploadFlow(
          this.title,
          this.description,
          this.image,
          this.voice
        );
        this.status = res.status;
      } catch (error) {
        console.error(error);
      }
    },
  },
});
</script>

<style lang="scss">
.upload-component {
  max-width: 800px;
}

.v-card-title {
  line-height: 1.2;
  white-space: break-spaces;
}

.v-chip-group {
  display: flex;
  flex-wrap: wrap;
}

.color-animation {
  position: relative;
  margin: auto;
  max-width: 500px;
  z-index: 0;
}

.color-animation::before,
.color-animation::after {
  content: "";
  position: absolute;
  left: -0;
  top: 0;
  background: linear-gradient(
    45deg,
    #fb0094,
    #0000ff,
    #00ff00,
    #ffff00,
    #ff0000,
    #fb0094,
    #0000ff,
    #00ff00,
    #ffff00,
    #ff0000
  );
  background-size: 400%;

  width: 100%;
  height: 100%;
  z-index: -1;
  border-radius: 4px;
  animation: steam 10s linear infinite;
}

.color-animation:after {
  filter: blur(50px);
}

.v-btn {
  margin-bottom: 1.5em;

  &.success-button {
    pointer-events: none;
  }
}

@keyframes steam {
  0% {
    background-position: 0 0;
  }
  50% {
    background-position: 400% 0;
  }
  100% {
    background-position: 0 0;
  }
}
</style>
