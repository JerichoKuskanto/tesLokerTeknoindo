<template>
  <div id="app">
    <div class="content">
      <HelloWorld msg="Welcome to Your Vue.js App" />
      <router-view></router-view>
    </div>
  </div>
</template>

<style>
html, body {
  margin: 0; 
  padding: 0;
  background-color: #121212; 
}

#app {
  background-image: url('assets/bg.jpg'); 
  background-size: cover; 
  background-position: center; 
  background-repeat: no-repeat; 
  min-height: 100vh; 
  color: #ffffff;
}
</style>
