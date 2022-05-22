<template>
  <div :class="$style.main">
    <video
      v-if="$store.state.main.list.length"
      ref="player"
      :src="$store.state.main.list[$store.state.main.settings.trackId].url"
      controls
      :volume="$store.state.main.settings.volume"
    ></video>
    <div :class="$style.list">
      <div
        @click="play(i)"
        class="clickable"
        :class="[$style.track, i === $store.state.main.settings.trackId ? $style.selected : null]"
        v-for="(x, i) in $store.state.main.list"
        :key="x.url"
      >
        <div style="font-weight: bold; margin-right: 10px">{{ i + 1 }}</div>
        <div :class="$style.name">{{ x.name }}</div>
        <div :class="$style.duration">
          {{ $root.moment.utc(x.duration * 1000).format('HH:mm:ss') }}
        </div>
        <div :class="$style.size">{{ $root.prettyBytes(x.size) }}</div>
      </div>

      <div :class="[$style.track, $style.total]">
        <div :class="$style.name">Total</div>
        <div :class="$style.duration">
          {{ $root.moment.utc(totalDuration * 1000).format('HH:mm:ss') }}
        </div>
        <div :class="$style.size">{{ $root.prettyBytes(totalSize) }}</div>
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
    await this.$store.dispatch('main/getSettings');

    (this.$refs['player'] as HTMLVideoElement).onended = () => {
      this.next();
    };
    (this.$refs['player'] as HTMLVideoElement).onvolumechange = () => {
      this.$store.dispatch('main/changeVolume', (this.$refs['player'] as HTMLVideoElement).volume);
    };
    setInterval(() => {
      this.$store.state.main.settings.currentTime = (
        this.$refs['player'] as HTMLVideoElement
      ).currentTime;
      this.$store.dispatch('main/saveSettings');
    }, 5000);

    await this.$nextTick(() => {
      (this.$refs['player'] as HTMLVideoElement).currentTime =
        this.$store.state.main.settings.currentTime;
    });

    // Total size
    this.$store.state.main.list.forEach((x) => {
      this.totalDuration += x.duration;
      this.totalSize += x.size;
    });
  },
  methods: {
    play(id: number): void {
      this.$store.dispatch('main/changeTrack', id);

      this.$nextTick(() => {
        (this.$refs['player'] as HTMLVideoElement).playbackRate = 1;
        (this.$refs['player'] as HTMLVideoElement).play();
      });
    },
    next() {
      let trackId = this.$store.state.main.settings.trackId;
      trackId += 1;
      if (trackId > this.$store.state.main.list.length - 1) {
        trackId = 0;
      }
      this.$store.dispatch('main/changeTrack', trackId);
      this.play(trackId);
    },
  },
  data: () => {
    return {
      totalDuration: 0,
      totalSize: 0,
    };
  },
});
</script>

<style lang="scss" module>
.main {
  height: 100%;

  video {
    width: 100%;
    height: 300px;
    display: block;
  }

  .list {
    height: calc(100% - 300px);
    overflow-y: auto;
    width: 100%;

    .track {
      display: flex;
      padding: 10px;
      background: #242424;
      font-size: 14px;

      .name {
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow-x: hidden;
        padding-right: 5px;
      }

      .duration {
        margin-left: auto;
        background: #05679c;
        padding: 2px 5px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
      }

      .size {
        margin-left: 10px;
        background: #359c05;
        padding: 2px 5px;
        border-radius: 4px;
        font-size: 12px;
        white-space: nowrap;
      }

      &.selected {
        background: #1b4988;
      }
    }

    .total {
      background: #2e2e2e;
    }
  }
}
</style>
