<template>
  <v-app>
    <v-main id="main">
      <router-view />
      <copyright-footer />
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { CopyrightFooter } from "./components";

@Component({
  components: {
    CopyrightFooter: CopyrightFooter
  }
})
export default class App extends Vue {}
</script>

<style>
#main {
  background: rgb(12, 0, 38);
  background: linear-gradient(
    6deg,
    rgb(6, 0, 19) 10%,
    rgb(11, 2, 41) 48%,
    rgb(11, 3, 49) 63%,
    rgb(0, 15, 68) 93%
  );
}
</style>
