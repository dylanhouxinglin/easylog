/* eslint-disable */
import axios from 'axios';

const http = axios.create({
  // headers: {
  //   'Content-type': 'application/x-www-form-urlencoded'
  // }
  // todo hrr
  //baseURL: 'http://192.168.242.17:8010',
  timeout: 8000
});
http.interceptors.request.use(
  config => {
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

export default http;
