<template>
  <div class="upload-component">
    <h3>Enter parameters for new video</h3>
    <form @submit.prevent="submitForm">
      <span>Title</span><br />
      <input
        v-model="title"
        type="text"
        placeholder="Title of YouTube clip"
      /><br />
      <span>Description</span><br />
      <input
        v-model="description"
        type="text"
        placeholder="Description"
      /><br />
      <span>Image generating text</span><br />
      <input
        v-model="image"
        type="text"
        placeholder="Happy cats abstract painting"
      /><br />
      <span>Voice</span><br />
      <input
        v-model="voice"
        type="text"
        placeholder="Hello spamtubers..."
      /><br />
      <div>Voice length: {{ voice.length }}</div>
      <input class="submit" type="submit" value="Submit" />
      <div class="error" v-if="errorMessage">{{ errorMessage }}</div>
    </form>
    <div v-if="status !== undefined">
      <h3>Posted with status</h3>
      <p>Status: {{ status }}</p>
      <p>Title: {{ title }}</p>
      <p>Description: {{ description }}</p>
      <p>Image: {{ image }}</p>
      <p>Voice: {{ voice }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { startUploadFlow } from "@/api/upload";

interface Data {
  title: string;
  description: string;
  image: string;
  voice: string;
  errorMessage: string;
  status?: number;
}

export default defineComponent({
  name: "UploadComponent",
  props: {},
  data(): Data {
    return {
      title: "",
      description: "",
      image: "",
      voice: "",
      errorMessage: "",
      status: undefined,
    };
  },
  methods: {
    submitForm: async function () {
      if (
        this.title === "" ||
        this.description === "" ||
        this.image === "" ||
        this.voice === ""
      ) {
        this.errorMessage = "Please fill out all fields";
        return;
      }
      this.errorMessage = "";
      const res = await startUploadFlow(
        this.title,
        this.description,
        this.image,
        this.voice
      );
      this.status = res.status;
    },
  },
});
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
input {
  border-color: #42b983;
  border-style: solid;
  color: white;
}
.error {
  color: red;
}
</style>
