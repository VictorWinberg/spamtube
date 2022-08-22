<template>
  <div class="configPage">
    <Config msg="config" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Config from "@/components/Config.vue"; // @ is an alias to /src

export default defineComponent({
  name: "configPage",
  components: {
    Config,
  },
});
</script>
