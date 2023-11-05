<template>
  <div>
    <vue3VideoPlay v-bind="options" />
  </div>
</template>

<script setup>
import 'vue3-video-play/dist/style.css'
import vue3VideoPlay from 'vue3-video-play'
import { reactive, defineProps } from 'vue'

const props = defineProps({
  video_url: String
})

const options = reactive({
  width: '910px',
  height: '510px',
  color: '#046286',
  title: '',
  src: props.video_url,
  muted: false,
  webFullScreen: false,
  speedRate: ['0.75', '1.0', '1.25', '1.5', '2.0'],
  autoPlay: true,
  loop: false,
  mirror: false,
  ligthOff: true,
  volume: 0.3,
  control: true,
  controlBtns: [
    'audioTrack',
    'quality',
    'speedRate',
    'volume',
    'setting',
    'pip',
    'pageFullScreen',
    'fullScreen'
  ]
})
</script>
