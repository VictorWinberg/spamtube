<template>
  <div class="stats">
    <img class="image" :src="graph" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import graph from "../assets/images/graph.svg";

export default defineComponent({
  name: "StatsPage",
  data() {
    return {
      graph,
    };
  },
});
</script>

<style scoped>
.stats {
  padding: 2em;
}

.image {
  max-width: 500px;
}
</style>
