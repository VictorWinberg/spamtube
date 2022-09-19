<template>
  <div class="home">
    <h1>This is a home page</h1>
    <img :src="logo" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import logo from "../assets/logo.png";

export default defineComponent({
  name: "HomePage",
  data() {
    return {
      logo,
    };
  },
});
</script>
