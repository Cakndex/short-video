<script setup>
import { ref } from 'vue'
const inputval = ref('')
</script>

<template>
  <footer>
    <a href="#">
      <img src="@/assets/logo.png" alt="" width="128" />
    </a>
    <div class="input-container">
      <input
        v-model="inputval"
        type="text"
        spellcheck="false"
        class="search-input"
        placeholder="搜你所想"
        autocomplete="false"
      />
      <svg
        t="1698941773287"
        class="icon"
        viewBox="0 0 1024 1024"
        version="1.1"
        xmlns="http://www.w3.org/2000/svg"
        p-id="49037"
        width="24"
        height="24"
      >
        <path
          d="M351.1808 59.2896A435.2 435.2 0 0 1 805.376 715.264 460.8 460.8 0 0 1 351.1808 59.3408z"
          fill="#20C997"
          p-id="49038"
        ></path>
        <path
          d="M754.3808 722.2272a358.4 358.4 0 1 0-267.8272 120.2176 51.2 51.2 0 0 1 0 102.4 460.8 460.8 0 1 1 365.1584-179.712l118.8864 121.2416c23.7568 24.2176 23.552 63.0272-0.4096 87.04l-0.4096 0.4096a61.184 61.184 0 0 1-86.9888-0.4608l-148.0192-150.9376a61.7984 61.7984 0 0 1 0.4096-86.9888l0.4096-0.4096c5.632-5.5808 11.9808-9.8304 18.7904-12.8z m-467.968-364.5952h409.6a51.2 51.2 0 1 1 0 102.4h-409.6a51.2 51.2 0 1 1 0-102.4z m0 204.8h256a51.2 51.2 0 0 1 0 102.4h-256a51.2 51.2 0 1 1 0-102.4z"
          fill="#2C6DD2"
          p-id="49039"
        ></path>
      </svg>
    </div>
    <div class="text">
      <button><span>创作中心</span></button>
      <button><span>商务合作</span></button>
    </div>
  </footer>
</template>

<style scoped lang="less">
footer {
  box-sizing: border-box;
  padding: 0 32px;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  height: 72px;
  z-index: 10;
  overflow: hidden;
  caret-color: transparent; /* 将光标颜色设置为透明 */
}
// 搜索框
.input-container {
  height: 40px;
  width: 400px;
  position: fixed;
  left: 50%;
  transform: translate(-50%);
  .search-input {
    padding: 0 84px 0 16px;
    width: 75%;
    height: 100%;
    font-size: 16px;
    line-height: 120%;
    border-radius: 40px;
    border: none;
    background-color: #f7f7f7;
    caret-color: #ee6775;
  }
}
.search-input:focus {
  outline: none;
}

// 文字部分
.text {
  width: 280px;
  display: flex;
  justify-content: space-between;
  button {
    padding: 10px 15px;
    margin: 0 20px;
    border: none;
    background: none;
    cursor: pointer;
    transition: all 0.1s ease-out;
    span {
      font-size: 1rem;
      font-weight: 600;
    }
  }
  button:hover {
    background-color: #f7f7f7;
    border-radius: 40px;
  }
}
.icon {
  position: absolute;
  right: 20px;
  top: 50%;
  transform: translateY(-50%);
  cursor: pointer;
}
</style>
