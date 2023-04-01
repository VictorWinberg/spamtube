<template>
  <v-dialog width="auto">
    <v-card>
      <v-card-text> Are you sure you want to remove this item? </v-card-text>
      <v-card-actions>
        <v-btn color="grey" @click="$emit('cancel')"> No </v-btn>
        <v-btn color="error" @click="$emit('remove')"> Yes </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "RemoveModal",
  emits: ["cancel", "remove"],
});
</script>

<style scoped lang="scss">
.v-card-actions {
  display: flex;
  justify-content: space-evenly;

  .v-btn {
    margin: 0;
  }
}
</style>
