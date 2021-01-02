<template>
  <v-container
    class="validate
   d-flex 
   justify-center 
   flex-column
   mt-16"
    fluid
  >
    <room-dialog :roomCode="roomId" :setUserName="onJoin" />
    <h1
      class="grey--text 
    text--lighten-1 
    d-flex 
    justify-center"
    >
      <span class="mr-3">ROOM</span>
      <span class="text--lighten-2">{{ roomId }}</span>
    </h1>
    <div
      class="d-flex 
    justify-center 
    mb-5 mt-5"
      v-if="!!firstUser && !!secondUser"
    >
      <v-btn
        class="start-btn"
        color="grey lighten-1"
        outlined
        x-large
        @click="onPressStart"
        v-if="firstUser === curUser && counter == null"
      >
        start race
      </v-btn>
      <div
        class="display-1 green--text mb-5 mt-5"
        v-if="firstUser !== curUser && counter == null"
      >
        Please be patient while the room owner starts race...
      </div>
      <div
        id="counter"
        class="display-4 green--text mb-5 mt-5"
        v-if="counter !== null && counter !== 0"
      >
        Starts in {{ counter }}
      </div>
      <div
        id="counter"
        class="display-4 green--text mb-5 mt-5"
        v-if="counter === 0"
      >
        Start Typing
      </div>
    </div>
    <div
      class="d-flex 
    display-2
    justify-center
    grey--text
    text--lighten-1 
    mb-10 mt-5"
      v-if="!!firstUser && !!secondUser"
    >
      {{ firstUser }} vs. {{ secondUser }}
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
    <v-alert :value="showAlert" dense elevation="15" type="success" id="alert">
      {{ newJoinedUser }} joined room
    </v-alert>
  </v-container>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import RandomWords from "random-words";
import { API_ENDPOINTS, Events, EventResponse } from "@/configs";
import { RoomDialog } from "../components";

@Component({
  components: {
    RoomDialog
  }
})
export default class Room extends Vue {
  ws: WebSocket | null = null;
  curUser = "";
  roomId = this.$route.params.id;
  firstUser = "";
  secondUser = "";
  randomWords = (RandomWords(30) as string[]).join(" ");
  incompleteString = this.randomWords;
  incompleteTextColor = "white";
  newJoinedUser = "";
  showAlert = false;
  completeString = "";
  counter: number | null = null;
  enteredTextCount = 0;

  onJoin(userName: string) {
    const ws = new WebSocket(API_ENDPOINTS.socket(this.roomId, userName));
    this.ws = ws;
    this.curUser = userName;

    this.firstUser = userName;

    ws.addEventListener("message", (e) => {
      console.log(e);
      const event: EventResponse = JSON.parse(e.data);

      switch (event.event) {
        case Events.JOIN:
          if (event?.userName !== userName) {
            this.secondUser = event?.userName;
            ws.send(
              JSON.stringify({
                event: Events.RECEIVED_TEST,
                data: this.randomWords,
                userName
              } as EventResponse)
            );
          }

          this.newJoinedUser = event?.userName;
          this.showAlert = true;
          setTimeout(() => (this.showAlert = false), 2000);
          break;

        case Events.TYPING: {
          const count = event?.data.length;

          this.incompleteString = this.randomWords.slice(
            count,
            this.randomWords.length
          );
          this.completeString = this.randomWords.slice(0, count);
          break;
        }

        case Events.RECEIVED_TEST:
          if (event?.userName !== userName) {
            this.firstUser = event?.userName;
            this.secondUser = userName;
            this.randomWords = event?.data;
            this.incompleteString = this.randomWords;
          }
          break;

        case Events.INITIATE_START:
          this.onStartRace();
          break;

        default:
          console.log("Error encountered");
      }
    });
  }

  initKeyPressListener() {
    document.addEventListener("keypress", (event) => {
      if (event.keyCode == 32) event.preventDefault();

      if (event.key === this.incompleteString[0])
        this.ws?.send(
          JSON.stringify({
            event: Events.TYPING,
            data: this.randomWords.slice(0, ++this.enteredTextCount),
            userName: this.curUser
          } as EventResponse)
        );
      else {
        this.incompleteTextColor = "red";
        setTimeout(() => (this.incompleteTextColor = "white"), 100);
      }
    });
  }

  onPressStart() {
    this.ws?.send(
      JSON.stringify({
        event: Events.INITIATE_START
      })
    );
  }

  onStartRace() {
    this.counter = 3;
    const interval = setInterval(() => {
      (this.counter as number)--;

      if (this.counter === 0) {
        window.clearInterval(interval);
        this.initKeyPressListener();
      }
    }, 1000);
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

#alert {
  position: fixed;
  bottom: 50px;
  left: 100px;
}
</style>
