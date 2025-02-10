const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    allowedHosts: "all", // Fix "Invalid Host Header"
    port: 8080
  }
});
