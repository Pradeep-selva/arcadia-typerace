<template>
  <v-dialog v-model="dialog" max-width="400px" persistent>
    <v-card>
      <v-card-title>
        <span class="headline">Welcome to {{ roomCode }}</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="1">
              <v-btn
                class="ma-2"
                text
                icon
                color="blue darken-3"
                @click="onCopy"
              >
                <v-icon color="blue darken-3">
                  mdi-clipboard-multiple
                </v-icon>
              </v-btn>
            </v-col>
            <v-col cols="11" class="d-flex align-center justify-center">
              <h4>Copy the joining code to share</h4>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12" class="d-flex align-center justify-center">
              <h4>
                **NOTE:<br />
                1. Wrongly entered characters will not be read. Simply enter the
                next right character to continue typing<br />2. Completed spaces
                will be represented by _
              </h4>
            </v-col>
          </v-row>
          <v-row>
            <v-col cols="12">
              <v-text-field
                label="Your name*"
                hint="What's your name?"
                v-model="userName"
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
        <v-btn color="blue darken-1" text @click="onConfirm">
          Confirm
        </v-btn>
      </v-card-actions>
    </v-card>
    <v-alert :value="alert" dense elevation="15" type="success" class="mt-16">
      Copied room Code
    </v-alert>
  </v-dialog>
</template>

<script lang="ts">
import { copyToClipBoard } from "@/services";
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class CreateDialog extends Vue {
  dialog = true;
  userName = "";
  alert = false;
  error = "";

  @Prop() roomCode!: string;
  @Prop() setUserName!: Function;

  onConfirm() {
    if (this.userName) {
      this.dialog = false;
      this.setUserName(this.userName);
    } else this.error = "This is a mandatory field.";
  }

  async onCopy() {
    copyToClipBoard(this.roomCode, () => {
      this.alert = true;
      setTimeout(() => (this.alert = false), 2000);
    });
  }
}
</script>
