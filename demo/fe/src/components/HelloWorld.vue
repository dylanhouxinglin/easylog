<template>
  <div>
    <el-row>
      <el-button type="primary" @click="req()">Get</el-button>
      <el-button type="danger">Post</el-button>
    </el-row>
  </div>
</template>

<script>
/* eslint-disable */
import http from '../assets/http.js';
export default {
  name: 'HelloWorld',
  methods: {
    req() {
      http.get('/demo/ready').then((response) => {
        console.log(response)
      })
    }
  }
}
</script>

<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
