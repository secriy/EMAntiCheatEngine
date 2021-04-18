import axios from 'axios';

// 列表
const hookList = () => axios.get('/api/v1/list/hook').then(res => res.data);

// 异常列表
const hookExpList = () => axios.get('/api/v1/list/hook/exp').then(res => res.data);

// 统计
const hookCount = () => axios.get('/api/v1/count/hook').then(res => res.data);

export { hookList, hookExpList, hookCount };
