import axios from 'axios';

// 新闻列表
const newsList = () => axios.get('/api/news/all').then(res => res.data);

// 新闻创建
const newsAdd = addForm => axios.post('/api/news/add', addForm).then(res => res.data);

// 新闻删除
const newsDel = id => axios.delete(`/api/news/del/${id}`).then(res => res.data);

// 新闻更新
const newsUpdate = (id, updateForm) => axios.put(`/api/news/update/${id}`, updateForm).then(res => res.data);

export { newsList, newsAdd, newsDel, newsUpdate };
