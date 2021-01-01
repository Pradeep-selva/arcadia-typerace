<template>
  <v-container
    class="validate
   d-flex 
   justify-center 
   flex-column
   mt-16"
    fluid
  >
    <h1
      class="grey--text 
    text--lighten-1 
    d-flex 
    justify-center"
    >
      ROOM--
      <span class="grey--text text--lighten-2">{{ $route.params.id }}</span>
    </h1>
    <div
      class="d-flex 
    justify-center 
    mb-10 mt-5"
    >
      <v-btn class="start-btn" color="grey lighten-1 mb-16" outlined x-large>
        start race
      </v-btn>
    </div>
    <v-divider light class="mb-10" id="divider" />
    <v-row justify="space-around" class="d-flex flex-row">
      <v-col class="editor" sm="5"></v-col>
      <v-col class="editor" sm="5"></v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";

@Component
export default class Room extends Vue {
  mounted() {
    console.log(this.$route);
  }
}
</script>

<style>
.editor {
  background-color: #090212;
  border-radius: 10px;
  max-height: 50vh;
}

.start-btn {
  max-width: 10vw;
}

#divider {
  border-color: #8e8c91;
}
</style>
