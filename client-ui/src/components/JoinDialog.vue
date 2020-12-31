<template>
  <v-dialog v-model="dialog" max-width="400px">
    <template v-slot:activator="{ on, attrs }">
      <v-btn
        class="mr-10 font-weight-bold"
        color="grey lighten-1"
        outlined
        x-large
        v-bind="attrs"
        v-on="on"
      >
        Join Room
      </v-btn>
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">Join A Room</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <v-text-field
                label="Room Code*"
                hint="Enter a room code to join"
                v-model="roomCode"
                :error="!!error"
                :error-messages="error"
                required
              />
            </v-col>
          </v-row>
        </v-container>
        <small>*indicates required field</small>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="dialog = false">
          Close
        </v-btn>
        <v-btn color="blue darken-1" text @click="onConfirm">
          Confirm
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { Component, Vue } from "vue-property-decorator";
import { ROUTES, VAL_TYPES } from "../configs";

@Component
export default class JoinDialog extends Vue {
  dialog = false;
  roomCode = "";
  error = "";

  onConfirm() {
    if (this.roomCode.length <= 4 && !!this.roomCode)
      this.$router.push({
        path: ROUTES.validate,
        params: {
          valType: VAL_TYPES.join,
          roomId: this.roomCode
        }
      });
    else this.error = "Enter a valid room code";
  }
}
</script>
