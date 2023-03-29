<template>
  <v-text-field
    v-model="subredditTextfield"
    label="Subreddit"
    placeholder="Write name of subreddit"
    variant="outlined"
    :rules="[() => !!subredditTextfield || 'Subreddit is required']"
  />

  <v-text-field
    v-model="periodicityTextfield"
    label="Periodicity"
    placeholder="Write your periodicity"
    variant="outlined"
    :hint="cronToString()"
    persistent-hint
    :rules="[
      () => isCronValid() || !periodicityTextfield || 'Invalid cron expression',
    ]"
  />

  <v-btn
    size="large"
    color="success"
    block
    @click="submit"
    :disabled="!subredditTextfield || !isCronValid()"
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
  data: () => ({
    subredditTextfield: "",
    periodicityTextfield: "",
  }),
  created() {
    this.subredditTextfield = this.item?.subreddit || "";
    this.periodicityTextfield = this.item?.periodicity || "";
  },
  methods: {
    submit() {
      this.$emit("submit", {
        ...this.item,
        subreddit: this.subredditTextfield,
        periodicity: this.periodicityTextfield,
      });
    },
    isCronValid() {
      return (
        isValidCron(this.periodicityTextfield) || !this.periodicityTextfield
      );
    },
    cronToString() {
      try {
        return cronstrue.toString(this.periodicityTextfield);
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
