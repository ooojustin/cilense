import axios from 'axios';

const frontend = "http://localhost:3000/";
const backend = "http://localhost:8080/";
const api = axios.create({
    baseURL: backend
});

export default {
    frontend,
    backend,
    api
}
