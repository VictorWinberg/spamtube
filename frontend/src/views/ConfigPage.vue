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
          <ConfigFormComponent submitText="Create" @submit="create" />
        </div>
      </v-list-item>

      <v-list-item v-for="item in items" :key="item.id">
        <div class="v-list-item__content">
          <div class="v-list-item__list-text">
            <v-list-item-title>
              <v-icon
                :color="item.periodicity ? 'primary' : 'secondary'"
                :icon="
                  item.periodicity ? 'mdi-play-circle' : 'mdi-pause-circle'
                "
                size="small"
              />
              <strong>{{ item.subreddit }}</strong>
            </v-list-item-title>
            <v-list-item-subtitle v-if="item.periodicity">
              {{ cronstrue.toString(item.periodicity) }}
            </v-list-item-subtitle>
          </div>
          <v-btn
            class="icon-btn"
            color="warning"
            size="small"
            icon="mdi-tools"
            @click="configure(item)"
          />
          <v-btn
            class="icon-btn"
            color="error"
            size="small"
            icon="mdi-trash-can"
            @click="openRemoveDialog(item)"
          />
        </div>
        <div
          class="v-list-item__configure"
          v-if="showConfigurationDetails(item)"
        >
          <ConfigFormComponent
            :item="items.find((i) => i.id === selectedItemId)"
            submitText="Save"
            @submit="save"
          />
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
import cronstrue from "cronstrue";
import ConfigFormComponent from "../components/ConfigFormComponent.vue";

interface ItemProps {
  id: number;
  subreddit: string;
  periodicity: string | null;
}

export default defineComponent({
  name: "ConfigPage",
  components: {
    ConfigFormComponent,
  },
  data() {
    return {
      cronstrue,
      selectedItemId: -1,
      subredditTextfield: "",
      periodicityTextfield: "",
      removeDialog: false,
      items: [
        {
          id: 1,
          subreddit: "Am I the asshole",
          periodicity: "0 14 * * *",
        },
        {
          id: 2,
          subreddit: "Sweden",
          periodicity: "0 20 * * 2,4",
        },
        {
          id: 3,
          subreddit: "Reddit News",
          periodicity: null,
        },
        {
          id: 4,
          subreddit: "Ask Reddit...",
          periodicity: "0 8,18 * * 1-5",
        },
      ],
    };
  },
  methods: {
    showNewConfiguration() {
      return this.selectedItemId === 0 && !this.removeDialog;
    },
    openNewConfiguration() {
      this.subredditTextfield = "";
      this.periodicityTextfield = "";
      if (this.selectedItemId === 0) {
        this.selectedItemId = -1;
      } else {
        this.selectedItemId = 0;
      }
    },
    showConfigurationDetails(item: ItemProps) {
      return this.selectedItemId === item.id && !this.removeDialog;
    },
    configure(item: ItemProps) {
      if (this.selectedItemId === item.id) {
        this.selectedItemId = -1;
        return;
      }
      this.selectedItemId = item.id;
      this.subredditTextfield = item.subreddit;
      this.periodicityTextfield = item.periodicity ? item.periodicity : "";
    },
    create(item: ItemProps) {
      this.items.push({
        id: this.items.length + 1,
        subreddit: item.subreddit,
        periodicity: item.periodicity || null,
      });
      this.selectedItemId = -1;
      this.subredditTextfield = "";
      this.periodicityTextfield = "";
    },
    save(item: ItemProps) {
      this.items = this.items.map((i) =>
        i.id === this.selectedItemId
          ? {
              ...item,
              subreddit: item.subreddit,
              periodicity: item.periodicity || null,
            }
          : i
      );
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

    .v-list-item-title {
      display: flex;
      align-items: center;
    }

    i {
      margin-right: 0.25em;
    }
  }

  &__configure {
    margin-top: 1.5em;

    .v-btn {
      margin-top: 1em;
    }
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
