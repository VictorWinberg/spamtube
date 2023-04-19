<template>
  <div class="specify-component mx-auto">
    <h1>Specify video details</h1>
    Specify the details of the video you want to generate.
    <div>
      <v-form class="mt-4" @submit.prevent="submitForm">
        <v-text-field
          v-model="title"
          label="Title"
          placeholder="Title of YouTube clip"
          prepend-inner-icon="mdi-format-text"
          variant="outlined"
        />
        <v-text-field
          v-model="description"
          label="Description"
          placeholder="Don't forget to like and subscribe"
          prepend-inner-icon="mdi-format-text"
          variant="outlined"
        />
        <v-combobox
          v-model="imageKeywords"
          label="Image generating text"
          placeholder="Happy cats abstract painting"
          prepend-inner-icon="mdi-image-area"
          variant="outlined"
          closable-chips
          multiple
          chips
        />
        <v-select
          v-model="voice"
          :items="voices"
          placeholder="Select voice"
          prepend-inner-icon="mdi-microphone"
          label="Voice"
          variant="outlined"
        />
        <v-combobox
          v-model="style"
          :items="styles"
          placeholder="Select style"
          prepend-inner-icon="mdi-auto-fix"
          label="Style"
          variant="outlined"
        />
        <v-combobox
          v-model="background"
          :items="backgrounds"
          placeholder="Select background"
          prepend-inner-icon="mdi-image-filter-hdr"
          label="Background"
          variant="outlined"
        />
        <v-textarea
          v-model="textContent"
          label="Text content"
          placeholder="Hello spamtubers..."
          variant="outlined"
        />
        <span class="pl-4 pt-2 d-flex text-medium-emphasis text-caption">
          Text content length: {{ textContent?.length || 0 }}
        </span>
        <div class="btn-group">
          <v-btn class="back mt-8" size="x-large" @click="goBack"> Back </v-btn>
          <v-btn
            class="mt-8 submit"
            type="submit"
            size="x-large"
            :disabled="!title || !description || !imageKeywords || !textContent"
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
  title?: string;
  description?: string;
  imageKeywords?: string[];
  voice?: string;
  voices: string[];
  style?: string;
  styles: string[];
  background?: string;
  backgrounds: string[];
  textContent?: string;
  errorMessage?: string;
  status?: number;
}

export default defineComponent({
  name: "SpecifyComponent",
  props: ["data"],
  data(): DataProps {
    return {
      title: "",
      description: "",
      imageKeywords: [],
      textContent: "",
      voice: undefined,
      voices: ["Matthew"],
      style: undefined,
      styles: [
        "cartoon",
        "comic book",
        "futuristic",
        "graffiti",
        "impressionism",
        "manga",
        "oil painting",
        "pencil sketch",
        "pop art",
        "surrealism",
        "watercolor",
      ],
      background: undefined,
      backgrounds: [
        "abstract",
        "apocalyptic",
        "beach",
        "bright neon lights",
        "city",
        "cyberpunk",
        "desert",
        "dystopia",
        "forest",
        "galaxy",
        "jungle",
        "lake",
        "mountain",
        "ocean",
        "ruins",
        "space",
        "underwater",
        "urban",
        "utopian",
        "volcano",
        "waterfall",
      ],
      errorMessage: "",
      status: undefined,
    };
  },
  updated() {
    this.errorMessage = "";
    this.title = this.data?.title;
    this.description = this.data?.description;
    this.imageKeywords = this.data?.keywords;
    this.textContent = this.data?.selftext;
  },
  methods: {
    submitForm: async function () {
      if (
        this.title === "" ||
        this.description === "" ||
        !this.imageKeywords ||
        this.textContent === ""
      ) {
        this.errorMessage = "Please fill out all fields correctly";
        return;
      }
      this.errorMessage = "";
      const post = {
        title: this.title,
        description: this.description,
        imageKeywords: this.imageKeywords,
        voice: this.voice,
        style: this.style,
        background: this.background,
        textContent: this.textContent,
      };
      this.$emit("submitStep", post);
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
