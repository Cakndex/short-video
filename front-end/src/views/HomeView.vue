<script setup>
import axios from 'axios'
import { ref, onMounted, onUnmounted } from 'vue'
// import { ref } from 'vue'
// 视频列表
const videolist = ref([])
const loading = ref(false)
const body = ref(null)
// 拿视频数据
const road = async () => {
  if (loading.value) return
  loading.value = true
  try {
    const response = await axios.get('/test.json')
    const data = response.data
    videolist.value = videolist.value.concat(data)
    console.log(data)
    console.log(videolist.value)
  } catch (error) {
    console.error(error)
  }
  loading.value = false
}
road()
// 无线滚动
onMounted(() => {
  body.value.addEventListener('scroll', scrollhandle, false)
})
onUnmounted(() => {
  window.removeEventListener('scroll', scrollhandle, false)
})

const scrollhandle = () => {
  // 页面总高
  const scrollHeight = body.value.scrollHeight
  // 滚动距离
  const scrollTop = body.value.scrollTop || body.value.scrollTop
  //可视区高度
  const clientHeight = +724
  const distance = scrollHeight - clientHeight - scrollTop
  console.log(distance + 'distance')
  if (distance < 300) {
    road()
    console.log('该加载了')
  }
}
</script>

<template>
  <CategoryPage></CategoryPage>
  <HeaderPage></HeaderPage>
  <NavPage></NavPage>
  <section id="body" ref="body">
    <button @click="road()">sjsh</button>
    <div class="videolist-left">
      <section
        class="video-left"
        v-for="(item, index) in videolist"
        :key="index"
      >
        <video :src="item.url" preload="true" controls width="250"></video>
        <h1>{{ item.name }}</h1>
      </section>
    </div>
  </section>
</template>

<style scoped lang="less">
#body {
  position: absolute;
  top: 22%;
  left: 17vw;
  width: 83vw;
  height: 78vh;
  overflow-y: scroll;
  overflow-x: hidden;
  border-radius: 20px;
  //   隐藏滚动条
  -ms-overflow-style: none; /* IE and Edge */
  scrollbar-width: none; /* Firefox */
}
#body::-webkit-scrollbar {
  display: none; /* Chrome Safari */
}
.videolist-left {
  width: 25%;
  background-color: aqua;
  .video-left {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
}
</style>
