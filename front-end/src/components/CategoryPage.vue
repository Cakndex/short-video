<script setup>
import { ref } from 'vue'
const category = [
  '推荐',
  '穿搭',
  '美食',
  '彩妆',
  '影视',
  '职场',
  '情感',
  '家具',
  '游戏',
  '旅行',
  '健身'
]
const nowindex = ref(0)
console.log(nowindex.value)
const changeindex = (index) => {
  nowindex.value = index
}
</script>

<template>
  <div class="category">
    <ul>
      <li
        v-for="(item, index) in category"
        :key="index"
        :class="{ active: index === nowindex }"
        @click="changeindex(index)"
      >
        {{ item }}
      </li>
    </ul>
  </div>
</template>

<style scoped lang="less">
.category {
  width: 100%;
  position: fixed;
  top: 10%;
  left: 20vw;
  z-index: 10;
}
.category ul {
  display: flex;
  flex-direction: row;
  li {
    list-style: none;
    padding: 10px 20px;
    color: #787878;
    transition: all 0.1s ease-out;
    caret-color: transparent; /* 将光标颜色设置为透明 */
  }
  li:hover {
    color: #333333;
    font-weight: 600;
    cursor: pointer;
    background-color: #f7f7f7;
    border-radius: 20px;
  }
}
.active {
  font-weight: 600;
  cursor: pointer;
  background-color: #f7f7f7;
  border-radius: 20px;
  color: #333333 !important;
}
</style>
