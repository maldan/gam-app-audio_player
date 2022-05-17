<template>
  <div :class="$style.main">
    <video
      v-if="$store.state.main.list.length"
      ref="player"
      :src="$store.state.main.list[id].url"
      controls
      :volume="volume"
    ></video>
    <div :class="$style.list">
      <div
        @click="play(i)"
        class="clickable"
        :class="[$style.track, i === id ? $style.selected : null]"
        v-for="(x, i) in $store.state.main.list"
        :key="x.url"
      >
        <div style="font-weight: bold; margin-right: 10px">{{ i + 1 }}</div>
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
    this.id = 0;
  },
  methods: {
    play(id: number): void {
      this.id = id;
      this.$nextTick(() => {
        (this.$refs['player'] as HTMLVideoElement).playbackRate = 1;
        (this.$refs['player'] as HTMLVideoElement).volume = 0.25;
        (this.$refs['player'] as HTMLVideoElement).play();
        (this.$refs['player'] as HTMLVideoElement).onended = () => {
          this.next();
        };
      });
    },
    next() {
      this.id += 1;
      if (this.id > this.$store.state.main.list.length - 1) {
        this.id = 0;
      }
      this.play(this.id);
    },
  },
  data: () => {
    return {
      // track: '',
      id: 0,
      volume: 0.25,
    };
  },
});
</script>

<style lang="scss" module>
.main {
  height: 100%;

  video {
    width: 100%;
    height: 240px;
    display: block;
  }

  .list {
    height: calc(100% - 240px);
    overflow-y: auto;

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
