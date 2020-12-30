<template>
  <v-footer
    id="footer"
    class="d-flex 
    justify-center 
    black 
    font-weight-bold 
    grey--text
    text--lighten-2"
  >
    Copyright &copy; 2020
  </v-footer>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class CopyrightFooter extends Vue {}
</script>

<style>
#footer {
  position: fixed;
  bottom: 0px;
  min-width: 100vw;
  max-height: 35px;
}
</style>
