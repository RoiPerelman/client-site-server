import axios from 'axios';

export default {
  signup: user => axios.post('/api/user/signup', user).then(res => res.data),
  authorize: () => axios.get('/api/user/authorize').then(res => res.data),
  login: user => axios.post('/api/user/login', user).then(res => res.data)
};
