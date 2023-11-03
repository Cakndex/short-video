<script setup></script>

<template>
  <div class="category">
    <ul>
      <li>推荐</li>
      <li>穿搭</li>
      <li>美食</li>
      <li>彩妆</li>
      <li>影视</li>
      <li>职场</li>
      <li>情感</li>
      <li>家具</li>
      <li>游戏</li>
      <li>旅行</li>
      <li>健身</li>
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
</style>
