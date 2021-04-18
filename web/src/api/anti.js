import axios from 'axios';

// 列表
const antiList = () => axios.get('/api/v1/list/anti').then(res => res.data);

// 异常列表
const antiExpList = () => axios.get('/api/v1/list/anti/exp').then(res => res.data);

// 统计
const antiCount = () => axios.get('/api/v1/count/anti').then(res => res.data);

export { antiList, antiExpList, antiCount };
