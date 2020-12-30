<template>
  <v-container class="d-flex">
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.png')"
          class="my-3"
          contain
          height="300"
        />
      </v-col>

      <v-col class="mb-16">
        <h1
          class="shadow-text
        display-4 
        font-weight-bold 
        mb-5 grey--text 
        text--lighten-1"
        >
          Arcadia TypeRace
        </h1>

        <p
          class="shadow-text
        overline
        font-weight-regular 
        grey--text
        text--lighten-2"
        >
          Join/Create a room and challenge your friends!
          <br />Find out who's faster. <br />
          <a
            href="https://github.com/Pradeep-selva/arcadia-typerace"
            target="_blank"
            rel="noreferrer noopener"
          >
            Source Code
          </a>
        </p>
        <v-row class="mt-16 d-flex justify-center">
          <join-dialog />
          <create-dialog />
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { JoinDialog, CreateDialog } from "../components";

@Component({
  components: {
    JoinDialog,
    CreateDialog
  }
})
export default class Home extends Vue {}
</script>

<style>
.shadow-text {
  text-shadow: 1px 1px 2px rgba(150, 150, 150, 1);
}
</style>
