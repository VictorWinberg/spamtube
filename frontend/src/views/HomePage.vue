<template>
  <div class="home">
    <h1>This is a home page</h1>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "HomePage",
  data() {
    return {};
  },
});
</script>

<style scoped>
.home {
  padding: 2em;
}
</style>
