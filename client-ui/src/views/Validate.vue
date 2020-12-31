<template>
  <div class="about">
    <h1>Validate</h1>
    <br />
    <h3>{{ roomId }} {{ valType }}</h3>
  </div>
</template>

<script lang="ts">
import { VAL_TYPES } from "@/configs";
import { Vue, Component } from "vue-property-decorator";
import { validateOperation } from "../apis";

@Component
export default class Validate extends Vue {
  valType = this.$route.params.valType;
  roomId = this.$route.params.roomId;

  mounted() {
    validateOperation(this.valType as keyof typeof VAL_TYPES, this.roomId).then(
      console.log
    );
  }
}
</script>
