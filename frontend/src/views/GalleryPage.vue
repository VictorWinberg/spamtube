<template>
  <div>
    <div>
      <ImgComponent v-for="img in images" :key="img.url" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { getImg, ImgData } from "../api/img";
import ImgComponent from "@/components/ImgComponent.vue";

interface DataProps {
  images: ImgData[];
}

export default defineComponent({
  name: "GalleryPage",
  components: {
    ImgComponent,
  },
  data(): DataProps {
    return {
      images: [],
    };
  },
  async created() {
    try {
      const response = await getImg();
      this.images = response.parsedBody || [];
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

</style>
