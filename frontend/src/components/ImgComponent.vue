<template>
  <div class="img-component">
    <img :src="img.url" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "ImgComponent",
  props: ["img"],
  methods: {},
});
</script>

<style scoped lang="scss">


</style>
