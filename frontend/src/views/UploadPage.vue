<template>
  <div class="upload">
    <UploadComponent />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import UploadComponent from "@/components/UploadComponent.vue";
export default defineComponent({
  name: "UploadPage",
  components: {
    UploadComponent,
  },
});
</script>

<style scoped>
.upload {
  padding: 2em;
}
</style>
