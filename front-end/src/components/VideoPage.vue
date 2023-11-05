<script setup>
import { ref, defineProps, defineEmits, computed } from 'vue'
import 'vue3-video-play/dist/style.css'
const props = defineProps({
  item: Object,
  showvideo: Boolean,
  url: String
})
// 关闭视频
const emit = defineEmits(['changeshow'])
const changeshowvideo = () => {
  emit('changeshow', false)
}
// 当前视频链接
const url = ref(props.item.url)

// 定义当前激活的视频索引
const activeIndex = ref(0)
// 定义视频列表数据
const videoList = ref([
  {
    url: 'https://prod-streaming-video-msn-com.akamaized.net/178161a4-26a5-4f84-96d3-6acea1909a06/2213bcd0-7d15-4da0-a619-e32d522572c0.mp4'
  },
  {
    url: 'https://prod-streaming-video-msn-com.akamaized.net/a8c412fa-f696-4ff2-9c76-e8ed9cdffe0f/604a87fc-e7bc-463e-8d56-cde7e661d690.mp4'
  },
  {
    url: 'https://prod-streaming-video-msn-com.akamaized.net/fe13f13c-c2cc-4998-b525-038b23bfa9b5/1a9d30ca-54be-411e-8b09-d72ef4488e05.mp4'
  },
  {
    url: 'https://prod-streaming-video-msn-com.akamaized.net/a8c412fa-f696-4ff2-9c76-e8ed9cdffe0f/604a87fc-e7bc-463e-8d56-cde7e661d690.mp4'
  },
  {
    url: 'https://prod-streaming-video-msn-com.akamaized.net/178161a4-26a5-4f84-96d3-6acea1909a06/2213bcd0-7d15-4da0-a619-e32d522572c0.mp4'
  }
  // 视频数据...
])

// 计算属性，根据当前激活的索引获取当前视频数据
const activeVideo = computed(() => videoList.value[activeIndex.value])
const nextVideo = computed(() => videoList.value[activeIndex.value + 1])
// 上一个视频
const transformY = ref(0)
function handleKeyDown(event) {
  if (event.key === 'ArrowUp') {
    // 上键切换到上一个视频
    activeIndex.value =
      (activeIndex.value - 1 + videoList.value.length) % videoList.value.length
    alert(activeIndex.value)
    transformY.value += 510
  } else if (event.key === 'ArrowDown') {
    // 下键切换到下一个视频
    activeIndex.value = (activeIndex.value + 1) % videoList.value.length
    alert(activeIndex.value)
    transformY.value -= 510
  }
}
// 评论
const comment = ref([
  {
    date: {
      year: '2022',
      month: '10',
      date: '2',
      hour: '12',
      minute: '13',
      second: '45'
    },
    content: '这个视频好棒'
  },
  {
    date: {
      year: '2022',
      month: '10',
      date: '2',
      hour: '12',
      minute: '13',
      second: '45'
    },
    content: '这个视频真不错'
  }
])
const content = ref('')
const commit = (value) => {
  // 获取当前的年份
  const currentYear = new Date().getFullYear()
  // 获取当前的月份（注意月份是从0开始计数的，所以需要加1）
  const currentMonth = new Date().getMonth() + 1
  // 获取当前的日期
  const currentDate = new Date().getDate()
  // 获取当前的小时
  const currentHour = new Date().getHours()
  // 获取当前的分钟
  const currentMinute = new Date().getMinutes()
  // 获取当前的秒数
  const currentSecond = new Date().getSeconds()
  comment.value.unshift({
    date: {
      year: currentYear,
      month: currentMonth,
      date: currentDate,
      hour: currentHour,
      minute: currentMinute,
      second: currentSecond
    },
    content: value
  })
  content.value = ''
}
</script>

<template>
  <div class="bg" @keydown.prevent="handleKeyDown">
    <!-- 关闭按钮 -->
    <div class="icon" @click="changeshowvideo">
      <p>{{ props.showvideo }}</p>
      <svg
        t="1699111590427"
        class="icon"
        viewBox="0 0 1024 1024"
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        p-id="8279"
        width="48"
        height="48"
      >
        <path
          d="M865.28 921.6c-14.336 0-28.672-5.632-39.936-16.384l-706.56-706.56c-22.016-22.016-22.016-57.856 0-79.872 22.016-22.016 57.856-22.016 79.872 0l706.56 706.56c22.016 22.016 22.016 57.856 0 79.872-11.264 10.752-25.6 16.384-39.936 16.384z"
          fill="#ffffff"
          p-id="8280"
        ></path>
        <path
          d="M158.72 921.6c-14.336 0-28.672-5.632-39.936-16.384-22.016-22.016-22.016-57.856 0-79.872l706.56-706.56c22.016-22.016 57.856-22.016 79.872 0 22.016 22.016 22.016 57.856 0 79.872l-706.56 706.56c-11.264 10.752-25.6 16.384-39.936 16.384z"
          fill="#ffffff"
          p-id="8281"
        ></path>
      </svg>
    </div>
    <section class="video">
      <div class="left">
        <div
          class="container"
          :style="{ transform: `translateY(${transformY}px)` }"
        >
          <div><video :src="url" autoplay controls></video></div>
          <div><video :src="activeVideo.url" autoplay controls></video></div>
          <div><video :src="nextVideo.url" autoplay controls></video></div>
        </div>
      </div>
      <div class="right">
        <div class="introduce">
          <h1>测试视频</h1>
          <p>
            &nbsp;&nbsp;痛了10年了，止痛片已经渐渐感觉没用了。真的，这种痛...懂的都懂，不懂的我希望你永远不要懂!早上起来发现来大姨妈了怕下午肚子疼没法拍照就提前吃了布洛芬然后画完妆突然一阵剧痛
          </p>
        </div>
        <div class="comment">
          <p class="comment-head">评论🎉</p>
          <div
            class="comment-content"
            v-for="(item, index) in comment"
            :key="index"
          >
            <img
              class="img"
              src="@/assets/usericon.png"
              alt=""
              width="35"
              height="35"
            />
            <div>
              <span class="date"
                >{{ item.date.year }}/{{ item.date.month }}/{{
                  item.date.date
                }}</span
              >
              <p>
                {{ item.content }}
              </p>
            </div>
          </div>
        </div>
        <div class="input">
          <input type="text" v-model="content" @keyup.enter="commit(content)" />
          <button @click="commit(content)">发送</button>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped lang="less">
.bg {
  position: fixed;
  left: 0;
  top: 0;
  display: flex;
  width: 100vw;
  height: 100vh;
  z-index: 999;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.4);
}
.video {
  z-index: 999;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 80vw;
  height: 68vh;
  display: flex;
  border-radius: 20px;
  background-color: #000000;
  .left {
    width: 75%;
    height: 100%;
    overflow-y: hidden;
    overflow-x: hidden;
    .container {
      height: 100%;
      margin: 0 0;
      transition: all 0.2s ease-out;
      video {
        height: 505px;
      }
    }
  }
  .right {
    width: 25%;
    height: 100%;
    background-color: #f7f7f7;
    border-radius: 0 20px;
    // 视频简介
    .introduce {
      padding: 20px;
      height: max-content;
      background-color: #fff;
      h1 {
        text-align: center;
        margin: 0 0 10px 0;
      }
    }
    // 评论区
    .comment {
      height: 300px;
      overflow-y: scroll;
      position: relative;
      .comment-head {
        padding-top: 10px;
        margin-left: 10px;
        position: sticky;
        color: #ff2741;
        font-weight: 600;
      }
      .comment-content {
        width: 300px;
        display: flex;
        .img {
          margin-left: 10px;
          margin-top: 10px;
        }
        .date {
          font-size: 12px;
          color: #858585;
          margin-left: 20px;
        }
        p {
          width: 200px;
          margin: 0 20px 15px;
          word-wrap: break-word; /* 在单词内部进行换行 */
        }
      }
    }
    .comment::-webkit-scrollbar {
      display: none; /* Chrome Safari */
    }
    // 写评论
    .input {
      position: absolute;
      bottom: 0;
      right: 0;
      width: 25%;
      height: 10%;
      background-color: #ece9e9;
      display: flex;
      align-items: center;
      border-radius: 0 0 20px 0;
      input {
        margin-left: 10px;
        padding-left: 10px;
        border-radius: 20px;
        border: #f4bfbf 2px solid;
        height: 30px;
        width: 13vw;
      }
      input:focus {
        outline: none;
        border: #e43939 2px solid;
      }
      button {
        letter-spacing: 3px;
        padding: 10px 10px;
        margin: 0 10px;
        border: none;
        border-radius: 25px;
        background-color: #ff2442;
        color: #fff;
        font-weight: 600;
        cursor: pointer;
      }
    }
  }
}
.icon {
  position: absolute;
  right: 5%;
  top: 5%;
  cursor: pointer;
}
</style>
