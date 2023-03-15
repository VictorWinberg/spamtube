<template>
  <div class="config">
    <h1>Configurations</h1>

    <v-list>
      <v-list-item :key="0">
        <div class="v-list-item__content">
          <v-btn
            size="small"
            class="icon-btn"
            color="green"
            icon="mdi-plus"
            @click="openNewConfiguration()"
          />
          <v-list-item-title>
            <strong>Create new configuration</strong>
          </v-list-item-title>
        </div>
        <div class="v-list-item__configure" v-if="showNewConfiguration()">
          <v-text-field
            v-model="subredditTextfield"
            label="Subreddit"
            placeholder="Write name of subreddit"
            variant="outlined"
          />
          <v-text-field
            v-model="periodicityTextfield"
            label="Periodicity"
            placeholder="Write your periodicity"
            variant="outlined"
          />
          <v-btn size="large" color="success" block @click="create()">
            Create
          </v-btn>
        </div>
      </v-list-item>

      <v-list-item v-for="item in items" :key="item.id">
        <div class="v-list-item__content">
          <v-btn
            variant="flat"
            disabled
            size="small"
            class="icon-btn"
            :color="item.subtitle ? 'primary' : 'secondary'"
            :icon="item.subtitle ? 'mdi-play-circle' : 'mdi-pause-circle'"
          />
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
            size="small"
            icon="mdi-cog"
            @click="configure(item)"
          />
          <v-btn
            class="icon-btn"
            color="error"
            size="small"
            icon="mdi-close-thick"
            @click="openRemoveDialog(item)"
          />
        </div>
        <div
          class="v-list-item__configure"
          v-if="showConfigurationDetails(item)"
        >
          <v-text-field
            v-model="subredditTextfield"
            label="Subreddit"
            placeholder="Write name of subreddit"
            variant="outlined"
          />
          <v-text-field
            v-model="periodicityTextfield"
            label="Periodicity"
            placeholder="Write your periodicity"
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
          <v-btn color="grey" @click="cancel()"> No </v-btn>
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
  subtitle: string | null;
}

export default defineComponent({
  name: "ConfigPage",
  data() {
    return {
      selectedItemId: -1,
      subredditTextfield: "",
      periodicityTextfield: "",
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
          subtitle: null,
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
    showConfigurationDetails(item: ItemProps) {
      return this.selectedItemId === item.id && !this.removeDialog;
    },
    showNewConfiguration() {
      return this.selectedItemId === 0 && !this.removeDialog;
    },
    openNewConfiguration() {
      if (this.selectedItemId === 0) {
        this.selectedItemId = -1;
      } else {
        this.selectedItemId = 0;
      }
    },
    configure(item: ItemProps) {
      if (this.selectedItemId === item.id) {
        this.selectedItemId = -1;
        return;
      }
      this.selectedItemId = item.id;
      this.subredditTextfield = item.title;
      this.periodicityTextfield = item.subtitle ? item.subtitle : "";
    },
    create() {
      this.items.push({
        id: this.items.length + 1,
        title: this.subredditTextfield,
        subtitle: this.periodicityTextfield || null,
      });
      this.selectedItemId = -1;
      this.subredditTextfield = "";
      this.periodicityTextfield = "";
    },
    save(item: ItemProps) {
      item.title = this.subredditTextfield;
      item.subtitle = this.periodicityTextfield || null;
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
    cancel() {
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
  margin: 0 auto;
  background: rgb(var(--v-theme-white));
  text-align: left;
  margin-bottom: 0.5em;
  max-width: 800px;

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

.v-avatar,
.icon-btn {
  margin: 0.5em;
}

.v-card-actions {
  display: flex;
  justify-content: space-evenly;

  .v-btn {
    margin: 0;
  }
}
</style>
