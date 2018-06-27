import axios from 'axios';

export default {
  signup: user => axios.post('/api/user/signup', user).then(res => res.data),
  authorize: () => axios.get('/api/user/authorize').then(res => res.data),
  login: user => axios.post('/api/user/login', user).then(res => res.data),
  setMultipleSections: isMulti =>
    axios.post('api/user/multipleSections', isMulti).then(res => res.data),
  addSection: section =>
    axios.post('/api/user/addSection', section).then(res => res.data),
  delSection: section =>
    axios.post('/api/user/delSection', section).then(res => res.data)
};