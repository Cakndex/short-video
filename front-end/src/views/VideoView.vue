<script setup>
import { ref } from 'vue'
const queryString = window.location.search
const url = queryString.slice(1)
const videoRef = ref(null)
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
    content: 'shadk方式是ds'
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
    content: 'shadk方式是ds'
  }
])
// 评论
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
  <div class="bg">
    <section class="video">
      <div class="left">
        <video
          :src="url"
          preload="true"
          height="502"
          ref="videoRef"
          controls
        ></video>
      </div>
      <div class="right">
        <div class="introduce">
          <h1>测试视频</h1>
          <p>
            &nbsp;&nbsp;这是简介Lorem ipsum dolor sit amet consectetur
            adipisicing elit. Optio voluptatibus odio dicta s
          </p>
        </div>
        <div class="comment">
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
  z-index: 20;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.25);
}
.video {
  z-index: 999;
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  width: 80vw;
  height: 80vh;
  display: flex;
  border-radius: 20px;
  background-color: #333;
  .left {
    width: 75%;
    height: 100%;

    video {
      margin-top: 45px;
    }
  }
  .right {
    width: 25%;
    height: 100%;
    background-color: #ece9e9;
    border-radius: 0 20px;
    // 视频简介
    .introduce {
      height: 25%;
      border-bottom: #fff 2px;
      h1 {
        text-align: center;
        margin: 0;
      }
    }
    // 评论区
    .comment {
      height: 65%;
      overflow-y: scroll;
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
</style>
