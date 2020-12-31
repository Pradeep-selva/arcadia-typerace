<template>
  <v-container
    class="validate
    d-flex
    flex-column 
    justify-center 
    align-center"
  >
    <v-progress-circular
      :size="150"
      color="primary"
      indeterminate
      class="mb-16"
      v-if="loading"
    />
    <h1 v-if="loading" class="grey--text text--lighten-2">
      Validating your request...
    </h1>
    <h1 v-if="!!error" class="grey--text text--lighten-2">{{ error }}</h1>
    <v-btn
      v-if="!!error"
      color="grey lighten-1"
      outlined
      x-large
      class="mt-16"
      @click="onNavigateHome"
    >
      Go Home
    </v-btn>
  </v-container>
</template>

<script lang="ts">
import { ROUTES, VAL_TYPES } from "@/configs";
import { Vue, Component } from "vue-property-decorator";
import { validateOperation } from "../apis";

@Component
export default class Validate extends Vue {
  valType = this.$route.params.valType;
  roomId = this.$route.params.roomId;
  error = "";
  loading = false;

  async mounted() {
    this.loading = true;

    try {
      const response = await validateOperation(
        this.valType as keyof typeof VAL_TYPES,
        this.roomId
      );

      if (response.ok)
        this.$router.push({
          path: ROUTES.room(this.roomId)
        });
      else this.error = response.data;
    } catch (error) {
      this.error = error;
    } finally {
      this.loading = false;
    }
  }

  onNavigateHome() {
    this.$router.push({ path: ROUTES.home });
  }
}
</script>

<style>
.validate {
  height: 100vh;
}
</style>
