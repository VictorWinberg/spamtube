<template>
  <div class="upload">
    <v-tabs v-model="step" grow>
      <v-tab :value="1" :disabled="step < 1">Pick Reddit</v-tab>
      <v-tab :value="2" :disabled="step < 2">Specify details</v-tab>
      <v-tab :value="3" :disabled="step < 3">Wait for success</v-tab>
    </v-tabs>

    <v-window v-model="step">
      <v-window-item class="pa-4" :value="1">
        <SearchComponent @submitStep="step1" />
      </v-window-item>
      <v-window-item class="pa-4" :value="2">
        <SpecifyComponent :data="selectedPost" @submitStep="step2" />
      </v-window-item>
      <v-window-item class="pa-4" :value="3">
        <UploadComponent :data="uploadContent" />
      </v-window-item>
    </v-window>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { TopPost } from "../api/reddit";
import SearchComponent from "@/components/SearchComponent.vue";
import SpecifyComponent from "@/components/SpecifyComponent.vue";
import UploadComponent from "@/components/UploadComponent.vue";

interface PostProps {
  title: string;
  description: string;
  image: string;
  voice: string;
}

interface DataProps {
  step: number;
  selectedPost?: TopPost;
  uploadContent?: PostProps;
}

export default defineComponent({
  name: "UploadPage",
  components: {
    SearchComponent,
    SpecifyComponent,
    UploadComponent,
  },
  data(): DataProps {
    return {
      step: 1,
      selectedPost: undefined,
      uploadContent: undefined,
    };
  },

  methods: {
    step1(post: TopPost) {
      this.selectedPost = post;
      this.step += 1;
    },
    step2(content: PostProps) {
      this.uploadContent = content;
      this.step += 1;
    },
  },
});
</script>

<style lang="scss" scoped>
.upload {
  padding: 0 2em;
  @media screen and (max-width: 900px) {
    padding: 0em;
  }
}
</style>
