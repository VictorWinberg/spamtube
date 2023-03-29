<template>
  <div class="specify-component mx-auto">
    <h1>Search video details</h1>
    Specify the details of the video you want to generate.
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
        <v-textarea
          v-model="voice"
          label="Voice"
          placeholder="Hello spamtubers..."
          :hint="'Voice length: ' + voice.length"
          persistent-hint
          variant="outlined"
        />
        <div class="btn-group">
          <v-btn class="back mt-8" size="x-large" @click="goBack"> Back </v-btn>
          <v-btn
            class="mt-8 submit"
            type="submit"
            size="x-large"
            :disabled="!title || !description || !image || !voice"
            >Submit</v-btn
          >
        </div>
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
    this.errorMessage = "";
    this.title = this.data?.title || "";
    this.description = this.data?.description || "";
    this.image = this.data?.keywords.join(" ") || "";
    this.voice = this.data?.selftext || "";
  },
  methods: {
    submitForm: async function () {
      if (
        this.title === "" ||
        this.description === "" ||
        this.image === "" ||
        this.voice === ""
      ) {
        this.errorMessage = "Please fill out all fields correctly";
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
    goBack: function () {
      this.$emit("back");
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

.btn-group {
  display: flex;
  gap: 1em;
}
.v-btn {
  flex: 1;
  opacity: 0.8;
  color: rgb(var(--v-theme-darkText));
  color: rgb(var(--v-theme-darkText));
  background-color: rgb(var(--v-theme-button));

  &.submit {
    background-color: rgb(var(--v-theme-button));
  }

  &.back {
    background-color: rgb(var(--v-theme-white));
  }
}
</style>
