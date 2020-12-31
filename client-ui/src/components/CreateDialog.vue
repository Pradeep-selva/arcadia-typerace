<template>
  <v-dialog v-model="dialog" max-width="400px">
    <template v-slot:activator="{ on, attrs }">
      <v-btn color="grey lighten-1" outlined x-large v-bind="attrs" v-on="on">
        Create Room
      </v-btn>
    </template>
    <v-card>
      <v-card-title>
        <span class="headline">Create A Room</span>
      </v-card-title>
      <v-card-text>
        <v-container>
          <v-row>
            <v-col cols="12">
              <h4>Are you sure you'd like to create a new room?</h4>
            </v-col>
          </v-row>
        </v-container>
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

<script>
import { API_PATHS, VAL_TYPES } from "@/configs";
import { generateUID } from "@/services";
import { Vue, Component } from "vue-property-decorator";

@Component
export default class CreateDialog extends Vue {
  dialog = false;

  onConfirm() {
    this.$router.push({
      path: API_PATHS.validate,
      params: {
        valType: VAL_TYPES.create,
        roomId: generateUID()
      }
    });
  }
}
</script>
