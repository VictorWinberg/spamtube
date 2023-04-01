<template>
  <v-text-field
    v-model="nameTextfield"
    label="Subreddit"
    placeholder="Write name of subreddit"
    variant="outlined"
    class="mb-2"
    :rules="[() => !!nameTextfield || 'Subreddit is required']"
  />

  <v-text-field
    v-model="cronTextfield"
    label="Periodicity"
    placeholder="Write your periodicity"
    variant="outlined"
    class="mb-2"
    :hint="cronToString()"
    persistent-hint
    :rules="[
      () => isCronValid() || !cronTextfield || 'Invalid cron expression',
    ]"
  />

  <v-btn
    size="x-large"
    color="success"
    block
    @click="submit"
    :disabled="!nameTextfield || !isCronValid()"
  >
    {{ submitText }}
  </v-btn>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import cronstrue from "cronstrue";
import { isValidCron } from "cron-validator";

export default defineComponent({
  name: "ConfigFormComponent",
  props: ["submitText", "item"],
  emits: ["submit"],
  data: () => ({
    nameTextfield: "",
    cronTextfield: "",
  }),
  created() {
    this.nameTextfield = this.item?.name || "";
    this.cronTextfield = this.item?.cron || "";
  },
  methods: {
    submit() {
      this.$emit("submit", {
        ...this.item,
        name: this.nameTextfield,
        cron: this.cronTextfield,
      });
    },
    isCronValid() {
      return isValidCron(this.cronTextfield) || !this.cronTextfield;
    },
    cronToString() {
      try {
        return cronstrue.toString(this.cronTextfield);
      } catch (e) {
        return;
      }
    },
  },
});
</script>

<style scoped lang="scss">
.form-component {
  max-width: 800px;
}
</style>
