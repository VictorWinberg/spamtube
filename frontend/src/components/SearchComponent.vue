<template>
  <div class="search-component mx-auto">
    <h1>Search for Subreddits</h1>
    <v-form class="mt-4" @submit.prevent="fetchTopPosts">
      <v-combobox
        v-model="autocomplete"
        :items="items"
        label="Subreddit"
        variant="outlined"
        :hide-no-data="true"
        :hide-selected="true"
        dense
      />
      <v-btn type="submit" size="large" icon="mdi-magnify"></v-btn>
    </v-form>
    <div>
      <p class="font-weight-regular" v-if="noResponse">
        No Response, please try again...
      </p>
    </div>
    <v-expansion-panels v-model="openedPanel" variant="accordion">
      <v-expansion-panel
        v-for="post in posts"
        :key="post.data.id"
        class="mx-auto"
        @click="selectPost(post.data.id)"
        :class="{ selected: isSelected(post.data.id) }"
      >
        <v-expansion-panel-title>
          {{ post.data.title }}
          <template v-slot:actions>
            <v-btn text v-if="post.data.selftext" class="read-btn ml-4">
              <span v-if="isSelected(post.data.id)">Read less</span>
              <span v-else>Read more</span>
            </v-btn>
          </template>
        </v-expansion-panel-title>
        <v-expansion-panel-text v-if="post.data.selftext">
          {{ post.data.selftext }}
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>
    <v-btn
      @click="submitStep()"
      class="mt-4"
      :disabled="!selectedId"
      block
      size="x-large"
    >
      Continue
    </v-btn>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { TopPost } from "../api/reddit";
import { getTopPosts, searchSubreddits } from "../api/reddit";

interface DataProps {
  autocomplete: string;
  items: string[];
  posts: TopPost[];
  selectedId: string;
  openedPanel: number[];
  noResponse: boolean;
  isLoading: boolean;
}

export default defineComponent({
  name: "SearchComponent",
  data: (): DataProps => ({
    autocomplete: "",
    selectedId: "",
    openedPanel: [],
    noResponse: false,
    posts: [],
    items: [],
    isLoading: false,
  }),
  methods: {
    isSelected(id: string) {
      return this.selectedId === id;
    },
    selectPost(id: string) {
      if (this.selectedId === id) {
        this.openedPanel = [];
        this.selectedId = "";
      } else {
        this.selectedId = id;
      }
    },
    submitStep() {
      const post = this.posts.find((post) => this.isSelected(post.data.id));
      this.$emit("submitStep", post?.data);
    },
    async fetchTopPosts() {
      this.noResponse = false;
      this.openedPanel = [];
      this.selectedId = "";

      try {
        let response = await getTopPosts(this.autocomplete);
        this.posts = response.parsedBody || [];
      } catch (error) {
        this.noResponse = true;
        this.posts = [];
      }
    },
  },
  watch: {
    async autocomplete() {
      if (this.isLoading) return;

      this.isLoading = true;
      try {
        const response = await searchSubreddits(this.autocomplete);
        const items = response.parsedBody || [];
        this.items = items.slice(0, 5);
      } catch (error) {
        console.error(error);
      }
      this.isLoading = false;
    },
  },
});
</script>

<style scoped lang="scss">
.search-component {
  max-width: 800px;
}
.v-expansion-panels {
  gap: 0.5em;
}

.v-expansion-panel {
  transition: all 0.4s;
  background-color: rgb(var(--v-theme-secondary));
  color: rgb(var(--v-theme-lightText));
  &.selected {
    background-color: rgb(var(--v-theme-lightText));
    color: rgb(var(--v-theme-secondary));
  }
}
.v-expansion-panel--active:not(:first-child),
.v-expansion-panel--active + .v-expansion-panel {
  margin-top: 0;
}
form {
  display: flex;
  gap: 0.5em;
}

.v-btn {
  opacity: 0.8;
  color: rgb(var(--v-theme-darkText));
  background-color: rgb(var(--v-theme-button));

  &.read-btn {
    width: 130px;
  }
}
</style>
