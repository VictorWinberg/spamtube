<template>
  <div class="specify-component mx-auto">
    <h1>Search for Subreddits</h1>
    <div>
      <v-form class="mt-4" @submit.prevent="submitForm">
        <v-text-field
          v-model="title"
          label="Title"
          placeholder="Title of YouTube clip"
          variant="outlined"
        />
        <v-text-field
          v-model="description"
          label="Description"
          placeholder="Don't forget to like and subscribe"
          variant="outlined"
        />
        <v-text-field
          v-model="image"
          label="Image generating text"
          placeholder="Happy cats abstract painting"
          variant="outlined"
        />
        <v-text-field
          v-model="voice"
          label="Voice"
          placeholder="Hello spamtubers..."
          :hint="'Voice length: ' + voice.length"
          persistent-hint
          variant="outlined"
        />
        <v-btn class="mt-8" type="submit" block size="x-large">Submit</v-btn>
        <p class="error" v-if="errorMessage">{{ errorMessage }}</p>
      </v-form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

interface DataProps {
  title: string;
  description: string;
  image: string;
  voice: string;
  errorMessage: string;
  status?: number;
}

export default defineComponent({
  name: "SpecifyComponent",
  props: ["data"],
  data(): DataProps {
    return {
      title: "",
      description: "",
      image: "",
      voice: "",
      errorMessage: "",
      status: undefined,
    };
  },
  updated() {
    this.title = this.data.title || "";
    this.description = this.data.description || "";
    this.image = this.data.image || "";
    this.voice = this.data.selftext.substring(0, 150) || "";
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

      const content = {
        title: this.title,
        description: this.description,
        image: this.image,
        voice: this.voice,
      };
      this.$emit("submitStep", content);
    },
  },
});
</script>

<style scoped lang="scss">
.specify-component {
  max-width: 800px;
}

.error {
  color: rgb(var(--v-theme-error));
}

.v-text-field {
  text-align: left;
  transition: none;
}

.v-btn {
  opacity: 0.8;
  color: rgb(var(--v-theme-darkText));
  background-color: rgb(var(--v-theme-button));
}
</style>
