<template>
  <v-container class="container d-flex">
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
          class="display-4 
        font-weight-bold 
        mb-5 grey--text 
        text--lighten-1"
        >
          Arcadia TypeRace
        </h1>

        <p
          class="overline
        font-weight-regular 
        grey--text
        text--lighten-2"
        >
          Join a room and challenge your friends!
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
          <v-btn class="mr-10" color="grey lighten-1" outlined x-large>
            Join Room
          </v-btn>
          <v-btn color="grey lighten-1" outlined x-large>
            Create Room
          </v-btn>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";

@Component
export default class Home extends Vue {}
</script>

<style>
.container {
  text-shadow: 1px 1px 2px rgba(150, 150, 150, 1);
}
</style>
