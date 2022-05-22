<template>
  <router-view />
  <ui-modal-view />
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import PrettyBytes from 'pretty-bytes';
import Moment from 'moment';

export default defineComponent({
  components: {},
  async mounted() {
    // b
  },
  methods: {},
  data: () => {
    return {
      // @ts-ignore
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      console: console,
      prettyBytes: PrettyBytes,
      moment: Moment,
    };
  },
});
</script>

<style lang="scss" scoped></style>
