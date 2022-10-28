<template>
  <div class="config-component">
    <h1>Search for Subreddits</h1>
    <v-form @submit.prevent="fetchTopPosts">
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
    <v-expansion-panels variant="accordion">
      <v-expansion-panel
        v-for="post in posts"
        :key="post.data.id"
        class="mx-auto"
        @click="selectPost(post.data.id)"
        :class="{ selected: post.data.id === selectedId }"
      >
        <v-expansion-panel-title>
          {{ post.data.title }}
          <template v-slot:actions>
            <v-btn text v-if="post.data.selftext">Read more</v-btn>
          </template>
        </v-expansion-panel-title>
        <v-expansion-panel-text v-if="post.data.selftext">
          {{ post.data.selftext }}
        </v-expansion-panel-text>
      </v-expansion-panel>
    </v-expansion-panels>
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
  noResponse: boolean;
  isLoading: boolean;
}

export default defineComponent({
  name: "ConfigComponent",
  props: {
    msg: String,
  },
  data: (): DataProps => ({
    autocomplete: "",
    selectedId: "",
    noResponse: false,
    posts: [],
    items: [],
    isLoading: false,
  }),
  methods: {
    selectPost(id: string) {
      if (this.selectedId === id) {
        this.selectedId = "";
      } else {
        this.selectedId = id;
      }
    },
    async fetchTopPosts() {
      this.noResponse = false;

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
.config-component {
  margin: 0 auto;
  max-width: 800px;
}
.v-expansion-panels {
  gap: 0.5em;
}

.v-expansion-panel {
  transition: all 0.4s;
  &.selected {
    background-color: #304362;
    color: white;
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
  color: black;
  background-color: #6eb2da;
}
</style>
