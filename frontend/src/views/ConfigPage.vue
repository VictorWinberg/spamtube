<template>
  <div class="config">
    <ConfigComponent msg="config" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import ConfigComponent from "@/components/ConfigComponent.vue"; // @ is an alias to /src

export default defineComponent({
  name: "ConfigPage",
  components: {
    ConfigComponent,
  },
});
</script>

<style scoped>
.config {
  padding: 2em;
}
</style>
