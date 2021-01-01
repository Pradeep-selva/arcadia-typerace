<template>
  <v-container
    class="validate
   d-flex 
   justify-center 
   flex-column
   mt-16"
    fluid
  >
    <room-dialog :roomCode="roomId" />
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
          <span :class="`${incompleteTextColor}--text`">
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
import { RoomDialog } from "../components";

@Component({
  components: {
    RoomDialog
  }
})
export default class Room extends Vue {
  roomId = this.$route.params.id;
  randomWords = (RandomWords(30) as string[]).join(" ");
  incompleteString = this.randomWords;
  incompleteTextColor = "white";
  completeString = "";
  enteredTextCount = 0;

  mounted() {
    const ws = new WebSocket(API_ENDPOINTS.socket(this.roomId));

    document.addEventListener("keypress", (event) => {
      if (event.keyCode == 32) event.preventDefault();

      if (event.key === this.incompleteString[0])
        ws.send(this.randomWords.slice(0, ++this.enteredTextCount));
      else {
        this.incompleteTextColor = "red";
        setTimeout(() => (this.incompleteTextColor = "white"), 100);
      }
    });

    ws.addEventListener("message", (event) => {
      const count = event.data.length;

      this.incompleteString = this.randomWords.slice(
        count,
        this.randomWords.length
      );
      this.completeString = this.randomWords.slice(0, count);
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
