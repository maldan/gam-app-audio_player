<template>
  <div :class="$style.main">
    <video ref="player" :src="track" controls></video>
    <div :class="$style.list">
      <div
        @click="play(x.url)"
        class="clickable"
        :class="[$style.track, x.url === track ? $style.selected : null]"
        v-for="x in $store.state.main.list"
        :key="x.url"
      >
        <div>{{ x.name }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  components: {},
  async mounted() {
    await this.$store.dispatch('main/getList');
  },
  methods: {
    play(url: string): void {
      this.track = url;
      // ab
      this.$nextTick(() => {
        (this.$refs['player'] as HTMLVideoElement).playbackRate = 1;
        (this.$refs['player'] as HTMLVideoElement).volume = 0.25;
        (this.$refs['player'] as HTMLVideoElement).play();
      });
    },
  },
  data: () => {
    return {
      track: '',
    };
  },
});
</script>

<style lang="scss" module>
.main {
  .list {
    .track {
      display: flex;
      padding: 10px;
      background: #242424;
      font-size: 14px;

      &.selected {
        background: #1b4988;
      }
    }
  }
}
</style>
