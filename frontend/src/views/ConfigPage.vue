<template>
  <div class="config">
    <h1>Configurations</h1>

    <v-list>
      <v-list-item>
        <div class="v-list-item__content">
          <v-btn
            size="small"
            class="icon-btn ma-2"
            color="green"
            icon="mdi-plus"
            @click="toggleNewConfiguration()"
          />
          <v-list-item-title>
            <strong>Create new configuration</strong>
          </v-list-item-title>
        </div>

        <ConfigFormComponent
          v-if="newConfigurationOpen"
          submitText="Create"
          @submit="create"
        />
      </v-list-item>

      <v-list-item v-for="item in query.data?.value?.parsedBody" :key="item.id">
        <div class="v-list-item__content">
          <div class="v-list-item__list-text">
            <v-list-item-title>
              <v-icon
                class="ma-2"
                size="small"
                :color="item.cron ? 'primary' : 'secondary'"
                :icon="item.cron ? 'mdi-play-circle' : 'mdi-pause-circle'"
              />
              <strong>{{ item.name }}</strong>
            </v-list-item-title>
            <v-list-item-subtitle v-if="item.cron">
              {{ cronstrue.toString(item.cron) }}
            </v-list-item-subtitle>
          </div>
          <v-btn
            class="icon-btn ma-2"
            variant="outlined"
            color="warning"
            size="small"
            icon="mdi-tools"
            @click="toggleConfigure(item)"
          />
          <v-btn
            class="icon-btn ma-2"
            variant="outlined"
            color="error"
            size="small"
            icon="mdi-trash-can"
            @click="openRemoveDialog(item)"
          />
        </div>

        <ConfigFormComponent
          v-if="selectedItemId === item.id"
          :item="item"
          submitText="Save"
          @submit="save"
        />
      </v-list-item>
    </v-list>

    <RemoveModal v-model="removeDialog" @cancel="cancel" @remove="remove" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import cronstrue from "cronstrue";
import { useQuery, useQueryClient } from "@tanstack/vue-query";
import ConfigFormComponent from "../components/ConfigFormComponent.vue";
import {
  getConfigurations,
  deleteConfigEntry,
  upsertConfig,
} from "../api/config";
import RemoveModal from "../modals/RemoveModal.vue";

interface ItemProps {
  id: string;
  name: string;
  cron: string | undefined;
  createdAt?: string;
}

export default defineComponent({
  name: "ConfigPage",
  components: {
    ConfigFormComponent,
    RemoveModal,
  },
  data() {
    return {
      cronstrue,
      queryClient: useQueryClient(),
      selectedItemId: "",
      removalItemId: "",
      nameTextfield: "",
      cronTextfield: "",
      removeDialog: false,
      newConfigurationOpen: false,
      items: [] as ItemProps[],
    };
  },
  methods: {
    toggleNewConfiguration() {
      this.nameTextfield = "";
      this.cronTextfield = "";
      if (this.newConfigurationOpen) {
        this.newConfigurationOpen = false;
      } else {
        this.selectedItemId = "";
        this.newConfigurationOpen = true;
      }
    },
    toggleConfigure(item: ItemProps) {
      if (this.selectedItemId === item.id) {
        this.selectedItemId = "";
        return;
      }
      if (this.newConfigurationOpen) {
        this.newConfigurationOpen = false;
      }
      this.selectedItemId = item.id;
      this.nameTextfield = item.name;
      this.cronTextfield = item.cron ? item.cron : "";
    },
    openRemoveDialog(item: ItemProps) {
      this.removeDialog = true;
      this.removalItemId = item.id;
    },
    async create(item: ItemProps) {
      try {
        await upsertConfig({
          name: item.name,
          cron: item.cron || undefined,
        });
        this.queryClient.invalidateQueries({ queryKey: ["items"] });
      } catch (error) {
        console.error(error);
      }
      this.newConfigurationOpen = false;
      this.nameTextfield = "";
      this.cronTextfield = "";
    },
    async save(item: ItemProps) {
      try {
        await upsertConfig({
          id: item.id,
          name: item.name,
          cron: item.cron || undefined,
        });
        this.queryClient.invalidateQueries({ queryKey: ["items"] });
      } catch (error) {
        console.error(error);
      }
      this.selectedItemId = "";
      this.nameTextfield = "";
      this.cronTextfield = "";
    },
    async remove() {
      try {
        await deleteConfigEntry(this.removalItemId);
        this.queryClient.invalidateQueries({ queryKey: ["items"] });
      } catch (error) {
        console.error(error);
      }
      this.selectedItemId = "";
      this.removeDialog = false;
    },
    cancel() {
      this.removeDialog = false;
    },
  },
  setup() {
    const query = useQuery({
      queryKey: ["items"],
      queryFn: getConfigurations,
    });
    return {
      query,
    };
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
  background: rgba(var(--v-theme-black), 0.5);
  color: rgb(var(--v-theme-white));
  text-align: left;
  max-width: 800px;

  &:not(:last-child) {
    .v-list-item__content {
      border-bottom: solid rgb(var(--v-theme-white), 0.2) 1px;
    }
  }

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
  }
}
</style>
