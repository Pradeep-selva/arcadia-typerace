<template>
  <v-app>
    <v-main id="main">
      <home />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import Home from "./views/Home.vue";

@Component({
  components: {
    Home
  }
})
export default class App extends Vue {}
</script>

<style>
#main {
  background: rgb(12, 0, 38);
  background: linear-gradient(
    6deg,
    rgba(12, 0, 38, 1) 10%,
    rgba(12, 0, 50, 1) 48%,
    rgba(11, 1, 64, 1) 63%,
    rgba(0, 20, 92, 1) 93%
  );
}
</style>
