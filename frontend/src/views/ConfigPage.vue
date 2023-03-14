<template>
  <div class="config">
    <h1>Configuration</h1>

    <v-list lines="three">
      <v-list-item v-for="item in items" :key="item.id">
        <div class="v-list-item__content">
          <v-avatar color="surface-variant">
            <v-icon color="white">mdi-play-circle</v-icon>
          </v-avatar>
          <div class="v-list-item__list-text">
            <v-list-item-title>
              <strong>{{ item.title }}</strong>
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ item.subtitle }}
            </v-list-item-subtitle>
          </div>
          <v-btn
            class="icon-btn"
            color="warning"
            icon="mdi-cog"
            @click="configure(item)"
          />
          <v-btn
            class="icon-btn"
            color="error"
            icon="mdi-close-thick"
            @click="openRemoveDialog(item)"
          />
        </div>
        <div class="v-list-item__configure" v-if="showConfigDetails(item)">
          <v-text-field
            v-model="configureTextfield1"
            label="Subreddit"
            placeholder="Name of subreddit"
            variant="outlined"
          />
          <v-text-field
            v-model="configureTextfield2"
            label="Title"
            placeholder="Title of YouTube clip"
            variant="outlined"
          />
          <v-btn size="large" color="success" block @click="save(item)">
            Save
          </v-btn>
        </div>
      </v-list-item>
    </v-list>

    <v-dialog v-model="removeDialog" width="auto">
      <v-card>
        <v-card-text> Are you sure you want to remove this item? </v-card-text>
        <v-card-actions>
          <v-btn color="grey" @click="doNotRemove()"> No </v-btn>
          <v-btn color="error" @click="remove(selectedItemId)"> Yes </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

interface ItemProps {
  id: number;
  title: string;
  subtitle: string;
}

export default defineComponent({
  name: "ConfigPage",
  data() {
    return {
      selectedItemId: -1,
      configureTextfield1: "",
      configureTextfield2: "",
      removeDialog: false,
      items: [
        {
          id: 1,
          title: "Am I the asshole",
          subtitle: "Mondays at 15pm",
        },
        {
          id: 2,
          title: "Sweden",
          subtitle: "Every day at 2pm",
        },
        {
          id: 3,
          title: "Reddit News",
          subtitle: "Twice a day, at 8am and 6pm",
        },
        {
          id: 4,
          title: "Ask Reddit...",
          subtitle: "Tuesday and Thursday at 10pm",
        },
      ],
    };
  },
  methods: {
    showConfigDetails(item: ItemProps) {
      return this.selectedItemId === item.id && !this.removeDialog;
    },
    configure(item: ItemProps) {
      if (this.selectedItemId === item.id) {
        this.selectedItemId = -1;
        return;
      }
      this.selectedItemId = item.id;
      this.configureTextfield1 = item.title;
      this.configureTextfield2 = item.subtitle;
    },
    save(item: ItemProps) {
      item.title = this.configureTextfield1;
      item.subtitle = this.configureTextfield2;
      this.selectedItemId = -1;
    },
    openRemoveDialog(item: ItemProps) {
      this.removeDialog = true;
      this.selectedItemId = item.id;
    },
    remove(item: number) {
      this.items = this.items.filter((i) => i.id !== item);
      this.removeDialog = false;
    },
    doNotRemove() {
      this.selectedItemId = -1;
      this.removeDialog = false;
    },
  },
});
</script>

<style scoped lang="scss">
.config {
  padding: 2em;
}

.v-list {
  background: none;
}

.v-list-item {
  background: rgb(var(--v-theme-white));
  text-align: left;
  margin-bottom: 0.5em;

  .v-avatar,
  .icon-btn {
    margin: 0.5em;
  }

  &__list-text {
    flex: 1;
  }

  &__content {
    display: flex;
    align-items: center;
  }

  &__configure {
    margin-top: 1.5em;
  }
}

.v-card-actions {
  display: flex;
  justify-content: space-evenly;
}
</style>
