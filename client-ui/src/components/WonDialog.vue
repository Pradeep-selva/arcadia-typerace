<template>
  <v-dialog v-model="dialog" max-width="400px">
    <v-card>
      <v-card-title>
        <span class="headline">{{ winner }} is faster! 😎</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="6" class="text-center">
              <h3>{{ first }}'s WPM</h3>
              <br />
              <h1>{{ firstWPM }}</h1>
            </v-col>
            <v-col cols="6" class="text-center">
              <h3>{{ second }}'s WPM</h3>
              <br />
              <h1>{{ secondWPM }}</h1>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="onConfirm">
          Go Home
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class CreateDialog extends Vue {
  @Prop() dialog!: boolean;
  @Prop() winner!: string;
  @Prop() first!: string;
  @Prop() firstWPM!: string;
  @Prop() second!: string;
  @Prop() secondWPM!: string;

  onConfirm() {
    window.location.replace("abc");
  }
}
</script>
