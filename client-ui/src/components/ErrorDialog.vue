<template>
  <v-dialog v-model="dialog" max-width="400px" persistent>
    <v-card>
      <v-card-title>
        <span class="headline">An Error Occured</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <h4>Please go back home and retry.</h4>
            </v-col>
          </v-row>
        </v-container>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn color="blue darken-1" text @click="onConfirm">
          Go home
        </v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script lang="ts">
import { DEPLOYED_URL } from "@/configs";
import { Vue, Component, Prop } from "vue-property-decorator";

@Component
export default class CreateDialog extends Vue {
  @Prop() dialog!: boolean;

  onConfirm() {
    window.location.replace(DEPLOYED_URL);
  }
}
</script>
