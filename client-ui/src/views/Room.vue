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
      <span class="mr-3">ROOM</span>
      <span class="grey--text text--lighten-2">{{ roomId }}</span>
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
      <v-col class="editor" sm="5">
        <div class="grey--text text--lighten-2 display-2">
          <span class="green--text">
            {{ completeString }}
          </span>
          <span>
            {{ incompleteString }}
          </span>
        </div>
      </v-col>
      <v-col class="editor" sm="5">
        <div class="grey--text text--lighten-2 display-2">
          <span class="green--text">
            {{ completeString }}
          </span>
          <span>
            {{ incompleteString }}
          </span>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import RandomWords from "random-words";
import { API_ENDPOINTS } from "@/configs";

@Component
export default class Room extends Vue {
  roomId = this.$route.params.id;
  incompleteString = (RandomWords(30) as string[]).join(" ");
  completeString = "";
  enteredText = "";

  mounted() {
    document.addEventListener("keypress", (event) => {
      this.enteredText += event.key;
    });
  }
}
</script>

<style>
.editor {
  background-color: #090212;
  border-radius: 10px;
  min-height: 50vh;
  padding: 2rem;
}

.start-btn {
  max-width: 10vw;
}

#divider {
  border-color: #8e8c91;
}
</style>
